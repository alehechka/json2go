package gen

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/alehechka/json2go/jenshared"
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

	if len(c.TimeFormat) == 0 {
		return time.RFC3339
	}

	switch c.TimeFormat {
	case "Layout":
		return time.Layout
	case "ANSIC":
		return time.ANSIC
	case "UnixDate":
		return time.UnixDate
	case "RubyDate":
		return time.RubyDate
	case "RFC822":
		return time.RFC822
	case "RFC822Z":
		return time.RFC822Z
	case "RFC850":
		return time.RFC850
	case "RFC1123":
		return time.RFC1123
	case "RFC1123Z":
		return time.RFC1123Z
	case "RFC3339":
		return time.RFC3339
	case "RFC3339Nano":
		return time.RFC3339Nano
	case "Kitchen":
		return time.Kitchen
	case "Stamp":
		return time.Stamp
	case "StampMilli":
		return time.StampMilli
	case "StampMicro":
		return time.StampMicro
	case "StampNano":
		return time.StampNano
	default:
		return c.TimeFormat
	}
}
