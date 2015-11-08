package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/models"
	"github.com/Gr1N/pacman/modules/helpers"
	"github.com/Gr1N/pacman/modules/session"
)

const (
	// ContextUserKey represents key which used for storing user
	// object in context.
	ContextUserKey = "user"
)

// UserFromCookie reads cookie, tries to find user in database
// and if user found attaches user object to the context.
func UserFromCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		if userID, exists := session.UserID(session.Get(c)); exists {
			if user, err := models.GetUserByID(userID); err == nil {
				c.Set(ContextUserKey, user)
			}
		}
	}
}

// Authenticated checks is user object attached to the context
// and if user object not attached redirects to the sign in page.
func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exists := c.Get(ContextUserKey); !exists {
			helpers.RedirectToSignIn(c)
			c.Abort()
		}
	}
}

// NotAuthenticated checks is user object attached to the context
// and if user object attached redirects to the home page.
func NotAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exists := c.Get(ContextUserKey); exists {
			helpers.RedirectToHome(c)
			c.Abort()
		}
	}
}
