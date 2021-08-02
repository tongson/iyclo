package main

import (
	"embed"
	"fmt"
	"os"
	"time"

	clir "github.com/leaanthony/clir"
	zerolog "github.com/rs/zerolog"
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
		zerolog.TimeFieldFormat = time.RFC3339
		jsonFile, err := os.OpenFile(log, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
		defer jsonFile.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			jsonFile.Close()
			os.Exit(1)
		}
		var jl zerolog.Logger
		jl = zerolog.New(jsonFile).With().Timestamp().Logger()
		jl.Info().Msg("Starting up...")
		return nil
	}
}
