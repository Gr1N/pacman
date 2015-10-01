package controllers

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/modules/helpers"
	"github.com/Gr1N/pacman/app/modules/oauth2"
)

type Auth struct {
	*revel.Controller
}

var (
	servicesAllowed *regexp.Regexp
)

func init() {
	revel.OnAppStart(initAuth)
}

func initAuth() {
	var buf bytes.Buffer

	servicesRaw, _ := revel.Config.String("auth.services")
	services := strings.Split(servicesRaw, ",")

	buf.WriteString("^(")
	buf.WriteString(strings.Join(services, "|"))
	buf.WriteString(")$")

	servicesAllowed = regexp.MustCompile(buf.String())
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

	// TODO: Store state per service and user

	return c.Redirect(services[service].AuthCodeUrl(state))
}
