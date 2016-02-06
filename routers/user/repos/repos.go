package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/models"
	"github.com/Gr1N/pacman/modules/middleware"
	"github.com/Gr1N/pacman/modules/oauth2"
)

// ReposList returns list of user repos for specified service.
func ReposList(c *gin.Context) {
	service := c.Param("service")
	if err := oauth2.HandleService(service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	user := middleware.UserFromContext(c)
	repos, _ := models.GetUserReposByService(user.ID, service)

	c.JSON(http.StatusOK, gin.H{
		"list":  repos,
		"total": len(repos),
	})
}

// ReposUpdate updates list of user repos for specified service.
func ReposUpdate(c *gin.Context) {
	// Not implemented yet
}
