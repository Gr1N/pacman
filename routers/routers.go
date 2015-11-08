package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/routers/health"
	"github.com/Gr1N/pacman/routers/user/auth"
)

// Init initializes application routers.
func Init(g *gin.Engine) {
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
		sin := u.Group("/signin")
		{
			sin.GET("/", auth.SignIn)
			sin.POST("/:service", auth.SignInPost)
			sin.GET("/:service/complete", auth.SignInComplete)
		}
	}
}
