package utils

import (
	"os"
	"strings"
)

// CreateFilePath will create the desired filePath if it does not exist
func CreateFilePath(filePath string) error {
	pathParts := strings.Split(filePath, "/")
	pathBuilder := ""
	for _, part := range pathParts {
		pathBuilder = pathBuilder + part + "/"
		if _, err := os.Stat(pathBuilder); os.IsNotExist(err) {
			if err := os.Mkdir(pathBuilder, 0755); err != nil {
				return err
			}
		}
	}
	return nil
}
