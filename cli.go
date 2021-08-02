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
	os.Exit(0)
}

func jLog(log string) zerolog.Logger {
	zerolog.TimeFieldFormat = time.RFC3339
	jsonFile, err := os.OpenFile(log, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
	return zerolog.New(jsonFile).With().Timestamp().Logger()
}

func handleCli(action *actionT) clir.Action {
	return func() error {
		if (*action).version {
			printVersion()
		}
		var log string
		if (*action).log == "" {
			log = defaultLog
		} else {
			log = (*action).log
		}
		jl := jLog(log)
		var db string
		if (*action).db == "" {
			db = defaultDb
		} else {
			db = (*action).db
		}
		// No need to test os.MkdirAll()
		os.MkdirAll(db, 0700)
		jl.Info().Msg("Starting up...")
		return nil
	}
}
