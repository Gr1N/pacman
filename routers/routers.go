package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/routers/health"
)

func Init(g *gin.Engine) {
	h := g.Group("/__health")
	{
		h.GET("/ping", health.Ping)
	}
}
