package controllers

import (
	"regexp"

	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/modules/oauth2"
)

type Auth struct {
	*revel.Controller
}

func (c Auth) Index() revel.Result {
	return c.Render()
}

func (c Auth) Login(service string) revel.Result {
	c.Validation.Match(service, regexp.MustCompile("^(github)$"))

	if c.Validation.HasErrors() {
		revel.INFO.Printf("Got not supported service name (%s)", service)
		return c.Redirect(Auth.Index)
	}

	services := map[string]*oauth2.Config{
		"github": oauth2.GitHub,
	}

	// TODO: Generate and store state per service and user

	return c.Redirect(services[service].AuthCodeUrl("state"))
}
