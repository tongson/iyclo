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

type httpT struct {
	logger    zerolog.Logger
	socket    net.Listener
	variables map[string]string
}

func customBanner(cli *clir.Cli) string {
	s, _ := assetsSrc.ReadFile("assets/iyclo.txt")
	return fmt.Sprintf("%s\n%s\n%s", string(s), cli.ShortDescription(), cli.Version())
}

func printVersion() {
	fmt.Printf("%s\n", versionStringG)
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
	return listener
}

func mainCli(action *actionT) clir.Action {
	return func() error {
		if (*action).version {
			printVersion()
		}

		var log string
		if (*action).log == "" {
			log = defaultLogG
		} else {
			log = (*action).log
		}

		var db string
		if (*action).db == "" {
			db = defaultDbG
		} else {
			db = (*action).db
		}
		// No need to test os.MkdirAll()
		os.MkdirAll(db, 0700)

		var socket string
		if (*action).socket == "" {
			socket = defaultSocketG
		} else {
			socket = (*action).socket
		}

		jl := jLog(log)
		l := unixSocket(socket)
		defer l.Close()
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)
		go func(c chan os.Signal) {
			sig := <-c
			jl.Debug().Msg("Shutting down...")
			fmt.Fprintf(os.Stderr, "Caught signal %s: shutting down.", sig)
			l.Close()
			os.Exit(0)
		}(sigc)

		var vars map[string]string
		vars = make(map[string]string)
		vars["db"] = (*action).db

		http := new(httpT)
		http.logger = jl
		http.socket = l
		http.variables = vars

		jl.Info().Msg("Starting up...")
		mainHttp(http)
		return nil
	}
}
