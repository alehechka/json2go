package jenshared

import (
	"log"
	"strings"
)

// Config presents configurations to CreateTypes
type Config struct {
	RootName        string
	PackageName     string
	OutputFileName  string
	OutputDirectory string
	Debugger        *log.Logger
}

// TypeItem represents a parsed JSON variable
type TypeItem struct {
	Name string
	Type string
}

// Title converts the JSON name to TitleCase
func (t TypeItem) Title() string {
	return strings.Title(t.Name)
}

// TypeItems is an array of TypeItem objects
type TypeItems []TypeItem

// TypeItemsMap is a map of TypeItems arrays
type TypeItemsMap map[string]TypeItems
