package main

import (
	"fmt"
	"os"

	"github.com/leaanthony/clir"
)

const versionString = "0.1.0 (Chill Hazelnut)"

type actionT struct {
	version bool
}

func main() {
	var action actionT
	cli := clir.NewCli("iyclo", "Well-oiled containers", versionString)
	cli.SetBannerFunction(customBanner)
	cli.BoolFlag("version", "Show version", &action.version)
	cli.BoolFlag("v", "Show version", &action.version)
	cli.Action(handleCli(&action))
	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error encountered: %v\n", err)
	}
}

func customBanner(cli *clir.Cli) string {
	return `iyclo
` + fmt.Sprintf("%s\n%s", cli.ShortDescription(), cli.Version())
}

func printVersion() {
	fmt.Printf("%s\n", versionString)
}

func handleCli(action *actionT) clir.Action {
	return func() error {
		if (*action).version {
			printVersion()
		}
		return nil
	}
}
