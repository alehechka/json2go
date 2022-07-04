package jenshared

import (
	"fmt"
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
	OmitEmpty       bool
	Debugger        *log.Logger
}

// TypeItem represents a parsed JSON variable
type TypeItem struct {
	Name      string
	Type      string
	OmitEmpty bool
}

// Title converts the JSON name to TitleCase
func (t TypeItem) Title() string {
	str := regexp.MustCompile(`[^a-zA-Z0-9]`).ReplaceAllString(t.Name, "_")

	numbers := regexp.MustCompile(`\d`)
	if len(str) > 0 && numbers.MatchString(str[0:1]) {
		str = fmt.Sprintf("_%s", str[1:])
	}

	return strings.Title(str)
}

// TagJSON prepares the tag for provided item
func (t TypeItem) TagJSON() string {
	tags := []string{t.Name}
	if t.OmitEmpty {
		tags = append(tags, "omitempty")
	}

	return strings.Join(tags, ",")
}

// Tags creates tha Tags map
func (t TypeItem) Tags() map[string]string {
	return map[string]string{"json": t.TagJSON()}
}

// TypeItems is an array of TypeItem objects
type TypeItems []TypeItem

// TypeItemsMap is a map of TypeItems arrays
type TypeItemsMap map[string]TypeItems
