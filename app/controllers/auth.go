package controllers

import (
	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/modules/auth"
)

type Auth struct {
	*revel.Controller
}

func (c Auth) Index() revel.Result {
	return c.Render()
}

func (c Auth) IndexEnd(service string) revel.Result {
	if !auth.ServiceEnabled(service, c.Validation) {
		return c.Redirect(Auth.Index)
	}

	return c.Render()
}

func (c Auth) Login(service string) revel.Result {
	if !auth.ServiceEnabled(service, c.Validation) {
		return c.Redirect(Auth.Index)
	}

	state := auth.IssueState(service, c.Session.Id())
	return c.Redirect(auth.SupportedServices[service].AuthCodeUrl(state))
}

func (c Auth) LoginEnd(service string) revel.Result {
	// TODO: handle OAuth2.0 cancel response
	if !auth.ServiceEnabled(service, c.Validation) {
		return c.Redirect(Auth.Index)
	}

	var (
		state string
		code  string
	)

	c.Params.Bind(&state, "state")
	c.Params.Bind(&code, "code")

	stateValid := auth.StateValid(service, c.Session.Id(), state, c.Validation)
	codeValid := auth.CodeValid(service, code, c.Validation)
	if !stateValid || !codeValid {
		// TODO: show error?
		return c.Redirect(Auth.Index)
	}

	return c.Redirect(Auth.Index)
}
