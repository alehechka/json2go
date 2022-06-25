package jenshared

import (
	"fmt"

	"github.com/alehechka/json2go/utils"
	"github.com/dave/jennifer/jen"
)

// GenerateTypesFile parses the provided JSON payload and generates matching Go types file.
func GenerateTypesFile(data interface{}, config *Config) error {
	f := generateTypes(data, config)

	if err := utils.CreateFilePath(config.OutputDirectory); err != nil {
		return err
	}

	return f.Save(fmt.Sprintf("%s%s", config.OutputDirectory, config.OutputFileName))
}

// GenerateTypes parses the provided JSON payload and returns the matching Go types.
func GenerateTypes(data interface{}, config *Config) string {
	f := generateTypes(data, config)

	return fmt.Sprintf("%#v", f)
}

func generateTypes(data interface{}, config *Config) *jen.File {
	f := jen.NewFile(config.PackageName)

	addStructsFromJSON(f, data, config)

	return f
}
