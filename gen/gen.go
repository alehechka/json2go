package gen

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/alehechka/json2go/utils"
)

// Gen presents a generate tool for json2go.
type Gen struct {
	readSTDIN       func() (*bufio.Reader, error)
	downloadPayload func(url string)
}

// New creates a new Gen.
func New() *Gen {
	return &Gen{
		readSTDIN: utils.ReadSTDIN,
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

	reader, err := g.readSTDIN()
	if err != nil {
		config.Logger.Fatal(err.Error())
	}
	str, err := ioutil.ReadAll(reader)
	if err != nil {
		config.Logger.Fatal(err.Error())
	}
	fmt.Println(string(str))

	return nil
}
