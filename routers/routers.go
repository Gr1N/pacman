package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/modules/middleware"
	"github.com/Gr1N/pacman/routers/health"
	"github.com/Gr1N/pacman/routers/user/auth"
)

// Init initializes application routers.
func Init(g *gin.Engine) {
	g.Use(middleware.UserFromToken())

	// Home page.
	g.GET("/", Home)

	// Health check group.
	h := g.Group("/h")
	{
		h.GET("/ping", health.Ping)
	}

	// User related group.
	u := g.Group("/user")
	{
		usin := u.Group("/signin")
		usin.Use(middleware.NotAuthenticated())
		{
			usin.POST("/:service", auth.SignIn)
			usin.GET("/:service/complete", auth.SignInComplete)
		}
	}
}
