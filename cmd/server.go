package cmd

import (
	"strings"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"

	"github.com/Gr1N/pacman/models"
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
	settings.Init()
	models.Init()

	g := initGin()

	routers.Init(g)

	port := strings.Join([]string{"", settings.S.Server.Port}, ":")
	g.Run(port)
}

func initGin() *gin.Engine {
	gin.SetMode(settings.S.RunMode)

	g := gin.New()

	g.Use(gin.Recovery(), gin.Logger())
	g.LoadHTMLGlob("templates/*")

	return g
}
