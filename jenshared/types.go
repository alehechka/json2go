package jenshared

import (
	"log"
	"regexp"
	"strings"
)

// Config presents configurations to CreateTypes
type Config struct {
	RootName        string
	PackageName     string
	OutputFileName  string
	OutputDirectory string
	TimeFormat      string
	Debugger        *log.Logger
}

// TypeItem represents a parsed JSON variable
type TypeItem struct {
	Name string
	Type string
}

// Title converts the JSON name to TitleCase
func (t TypeItem) Title() string {
	specialCharacters := regexp.MustCompile(`[^a-zA-Z0-9]`)
	return strings.Title(specialCharacters.ReplaceAllString(t.Name, "_"))
}

// TypeItems is an array of TypeItem objects
type TypeItems []TypeItem

// TypeItemsMap is a map of TypeItems arrays
type TypeItemsMap map[string]TypeItems
