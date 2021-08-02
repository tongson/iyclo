package main

import (
	"fmt"
	"os"

	clir "github.com/leaanthony/clir"
)

func main() {
	cli := clir.NewCli("iyclo", "Well-oiled containers", versionString)
	cli.SetBannerFunction(customBanner)
	var action actionT
	cli.BoolFlag("version", "Show version", &action.version)
	cli.BoolFlag("v", "Show version", &action.version)
	cli.StringFlag("log", "Path to JSON log", &action.log)
	cli.Action(handleCli(&action))
	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error encountered: %v\n", err)
		os.Exit(1)
	}
}
