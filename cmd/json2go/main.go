package main

import (
	"log"
	"os"

	"github.com/alehechka/json2go"
	"github.com/alehechka/json2go/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Version = json2go.Version
	app.Usage = "Automatically generate deeply nested Go types from a JSON payload."
	app.Commands = []*cli.Command{
		cmd.GenerateCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
