package auth

import (
	"strings"

	"github.com/revel/revel/cache"

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
