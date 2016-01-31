package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/models"
	"github.com/Gr1N/pacman/modules/middleware"
)

// Home returns context for front-end initialization.
func Home(c *gin.Context) {
	userR, _ := c.Get(middleware.ContextUserKey)
	user := userR.(*models.User)

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id": user.ID,
		},
	})
}
