package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	urlFlag        = "url"
	fileFlag       = "file"
	rootFlag       = "root"
	packageFlag    = "package"
	outputFileFlag = "output"
)

var generateFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    urlFlag,
		Aliases: []string{"u"},
		Usage:   "Download JSON payload from URL",
	},
	&cli.StringFlag{
		Name:    fileFlag,
		Aliases: []string{"f"},
		Usage:   "Read JSON payload from file on disk",
	},
	&cli.StringFlag{
		Name:    rootFlag,
		Aliases: []string{"r"},
		Usage:   "The name of the root object or array",
		Value:   "Root",
	},
	&cli.StringFlag{
		Name:    packageFlag,
		Aliases: []string{"p"},
		Usage:   "The name of the package to generate the types under.",
		Value:   "main",
	},
	&cli.StringFlag{
		Name:    outputFileFlag,
		Aliases: []string{"o", "out"},
		Usage: `The name of the file that is generated. If a file is provided as input, will use matching name unless explicitly provided. 
		The ".go" extension is not required and will be automatically appended.`,
		Value: "types.go",
	},
}

func generateTypes(ctx *cli.Context) error {
	fmt.Println("URL:", ctx.String(urlFlag))
	return nil
}

func main() {
	app := cli.NewApp()
	app.Version = "v0.1.0"
	app.Usage = "Automatically generate deeply nested Go types from a JSON payload."
	app.Commands = []*cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "Generate Go Types",
			Action:  generateTypes,
			Flags:   generateFlags,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
