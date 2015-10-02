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

func (c Auth) IndexEnd(service string) revel.Result {
	if !serviceAllowed(service, c.Validation) {
		return c.Redirect(Auth.Index)
	}

	return c.Render()
}

func (c Auth) Login(service string) revel.Result {
	if !serviceAllowed(service, c.Validation) {
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

func serviceAllowed(service string, v *revel.Validation) bool {
	v.Match(service, servicesAllowed)

	if v.HasErrors() {
		revel.INFO.Printf("Got not supported service name (%s)", service)
		return false
	}

	return true
}
