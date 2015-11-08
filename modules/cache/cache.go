package cache

import (
	"time"

	"github.com/gin-gonic/contrib/cache"

	"github.com/Gr1N/pacman/modules/settings"
)

var (
	c cache.CacheStore
)

// Init initializes application cache.
func Init() {
	caches := settings.S.Cache

	expiration, _ := time.ParseDuration(caches.DefaultExpiration)
	c = cache.NewRedisCache(
		caches.Host,
		caches.Password,
		expiration)
}

func Set(key string, value interface{}, expires time.Duration) error {
	return c.Set(key, value, expires)
}

func Get(key string, ptrValue interface{}) error {
	return c.Get(key, ptrValue)
}

func Delete(key string) error {
	return c.Delete(key)
}
