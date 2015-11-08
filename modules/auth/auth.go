package auth

import (
	"crypto/subtle"
	"strings"
	"time"

	"github.com/Unknwon/com"

	"github.com/Gr1N/pacman/models"
	"github.com/Gr1N/pacman/modules/auth/oauth2"
	"github.com/Gr1N/pacman/modules/cache"
	"github.com/Gr1N/pacman/modules/errors"
	"github.com/Gr1N/pacman/modules/helpers"
	"github.com/Gr1N/pacman/modules/settings"
)

const (
	stateLength = 32
)

var (
	gitHub            *oauth2.Config
	supportedServices map[string]*oauth2.Config
	stateCacheTimeout time.Duration

	errServiceNotSupported = errors.New(
		"authorization_service_not_supported", "Specified service is not supported")
	errStateInvalid = errors.New(
		"state_invalid", "State value is not valid")
)

// Init initializes auth module.
func Init() {
	auths := settings.S.Auth

	gitHub = oauth2.NewGitHub(auths.GitHubClientID, auths.GitHubClientSecret,
		auths.GitHubRedirectURL, auths.GitHubScopes)

	supportedServices = map[string]*oauth2.Config{
		"github": gitHub,
	}
	stateCacheTimeout, _ = time.ParseDuration(auths.StateCacheTimeout)
}

// HandleService validates that service name is valid and allowed.
func HandleService(serviceName string) error {
	if !com.IsSliceContainsStr(settings.S.Auth.EnabledServices, serviceName) {
		return errServiceNotSupported
	}

	return nil
}

// HandleAuthorizeRequest generates authorization URL for specified service.
func HandleAuthorizeRequest(serviceName, sessionID string) string {
	state := issueState(serviceName, sessionID)
	service := supportedServices[serviceName]

	return service.AuthCodeURL(state)
}

// ValidateAuthorizeRequest validates that state value is valid.
func ValidateAuthorizeRequest(serviceName, sessionID, state string) error {
	cachedState, err := retriveState(serviceName, sessionID)
	if err != nil {
		return errStateInvalid
	}

	if eq := subtle.ConstantTimeCompare([]byte(state), []byte(cachedState)); eq != 1 {
		return errStateInvalid
	}

	return nil
}

// FinishAuthorizeRequest validates code value and tries to find user.
func FinishAuthorizeRequest(serviceName, code string) (*models.User, error) {
	service := supportedServices[serviceName]

	token, err := service.Exchange(code)
	if err != nil {
		return nil, err
	}

	serviceUser, err := service.User(token)
	if err != nil {
		return nil, err
	}

	if user, err := models.GetUserByService(serviceName, serviceUser.ID); err == nil {
		return user, nil
	}

	user, _ := models.CreateUserByService(serviceName, token.Access, serviceUser.ID,
		serviceUser.Name, serviceUser.Email)
	return user, nil
}

func issueState(serviceName, sessionID string) string {
	state := helpers.RandomString(stateLength)

	key := makeStateCacheKey(serviceName, sessionID)
	go cache.Set(key, state, stateCacheTimeout)

	return state
}

func retriveState(serviceName, sessionID string) (string, error) {
	var state string

	key := makeStateCacheKey(serviceName, sessionID)
	if err := cache.Get(key, &state); err != nil {
		return "", err
	}

	go cache.Delete(key)

	return state, nil
}

func makeStateCacheKey(serviceName, sessionID string) string {
	return strings.Join([]string{
		sessionID,
		serviceName,
	}, ":")
}
