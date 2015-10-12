package controllers

import (
	"strconv"

	"github.com/revel/revel"

	auths "github.com/Gr1N/pacman/app/modules/auth/social"
	"github.com/Gr1N/pacman/app/routes"
)

type AuthSocial struct {
	Base
}

func (c AuthSocial) checkAuthentication() revel.Result {
	if user := c.withUser(); user != nil {
		return c.Redirect(routes.Application.Index())
	}

	return nil
}

func (c AuthSocial) Index() revel.Result {
	return c.Render()
}

func (c AuthSocial) IndexEnd(service string) revel.Result {
	if !auths.ServiceEnabled(service, c.Validation) {
		return c.Redirect(routes.AuthSocial.Index())
	}

	return c.Render()
}

func (c AuthSocial) Login(service string) revel.Result {
	if !auths.ServiceEnabled(service, c.Validation) {
		return c.Redirect(routes.AuthSocial.Index())
	}

	state := auths.IssueState(service, c.Session.Id())
	return c.Redirect(auths.SupportedServices[service].AuthCodeUrl(state))
}

func (c AuthSocial) LoginEnd(service string) revel.Result {
	// TODO: handle OAuth2.0 cancel response
	if !auths.ServiceEnabled(service, c.Validation) {
		return c.Redirect(routes.AuthSocial.Index())
	}

	var (
		state string
		code  string
	)

	c.Params.Bind(&state, "state")
	c.Params.Bind(&code, "code")

	stateValid := auths.StateValid(service, c.Session.Id(), state, c.Validation)
	codeValid := auths.CodeValid(service, code, c.Validation)
	if !stateValid || !codeValid {
		// TODO: handle error
		return c.Redirect(routes.AuthSocial.Index())
	}

	user, found := auths.FindOrCreateUser(service, code, c.Txn)
	if !found {
		// TODO: handle error
		return c.Redirect(routes.AuthSocial.Index())
	}

	c.flushSession()

	c.Session["user_id"] = strconv.FormatInt(user.Id, 10)

	return c.Redirect(routes.Application.Index())
}
