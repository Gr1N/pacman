package main

import (
	"os"

	"github.com/codegangsta/cli"

	"github.com/Gr1N/pacman/cmd"
)

const (
	Version = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = "PMan"
	app.Usage = "TBD"
	app.Version = Version
	app.Commands = []cli.Command{
		cmd.CmdServer,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
