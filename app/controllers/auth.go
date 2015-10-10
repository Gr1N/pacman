package controllers

import (
	"strconv"

	"github.com/revel/revel"

	"github.com/Gr1N/pacman/app/modules/auth"
)

type Auth struct {
	Application
}

type Authd struct {
	Application
}

func (c Auth) checkAuthentication() revel.Result {
	if user := c.withUser(); user != nil {
		return c.Redirect(Application.Index)
	}

	return nil
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
		// TODO: handle error
		return c.Redirect(Auth.Index)
	}

	user, found := auth.FindOrCreateUserUsingService(service, code, c.Txn)
	if !found {
		// TODO: handle error
		return c.Redirect(Auth.Index)
	}

	flushSession(c.Session)
	c.Session["user_id"] = strconv.FormatInt(user.Id, 10)

	return c.Redirect(Application.Index)
}

func (c Authd) Logout() revel.Result {
	flushSession(c.Session)
	return c.Redirect(Application.Index)
}

func flushSession(session revel.Session) {
	for k := range session {
		delete(session, k)
	}
}
