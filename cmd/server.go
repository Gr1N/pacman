package cmd

import (
	"strings"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/models"
	"github.com/Gr1N/pacman/modules/auth"
	"github.com/Gr1N/pacman/modules/cache"
	"github.com/Gr1N/pacman/modules/logger"
	"github.com/Gr1N/pacman/modules/session"
	"github.com/Gr1N/pacman/modules/settings"
	"github.com/Gr1N/pacman/routers"
)

var CmdServer = cli.Command{
	Name:        "server",
	Usage:       "Start PMan web server",
	Description: "TBD",
	Action:      runServer,
}

func runServer(ctx *cli.Context) {
	preInit()

	gin.SetMode(settings.S.RunMode)

	g := gin.New()

	g.Use(gin.Recovery(), gin.Logger())
	g.LoadHTMLGlob("templates/*")

	postInit(g)

	port := strings.Join([]string{"", settings.S.Server.Port}, ":")
	g.Run(port)
}

func preInit() {
	settings.Init()
	logger.Init()
	cache.Init()
	models.Init()
	auth.Init()
}

func postInit(g *gin.Engine) {
	session.Init(g)
	routers.Init(g)
}
