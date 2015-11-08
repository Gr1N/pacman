package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/modules/helpers"
	"github.com/Gr1N/pacman/modules/session"
)

// SignOut flushes user's session.
func SignOut(c *gin.Context) {
	sessionObj := session.Get(c)
	sessionObj.Clear()
	sessionObj.Save()

	helpers.RedirectToHome(c)
}
