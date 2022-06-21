package gen

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/alehechka/json2go/jenshared"
)

// Config presents Gen configurations.
type Config struct {
	Logger         *log.Logger
	URL            string
	File           string
	RootName       string
	PackageName    string
	OutputFileName string
}

func (c *Config) toJensharedConfig() *jenshared.Config {
	c.prepareOutputFileName()

	return &jenshared.Config{
		RootName:       c.RootName,
		PackageName:    c.PackageName,
		OutputFileName: c.OutputFileName,
	}
}

func (c *Config) prepareOutputFileName() {
	if len(c.File) > 0 && c.OutputFileName == DefaultOutputFile {
		filename := filepath.Base(c.File)
		filename = strings.TrimSuffix(filename, filepath.Ext(filename))

		badNames := map[string]bool{
			"..": true,
			".":  true,
			"\\": true,
			"/":  true,
			"":   true,
		}

		if _, exist := badNames[filename]; !exist {
			c.OutputFileName = filename
		}
	}

	if !strings.HasSuffix(c.OutputFileName, ".go") {
		c.OutputFileName += ".go"
	}
}
