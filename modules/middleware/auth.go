package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/models"
)

const (
	// ContextUserKey represents key which used for storing user
	// object in context.
	ContextUserKey = "user"
)

// UserFromToken reads `Authorization` header, tries to find user in database
// and if user found attaches user object to the context.
func UserFromToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if values, _ := c.Request.Header["Authorization"]; len(values) > 0 {
			auth := strings.Split(values[0], " ")
			if len(auth) == 2 && strings.ToLower(auth[0]) == "token" {
				token := auth[1]
				if user, err := models.GetUserByToken(token); err == nil {
					c.Set(ContextUserKey, user)
				}
			}
		}
	}
}

// UserFromContext returns the user object which attached to the context.
func UserFromContext(c *gin.Context) *models.User {
	user, _ := c.Get(ContextUserKey)

	return user.(*models.User)
}

// Authenticated checks is user object attached to the context
// and if user object not attached redirects to the sign in page.
func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exists := c.Get(ContextUserKey); !exists {
			c.JSON(http.StatusForbidden, gin.H{})
			c.Abort()
		}
	}
}

// NotAuthenticated checks is user object attached to the context
// and if user object attached redirects to the home page.
func NotAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exists := c.Get(ContextUserKey); exists {
			c.JSON(http.StatusForbidden, gin.H{})
			c.Abort()
		}
	}
}
