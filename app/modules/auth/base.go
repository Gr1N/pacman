package auth

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/modules/oauth2"
)

var (
	SupportedServices map[string]*oauth2.Config

	enabledServices   *regexp.Regexp
	stateCacheTimeout time.Duration
)

func init() {
	revel.OnAppStart(initAuth)
}

func initAuth() {
	SupportedServices = map[string]*oauth2.Config{
		"github": oauth2.GitHub,
	}

	services := strings.Split(
		revel.Config.StringDefault("auth.services", ""), ",")
	enabledServices = regexp.MustCompile(
		fmt.Sprintf("^(%v)$", strings.Join(services, "|")))

	stateCacheTimeout, _ = time.ParseDuration(
		revel.Config.StringDefault("auth.services.cache.timeout", "10m"))
}
