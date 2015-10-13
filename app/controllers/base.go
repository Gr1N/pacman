package controllers

import (
	"strconv"

	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/models"
)

type Base struct {
	*revel.Controller
}

func (c Base) tryAuthenticate() revel.Result {
	if user := c.withUser(); user != nil {
		c.RenderArgs["user"] = user
	}

	return nil
}

func (c Base) withUser() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}

	if userId, ok := c.Session["user_id"]; ok {
		if id, err := strconv.ParseInt(userId, 10, 64); err == nil {
			if user, err := models.GetUserById(id); err == nil {
				return user
			}
		}
	}

	return nil
}

func (c Base) flushSession() {
	for k := range c.Session {
		delete(c.Session, k)
	}
}
