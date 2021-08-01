package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/leaanthony/clir"
)

const versionString = "0.1.0 (Chill Hazelnut)"

//go:embed assets/*
var assetsSrc embed.FS

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
	}
}

func customBanner(cli *clir.Cli) string {
	s, _ := assetsSrc.ReadFile("assets/iyclo.txt")
	return fmt.Sprintf("%s\n%s\n%s", string(s), cli.ShortDescription(), cli.Version())
}

func printVersion() {
	fmt.Printf("%s\n", versionString)
}

func handleCli(action *actionT) clir.Action {
	return func() error {
		if (*action).version {
			printVersion()
		}
		var log string
		if (*action).log == "" {
			log = "/var/log/iyclo.json"
		} else {
			log = (*action).log
		}
		file, err := os.Create(log)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		}
		defer file.Close()
		return nil
	}
}
