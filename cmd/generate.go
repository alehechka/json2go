package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/alehechka/json2go/gen"
	"github.com/urfave/cli/v2"
)

const (
	urlFlag          = "url"
	fileFlag         = "file"
	rootFlag         = "root"
	packageFlag      = "package"
	outputFileFlag   = "output"
	debugFlag        = "debug"
	quietFlag        = "quiet"
	stdoutFlag       = "out"
	timeFormatFlag   = "time"
	omitEmptyFlag    = "omitempty"
	alphabeticalFlag = "alphabetical"
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
		Value:   gen.DefaultRootName,
	},
	&cli.StringFlag{
		Name:    packageFlag,
		Aliases: []string{"p"},
		Usage:   "The name of the package to generate the types under.",
		Value:   gen.DefaultPackage,
	},
	&cli.StringFlag{
		Name:    outputFileFlag,
		Aliases: []string{"o"},
		Usage: `The name of the file that is generated. If a file is provided as input, will use matching name unless explicitly provided. 
		The ".go" extension is not required and will be automatically appended.`,
		Value: gen.DefaultOutputFile,
	},
	&cli.StringFlag{
		Name:    timeFormatFlag,
		Aliases: []string{"t"},
		Usage:   "Time format to use while parsing strings for potential time.Time variables. View time.Time constants for possible defaults: https://pkg.go.dev/time#pkg-constants",
		Value:   time.RFC3339,
	},
	&cli.BoolFlag{
		Name:  omitEmptyFlag,
		Usage: "Appends the omitempty to all object variable tags.",
	},
	&cli.BoolFlag{
		Name:    alphabeticalFlag,
		Aliases: []string{"a", "alpha"},
		Usage:   "Sorts all keys into alphabetical order before generation.",
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
	&cli.BoolFlag{
		Name:  stdoutFlag,
		Usage: "Print Go structs to STDOUT instead of saving to file.",
	},
}

func generateTypes(ctx *cli.Context) (err error) {
	var debugger *log.Logger

	if ctx.Bool(debugFlag) {
		debugger = log.New(log.Writer(), "", 0)
	}

	config := &gen.Config{
		Debugger:       debugger,
		URL:            ctx.String(urlFlag),
		File:           ctx.String(fileFlag),
		RootName:       ctx.String(rootFlag),
		PackageName:    ctx.String(packageFlag),
		OutputFileName: ctx.String(outputFileFlag),
		TimeFormat:     ctx.String(timeFormatFlag),
		OmitEmpty:      ctx.Bool(omitEmptyFlag),
		Alphabetical:   ctx.Bool(alphabeticalFlag),
	}

	if ctx.Bool(stdoutFlag) {
		var out string
		out, err = gen.New().Build(config)
		fmt.Println(out)
	} else {
		err = gen.New().Generate(config)
	}

	if err != nil && ctx.Bool(quietFlag) {
		log.Println(err.Error())
		return nil
	}

	return err
}

// GenerateCommand provides the config for the "generate" CLI command
var GenerateCommand = &cli.Command{
	Name:    "generate",
	Aliases: []string{"g"},
	Usage:   "Generate Go Types",
	Action:  generateTypes,
	Flags:   generateFlags,
}
