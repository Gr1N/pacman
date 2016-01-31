package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/modules/middleware"
)

// Home returns context for front-end initialization.
func Home(c *gin.Context) {
	user := middleware.UserFromContext(c)

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id": user.ID,
		},
	})
}
