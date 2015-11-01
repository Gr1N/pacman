package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/routers/health"
)

func Init(g *gin.Engine) {
	// Home page.
	g.GET("/", Home)

	// Health check group.
	h := g.Group("/h")
	{
		h.GET("/ping", health.Ping)
	}
}
