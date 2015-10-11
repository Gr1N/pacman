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
		return &models.User{}, false
	}

	externalUser := SupportedServices[service].User(token)
	if externalUser == nil {
		return &models.User{}, false
	}

	var (
		user        models.User
		userService models.Service
	)

	txn.Where(&models.Service{
		Name:          service,
		UserServiceId: externalUser.Id,
	}).First(&userService)

	if userService.Id == 0 {
		user = models.User{
			Services: []models.Service{{
				Name:             service,
				UserServiceId:    externalUser.Id,
				UserServiceName:  externalUser.Name,
				UserServiceEmail: externalUser.Email,
			}},
		}
		txn.Create(&user)
	} else {
		txn.Model(&userService).Related(&user)
	}

	return &user, true
}
