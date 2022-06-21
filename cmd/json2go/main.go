package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/alehechka/json2go"
	"github.com/alehechka/json2go/gen"
	"github.com/urfave/cli/v2"
)

const (
	urlFlag        = "url"
	fileFlag       = "file"
	rootFlag       = "root"
	packageFlag    = "package"
	outputFileFlag = "output"
	debugFlag      = "debug"
	quietFlag      = "quiet"
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
		Value:   gen.DefaultPackage,
	},
	&cli.StringFlag{
		Name:    outputFileFlag,
		Aliases: []string{"o", "out"},
		Usage: `The name of the file that is generated. If a file is provided as input, will use matching name unless explicitly provided. 
		The ".go" extension is not required and will be automatically appended.`,
		Value: gen.DefaultOutputFile,
	},
	&cli.BoolFlag{
		Name:  debugFlag,
		Usage: "Log debug messages.",
	},
	&cli.BoolFlag{
		Name:    quietFlag,
		Aliases: []string{"q"},
		Usage:   "Quiets fatal errors.",
	},
}

func generateTypes(ctx *cli.Context) error {
	debugger := log.New(ioutil.Discard, "", log.LstdFlags)

	if ctx.Bool(debugFlag) {
		debugger = log.New(log.Writer(), "", 0)
	}

	err := gen.New().Build(&gen.Config{
		Debugger:       debugger,
		URL:            ctx.String(urlFlag),
		File:           ctx.String(fileFlag),
		RootName:       ctx.String(rootFlag),
		PackageName:    ctx.String(packageFlag),
		OutputFileName: ctx.String(outputFileFlag),
	})

	if err != nil && ctx.Bool(quietFlag) {
		log.Println(err)
		return nil
	}

	return err
}

func main() {
	app := cli.NewApp()
	app.Version = json2go.Version
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
