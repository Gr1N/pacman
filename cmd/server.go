package cmd

import (
	"strings"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/models"
	"github.com/Gr1N/pacman/modules/cache"
	"github.com/Gr1N/pacman/modules/oauth2"
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

	postInit(g)

	port := strings.Join([]string{"", settings.S.Server.Port}, ":")
	g.Run(port)
}

func preInit() {
	settings.Init()
	cache.Init()
	models.Init()
	oauth2.Init()
}

func postInit(g *gin.Engine) {
	routers.Init(g)
}
