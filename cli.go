package main

import (
	"embed"
	"fmt"
	"os"

	clir "github.com/leaanthony/clir"
)


//go:embed assets/*
var assetsSrc embed.FS

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
		defer file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			file.Close()
			os.Exit(1)
		}
		return nil
	}
}
