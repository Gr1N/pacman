package controllers

import (
	"crypto/subtle"
	"fmt"
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
	servicesSupported    map[string]*oauth2.Config
	servicesCacheTimeout time.Duration
)

func init() {
	revel.OnAppStart(initAuth)
}

func initAuth() {
	servicesRaw, _ := revel.Config.String("auth.services")
	services := strings.Split(servicesRaw, ",")

	servicesAllowed = regexp.MustCompile(
		fmt.Sprintf("^(%v)$", strings.Join(services, "|")))
	servicesSupported = map[string]*oauth2.Config{
		"github": oauth2.GitHub,
	}

	servicesCacheTimeoutRaw, _ := revel.Config.String("auth.services.cache.timeout")
	servicesCacheTimeout, _ = time.ParseDuration(servicesCacheTimeoutRaw)
}

func (c Auth) Index() revel.Result {
	return c.Render()
}

func (c Auth) IndexEnd(service string) revel.Result {
	if !serviceValid(service, c.Validation) {
		return c.Redirect(Auth.Index)
	}

	return c.Render()
}

func (c Auth) Login(service string) revel.Result {
	if !serviceValid(service, c.Validation) {
		return c.Redirect(Auth.Index)
	}

	state := helpers.RandomString(32)

	key := serviceCacheKey(service, c.Session.Id())
	go cache.Set(key, state, servicesCacheTimeout)

	return c.Redirect(servicesSupported[service].AuthCodeUrl(state))
}

func (c Auth) LoginEnd(service string) revel.Result {
	// TODO: handle OAuth2.0 cancel response
	if !serviceValid(service, c.Validation) {
		return c.Redirect(Auth.Index)
	}

	var (
		state string
		code  string
	)

	c.Params.Bind(&state, "state")
	c.Params.Bind(&code, "code")

	stateValid := serviceStateValid(service, c.Session.Id(), state, c.Validation)
	codeValid := serviceCodeValid(service, code, c.Validation)
	if !stateValid || !codeValid {
		// TODO: show error?
		return c.Redirect(Auth.Index)
	}


	return c.Redirect(Auth.Index)
}

func serviceValid(service string, v *revel.Validation) bool {
	v.Match(service, servicesAllowed)

	if v.HasErrors() {
		revel.INFO.Printf("Got not supported service name (%s)", service)
		return false
	}

	return true
}

func serviceStateValid(service, sessionId, state string, v *revel.Validation) bool {
	v.Required(state)
	v.Length(state, 32)

	if v.HasErrors() {
		revel.INFO.Printf("Got invalid state (%s) value", state)
		return false
	}

	var cachedState string

	key := serviceCacheKey(service, sessionId)
	if err := cache.Get(key, &cachedState); err != nil {
		revel.INFO.Printf("Cached state value not found for key (%s)", key)
		return false
	}

	go cache.Delete(key)

	if eq := subtle.ConstantTimeCompare([]byte(state), []byte(cachedState)); eq != 1 {
		revel.INFO.Printf("State (%s) from request and cached state (%s) not equal",
			state, cachedState)
		return false
	}

	return true
}

func serviceCodeValid(service, code string, v *revel.Validation) bool {
	v.Required(code)
	v.Length(code, 20)

	if v.HasErrors() {
		revel.INFO.Printf("Got invalid code (%s) value", code)
		return false
	}

	return true
}

func serviceCacheKey(service, sessionId string) string {
	return strings.Join([]string{
		sessionId,
		service,
	}, ":")
}
