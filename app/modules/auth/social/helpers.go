package social

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/revel/revel/cache"

	"github.com/Gr1N/pacman/app/models"
	"github.com/Gr1N/pacman/app/modules/helpers"
)

func IssueState(service, sessionId string) string {
	state := helpers.RandomString(32)

	key := makeStateCacheKey(service, sessionId)
	go cache.Set(key, state, stateCacheTimeout)

	return state
}

func RetriveState(service, sessionId string) (string, bool) {
	var state string

	key := makeStateCacheKey(service, sessionId)
	if err := cache.Get(key, &state); err != nil {
		return "", false
	}

	go cache.Delete(key)

	return state, true
}

func makeStateCacheKey(service, sessionId string) string {
	return strings.Join([]string{
		sessionId,
		service,
	}, ":")
}

func FindOrCreateUser(service, code string, txn *gorm.DB) (*models.User, bool) {
	token := SupportedServices[service].Exchange(code)
	if token == nil {
		return nil, false
	}

	externalUser := SupportedServices[service].User(token)
	if externalUser == nil {
		return nil, false
	}

	if user, err := models.GetUserByService(txn, service, externalUser.Id); err == nil {
		return user, true
	}

	user := models.CreateUserByService(txn, service, externalUser.Id, externalUser.Name,
		externalUser.Email)
	return user, true
}
