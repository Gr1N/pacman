package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/routers/health"
)

func Init(g *gin.Engine) {
	// Health check group.
	h := g.Group("/h")
	{
		h.GET("/ping", health.Ping)
	}
}
