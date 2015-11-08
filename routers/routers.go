package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/modules/middleware"
	"github.com/Gr1N/pacman/routers/health"
	"github.com/Gr1N/pacman/routers/user/auth"
)

// Init initializes application routers.
func Init(g *gin.Engine) {
	g.Use(middleware.UserFromCookie())

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
			usin.GET("/", auth.SignIn)
			usin.POST("/:service", auth.SignInPost)
			usin.GET("/:service/complete", auth.SignInComplete)
		}

		usout := u.Group("/signout")
		usout.Use(middleware.Authenticated())
		{
			usout.POST("/", auth.SignOut)
		}
	}
}
