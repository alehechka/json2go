package jenshared

import (
	"fmt"

	"github.com/alehechka/json2go/utils"
	"github.com/dave/jennifer/jen"
)

// GenerateTypes parses the provided JSON payload and generates matching Go types.
func GenerateTypes(data interface{}, config *Config) error {
	f, err := generateTypes(data, config)
	if err != nil {
		return err
	}

	if err := utils.CreateFilePath(config.OutputDirectory); err != nil {
		return err
	}

	return f.Save(fmt.Sprintf("%s%s", config.OutputDirectory, config.OutputFileName))
}

func generateTypes(data interface{}, config *Config) (*jen.File, error) {
	f := jen.NewFile(config.PackageName)

	return f, nil
}
