package gen

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/alehechka/json2go/jenshared"
	"github.com/alehechka/json2go/utils"
)

// Gen presents a generate tool for json2go.
type Gen struct {
	readSTDIN         func() ([]byte, error)
	downloadPayload   func(url string) ([]byte, error)
	readFile          func(filepath string) ([]byte, error)
	decodeJSON        func(data []byte, v any) error
	generateTypesFile func(data interface{}, config *jenshared.Config) error
	generateTypes     func(data interface{}, config *jenshared.Config) string
	jsonPayload       interface{}
	bytes             []byte
}

// New creates a new Gen.
func New() *Gen {
	return &Gen{
		readSTDIN:         utils.ReadSTDIN,
		downloadPayload:   utils.DownloadPayload,
		readFile:          os.ReadFile,
		decodeJSON:        json.Unmarshal,
		generateTypesFile: jenshared.GenerateTypesFile,
		generateTypes:     jenshared.GenerateTypes,
	}
}

// Generate builds the type structs Go file.
func (g *Gen) Generate(config *Config) error {
	if config.Debugger == nil {
		config.Debugger = log.New(ioutil.Discard, "", log.LstdFlags)
	}

	err := g.prepareJSON(config)
	if err != nil {
		return err
	}

	return g.generateTypesFile(g.jsonPayload, config.toJensharedConfig())
}

// Build builds the type structs Go payload.
func (g *Gen) Build(config *Config) (string, error) {
	if config.Debugger == nil {
		config.Debugger = log.New(ioutil.Discard, "", log.LstdFlags)
	}

	err := g.prepareJSON(config)
	if err != nil {
		return "", err
	}

	return g.generateTypes(g.jsonPayload, config.toJensharedConfig()), nil
}

func (g *Gen) prepareJSON(config *Config) (err error) {
	if len(config.File) > 0 {
		config.Debugger.Printf("Reading file: %s\n", config.File)
		if g.bytes, err = g.readFile(config.File); err != nil {
			config.Debugger.Printf("Failed to read file: %s\n", config.File)
		}

	}

	if len(g.bytes) == 0 && len(config.URL) > 0 {
		config.Debugger.Printf("Downloading data from: %s\n", config.URL)
		if g.bytes, err = g.downloadPayload(config.URL); err != nil {
			config.Debugger.Printf("Failed to download data from: %s\n", config.URL)
		}
	}

	if len(g.bytes) == 0 {
		config.Debugger.Println("Reading data from STDIN")
		if g.bytes, err = g.readSTDIN(); err != nil {
			config.Debugger.Println("Failed to read data from STDIN")
		}
	}

	if len(g.bytes) == 0 {
		return errors.New("no JSON payload provided")
	}

	if err := g.decodeJSON(g.bytes, &g.jsonPayload); err != nil {
		return err
	}

	return nil
}
