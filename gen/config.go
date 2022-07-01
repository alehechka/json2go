package gen

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/alehechka/json2go/jenshared"
	"github.com/alehechka/json2go/utils"
)

// Config presents Gen configurations.
type Config struct {
	Debugger       *log.Logger
	URL            string
	File           string
	RootName       string
	PackageName    string
	OutputFileName string
	TimeFormat     string
}

func (c *Config) toJensharedConfig() *jenshared.Config {
	c.prepareOutputFileName()

	var dir string
	if c.PackageName != DefaultPackage {
		dir = fmt.Sprintf("%s/", c.PackageName)
	}

	c.PackageName = filepath.Base(c.PackageName)

	return &jenshared.Config{
		RootName:        c.RootName,
		PackageName:     c.PackageName,
		OutputFileName:  c.OutputFileName,
		OutputDirectory: dir,
		TimeFormat:      c.getTimeFormat(),
		Debugger:        c.Debugger,
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

func (c *Config) getTimeFormat() string {
	return utils.GetTimeFormat(c.TimeFormat, time.RFC3339)
}
