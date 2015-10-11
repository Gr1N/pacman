package controllers

import (
	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/routes"
)

type Auth struct {
	Application
}

func (c Auth) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}

	return c.Redirect(routes.Application.Index())
}
