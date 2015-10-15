package controllers

import (
	"strconv"

	"github.com/revel/revel"

	auths "github.com/Gr1N/pacman/app/modules/auth/social"
	"github.com/Gr1N/pacman/app/routes"
)

type AuthSocial struct {
	NotAuthenticated
}

func (c AuthSocial) Index() revel.Result {
	return c.Render()
}

func (c AuthSocial) IndexEnd(service string) revel.Result {
	if err := auths.HandleService(service, c.Validation); err != nil {
		return c.Redirect(routes.AuthSocial.Index())
	}

	return c.Render()
}

func (c AuthSocial) Login(service string) revel.Result {
	if err := auths.HandleService(service, c.Validation); err != nil {
		return c.Redirect(routes.AuthSocial.Index())
	}

	redirectUrl := auths.HandleAuthorizeRequest(service, c.Session.Id())
	return c.Redirect(redirectUrl)
}

func (c AuthSocial) LoginEnd(service string) revel.Result {
	// TODO: handle OAuth2.0 cancel response
	if err := auths.HandleService(service, c.Validation); err != nil {
		return c.Redirect(routes.AuthSocial.Index())
	}

	var (
		state string
		code  string
	)

	c.Params.Bind(&state, "state")
	c.Params.Bind(&code, "code")

	if err := auths.ValidateAuthorizeRequest(service, c.Session.Id(),
		state, code, c.Validation); err != nil {
		// TODO: handle error
		return c.Redirect(routes.AuthSocial.Index())
	}

	user, err := auths.FinishAuthorizeRequest(service, code)
	if err != nil {
		// TODO: handle error
		return c.Redirect(routes.AuthSocial.Index())
	}

	c.flushSession()

	c.Session["user_id"] = strconv.FormatInt(user.Id, 10)

	return c.Redirect(routes.Application.Index())
}
