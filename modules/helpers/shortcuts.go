package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RedirectToHome is shortcut for redirecting to home page.
func RedirectToHome(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}

// RedirectToSignIn is shortcut for redirecting to sign in page.
func RedirectToSignIn(c *gin.Context) {
	c.Redirect(http.StatusFound, "/user/signin")
}
