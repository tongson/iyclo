package main

import (
	"os"

	clir "github.com/leaanthony/clir"
)

func main() {
	cli := clir.NewCli("iyclo", "Well-oiled containers", versionStringG)
	cli.SetBannerFunction(customBanner)
	var action actionT
	cli.BoolFlag("V", "Show version", &action.version)
	cli.StringFlag("log", "Path to JSON log", &action.log)
	cli.StringFlag("db", "Path to state directory", &action.db)
	cli.StringFlag("socket", "Path to the Unix socket", &action.socket)
	cli.Action(mainCli(&action))
	if err := cli.Run(); err != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
