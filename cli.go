package main

import (
	"embed"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
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

func unixSocket(path string) net.Listener {
	listener, err := net.Listen("unix", path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to initialize socket: %v\n", err.Error())
		os.Exit(1)
	}
	os.Chmod(path, 0600)
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func(c chan os.Signal) {
		sig := <-c
		fmt.Fprintf(os.Stderr, "Caught signal %s: shutting down.", sig)
		listener.Close()
		os.Exit(255)
	}(sigc)

	return listener
}

func mainCli(action *actionT) clir.Action {
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

		var socket string
		if (*action).socket == "" {
			socket = defaultSocket
		} else {
			socket = (*action).socket
		}
		l := unixSocket(socket)
		defer l.Close()
		mainHttp(l)

		jl.Info().Msg("Starting up...")
		return nil
	}
}
