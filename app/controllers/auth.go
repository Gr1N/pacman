package controllers

import (
	"regexp"
	"strings"
	"time"

	"github.com/revel/revel"
	"github.com/revel/revel/cache"

	"github.com/Gr1N/pacman/app/modules/helpers"
	"github.com/Gr1N/pacman/app/modules/oauth2"
)

type Auth struct {
	*revel.Controller
}

var (
	servicesAllowed      *regexp.Regexp
	servicesCacheTimeout time.Duration
)

func init() {
	revel.OnAppStart(initAuth)
}

func initAuth() {
	servicesRaw, _ := revel.Config.String("auth.services")
	services := strings.Split(servicesRaw, ",")
	servicesAllowed = regexp.MustCompile(
		helpers.JoinStrings("^(", strings.Join(services, "|"), ")$"))

	servicesCacheTimeoutRaw, _ := revel.Config.String("auth.services.cache.timeout")
	servicesCacheTimeout, _ = time.ParseDuration(servicesCacheTimeoutRaw)
}

func (c Auth) Index() revel.Result {
	return c.Render()
}

func (c Auth) Login(service string) revel.Result {
	c.Validation.Match(service, servicesAllowed)

	if c.Validation.HasErrors() {
		revel.INFO.Printf("Got not supported service name (%s)", service)
		return c.Redirect(Auth.Index)
	}

	services := map[string]*oauth2.Config{
		"github": oauth2.GitHub,
	}
	state := helpers.RandomString(32)

	cacheKey := helpers.JoinStrings(service, ":", state)
	go cache.Set(c.Session.Id(), cacheKey, servicesCacheTimeout)

	return c.Redirect(services[service].AuthCodeUrl(state))
}
