package social

import (
	"crypto/subtle"
	"errors"
	"strings"

	"github.com/revel/revel"
	"github.com/revel/revel/cache"

	"github.com/Gr1N/pacman/app/models"
	"github.com/Gr1N/pacman/app/modules/helpers"
)

const (
	stateLength = 32
	codeLength  = 20
)

var (
	ErrServiceRequired = errors.New("Service invalid or disabled")
	ErrStateRequired   = errors.New("State does not match requirements")
	ErrStateNotFound   = errors.New("Cached state not found")
	ErrStateInvalid    = errors.New("State does not match to the cached one")
	ErrCodeRequired    = errors.New("Code does not match requirements")
)

func HandleService(serviceName string, v *revel.Validation) error {
	v.Match(serviceName, enabledServices)

	if v.HasErrors() {
		return ErrServiceRequired
	}

	return nil
}

func HandleAuthorizeRequest(serviceName, sessionID string) string {
	state := issueState(serviceName, sessionID)
	service := supportedServices[serviceName]

	return service.AuthCodeURL(state)
}

func ValidateAuthorizeRequest(serviceName, sessionID, state, code string,
	v *revel.Validation) error {

	if err := validateState(serviceName, sessionID, state, v); err != nil {
		return err
	}

	if err := validateCode(code, v); err != nil {
		return err
	}

	return nil
}

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

func validateState(serviceName, sessionID, state string, v *revel.Validation) error {
	v.Required(state)
	v.Length(state, stateLength)

	if v.HasErrors() {
		return ErrStateRequired
	}

	cachedState, err := retriveState(serviceName, sessionID)
	if err != nil {
		return ErrStateNotFound
	}

	if eq := subtle.ConstantTimeCompare([]byte(state), []byte(cachedState)); eq != 1 {
		return ErrStateInvalid
	}

	return nil
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

func validateCode(code string, v *revel.Validation) error {
	v.Required(code)
	v.Length(code, codeLength)

	if v.HasErrors() {
		return ErrCodeRequired
	}

	return nil
}
