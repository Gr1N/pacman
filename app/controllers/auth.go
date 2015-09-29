package controllers

import (
	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/modules/oauth2"
)

type Auth struct {
	*revel.Controller
}

func (c Auth) Login() revel.Result {
	return c.Redirect(oauth2.GitHub.AuthCodeUrl("state"))
}
