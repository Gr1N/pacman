package controllers

import (
	"strconv"
	"strings"

	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/models"
	"github.com/Gr1N/pacman/app/routes"
)

type Base struct {
	*revel.Controller
}

type Any struct {
	Base
}

type NotAuthenticated struct {
	Base
}

type AnyAuthenticated struct {
	Base
}

type SessionAuthenticated struct {
	Base
}

type TokenAuthenticated struct {
	Base
}

func (c Base) attachUser() revel.Result {
	if user := c.getUser(); user != nil {
		return nil
	}

	if user := c.getUserFromSessionCookie(); user != nil {
		c.RenderArgs["user"] = user
		return nil
	}

	return nil
}

func (c Base) checkUser() revel.Result {
	if user := c.getUser(); user == nil {
		return c.Redirect(routes.AuthSocial.Index())
	}

	return nil
}

func (c Base) getUser() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}

	return nil
}

func (c Base) getUserFromSessionCookie() *models.User {
	if userID, ok := c.Session["user_id"]; ok {
		if id, err := strconv.ParseInt(userID, 10, 64); err == nil {
			if user, err := models.GetUserByID(id); err == nil {
				return user
			}
		}
	}

	return nil
}

func (c Base) getUserFromToken() *models.User {
	auth := strings.Split(c.Request.Header.Get("Authorization"), " ")
	if len(auth) != 2 || strings.ToLower(auth[0]) != "token" {
		return nil
	}

	token := auth[1]
	if user, err := models.GetUserByToken(token); err == nil {
		return user
	}

	return nil
}

func (c Base) flushSession() {
	for k := range c.Session {
		delete(c.Session, k)
	}
}

func (c Any) attachUser() revel.Result {
	return c.Base.attachUser()
}

func (c NotAuthenticated) attachUser() revel.Result {
	return c.Base.attachUser()
}

func (c NotAuthenticated) checkUser() revel.Result {
	if user := c.getUser(); user != nil {
		return c.Redirect(routes.Application.Index())
	}

	return nil
}

func (c AnyAuthenticated) attachUser() revel.Result {
	if user := c.getUser(); user != nil {
		return nil
	}

	if user := c.getUserFromSessionCookie(); user != nil {
		c.RenderArgs["user"] = user
		return nil
	}

	if user := c.getUserFromToken(); user != nil {
		c.RenderArgs["user"] = user
		return nil
	}

	return nil
}

func (c AnyAuthenticated) checkUser() revel.Result {
	return c.Base.checkUser()
}

func (c SessionAuthenticated) attachUser() revel.Result {
	return c.Base.attachUser()
}

func (c SessionAuthenticated) checkUser() revel.Result {
	return c.Base.checkUser()
}

func (c TokenAuthenticated) attachUser() revel.Result {
	if user := c.getUser(); user != nil {
		return nil
	}

	if user := c.getUserFromToken(); user != nil {
		c.RenderArgs["user"] = user
		return nil
	}

	return nil
}

func (c TokenAuthenticated) checkUser() revel.Result {
	return c.Base.checkUser()
}
