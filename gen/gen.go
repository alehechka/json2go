package gen

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/alehechka/json2go/utils"
)

// Gen presents a generate tool for json2go.
type Gen struct {
	readSTDIN       func() ([]byte, error)
	downloadPayload func(url string) ([]byte, error)
	readFile        func(filepath string) ([]byte, error)
	decodeJSON      func(data []byte, v any) error
	jsonPayload     interface{}
	bytes           []byte
}

// New creates a new Gen.
func New() *Gen {
	return &Gen{
		readSTDIN:       utils.ReadSTDIN,
		downloadPayload: utils.DownloadPayload,
		readFile:        os.ReadFile,
		decodeJSON:      json.Unmarshal,
	}
}

// Config presents Gen configurations.
type Config struct {
	Logger         *log.Logger
	URL            string
	File           string
	RootName       string
	PackageName    string
	OutputFileName string
}

// Build builds the type structs go file.
func (g *Gen) Build(config *Config) error {

	err := g.prepareJSON(config)
	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", g.jsonPayload)

	return nil
}

func (g *Gen) prepareJSON(config *Config) error {
	if len(config.File) > 0 {
		g.bytes, _ = g.readFile(config.File)
	}

	if len(g.bytes) == 0 && len(config.URL) > 0 {
		g.bytes, _ = g.downloadPayload(config.URL)
	}

	if len(g.bytes) == 0 {
		g.bytes, _ = g.readSTDIN()
	}

	if len(g.bytes) == 0 {
		return errors.New("no JSON payload provided")
	}

	if err := g.decodeJSON(g.bytes, &g.jsonPayload); err != nil {
		return err
	}

	return nil
}
