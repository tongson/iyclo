package main

import (
	"fmt"
	"os"

	"github.com/leaanthony/clir"
)

const versionString = "0.1.0 (Chill Hazelnut)"

type actionT struct {
	version bool
	log string
}

func main() {
	var action actionT
	cli := clir.NewCli("iyclo", "Well-oiled containers", versionString)
	cli.SetBannerFunction(customBanner)
	cli.BoolFlag("version", "Show version", &action.version)
	cli.BoolFlag("v", "Show version", &action.version)
	cli.StringFlag("log", "Path to JSON log", &action.log)
	cli.Action(handleCli(&action))
	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error encountered: %v\n", err)
		os.Exit(1)
	}
}
