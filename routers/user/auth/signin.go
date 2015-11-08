package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/modules/auth"
	"github.com/Gr1N/pacman/modules/helpers"
	"github.com/Gr1N/pacman/modules/session"
)

const (
	signInTmpl = "user.auth.signin.tmpl"
)

type signInCompleteBinding struct {
	State string `form:"state" binding:"required,len=32"`
	Code  string `form:"code" binding:"required,len=20"`
}

// SignIn renders sign in page.
func SignIn(c *gin.Context) {
	c.HTML(http.StatusOK, signInTmpl, gin.H{})
}

// SignInPost starts authentication process.
func SignInPost(c *gin.Context) {
	service := c.Param("service")
	if err := auth.HandleService(service); err != nil {
		helpers.RedirectToSignIn(c)
		return
	}

	sessionID := session.ID(session.Get(c))
	redirectURL := auth.HandleAuthorizeRequest(service, sessionID)

	c.Redirect(http.StatusFound, redirectURL)
}

// SignInComplete finishes authentication process.
func SignInComplete(c *gin.Context) {
	service := c.Param("service")
	if err := auth.HandleService(service); err != nil {
		helpers.RedirectToSignIn(c)
		return
	}

	var b signInCompleteBinding
	if err := c.Bind(&b); err != nil {
		helpers.RedirectToSignIn(c)
		return
	}

	sessionObj := session.Get(c)
	sessionID := session.ID(sessionObj)
	if err := auth.ValidateAuthorizeRequest(service, sessionID, b.State); err != nil {
		helpers.RedirectToSignIn(c)
		return
	}

	user, err := auth.FinishAuthorizeRequest(service, b.Code)
	if err != nil {
		helpers.RedirectToSignIn(c)
		return
	}

	sessionObj.Clear()
	session.SetUserID(sessionObj, user.ID)

	helpers.RedirectToHome(c)
}
