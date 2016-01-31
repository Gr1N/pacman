package oauth2

import (
	"strings"
	"time"

	"github.com/Unknwon/com"

	"github.com/Gr1N/pacman/models"
	"github.com/Gr1N/pacman/modules/cache"
	"github.com/Gr1N/pacman/modules/helpers"
	"github.com/Gr1N/pacman/modules/settings"
)

const (
	stateLength = 32
)

var (
	gitHub            *Config
	supportedServices map[string]*Config
	stateCacheTimeout time.Duration
)

// Init initializes oauth2 module.
func Init() {
	auths := settings.S.OAuth2

	gitHub = NewGitHub(auths.GitHubClientID, auths.GitHubClientSecret,
		auths.GitHubRedirectURL, auths.GitHubScopes)

	supportedServices = map[string]*Config{
		"github": gitHub,
	}
	stateCacheTimeout, _ = time.ParseDuration(auths.StateCacheTimeout)
}

// HandleService validates that service name is valid and allowed.
func HandleService(serviceName string) error {
	if !com.IsSliceContainsStr(settings.S.OAuth2.EnabledServices, serviceName) {
		return errServiceNotSupported
	}

	return nil
}

// HandleAuthorizeRequest generates authorization URL for specified service.
func HandleAuthorizeRequest(serviceName string) string {
	state := issueState(serviceName)
	service := supportedServices[serviceName]

	return service.AuthCodeURL(state)
}

// ValidateAuthorizeRequest validates that state value is valid.
func ValidateAuthorizeRequest(serviceName, state string) error {
	if !verifyState(serviceName, state) {
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

func issueState(serviceName string) string {
	state := helpers.RandomString(stateLength)

	key := makeStateCacheKey(state)
	go cache.Set(key, serviceName, stateCacheTimeout)

	return state
}

func verifyState(serviceName, state string) bool {
	var cachedServiceName string

	key := makeStateCacheKey(state)
	if err := cache.Get(key, &cachedServiceName); err != nil {
		return false
	}

	go cache.Delete(key)

	if cachedServiceName != serviceName {
		return false
	}

	return true
}

func makeStateCacheKey(state string) string {
	return strings.Join([]string{
		"auth",
		state,
	}, ":")
}
