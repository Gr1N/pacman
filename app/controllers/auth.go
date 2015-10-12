package controllers

import (
	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/routes"
)

type Auth struct {
	Base
}

func (c Auth) Logout() revel.Result {
	c.flushSession()

	return c.Redirect(routes.Application.Index())
}
