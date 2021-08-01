package main

import (
	"fmt"

	"github.com/leaanthony/clir"
)

const versionString = "0.1.0 (Chill Hazelnut)"

type actionsT struct {
	version bool
}

func main() {
	var actions actionsT
	cli := clir.NewCli("iyclo", "Well-oiled containers", versionString)
	cli.BoolFlag("version", "Show version", &actions.version)
	cli.Action(handleCli(&actions))
	if err := cli.Run(); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	}
}

func printVersion() {
	fmt.Printf("%s\n", versionString)
}

func handleCli(actions *actionsT) clir.Action {
	return func() error {
		if (*actions).version {
			printVersion()
		}
		return nil
	}
}
