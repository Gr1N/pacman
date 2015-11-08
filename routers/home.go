package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/modules/middleware"
)

// Home renders home page.
func Home(c *gin.Context) {
	user, _ := c.Get(middleware.ContextUserKey)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"user": user,
	})
}
