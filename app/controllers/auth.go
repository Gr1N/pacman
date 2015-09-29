package controllers

import (
	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/modules"
)

type Auth struct {
	*revel.Controller
}

func (c Auth) Login() revel.Result {
	return c.Redirect(modules.GitHub.AuthCodeUrl("state"))
}
