package jenshared

import (
	"sort"

	"github.com/dave/jennifer/jen"
)

func addStructs(f *jen.File, itemMap TypeItemsMap, config *Config) {
	keys := make([]string, 0, len(itemMap))
	for key := range itemMap {
		keys = append(keys, key)
	}

	if config.Alphabetical {
		sort.Strings(keys)
	}

	for _, key := range keys {
		f.Add(createStruct(key, itemMap[key], config))
		f.Line()
	}
}

func addStruct(f *jen.File, name string, items TypeItems, config *Config) {
	f.Add(createStruct(name, items, config))
}

func createStruct(name string, items TypeItems, config *Config) *jen.Statement {

	if len(items) == 1 && name == items[0].Name {
		return jen.Type().Id(name).Id(items[0].Type)
	}

	structItems := createStructItems(items, config)

	return jen.Type().Id(name).Struct(structItems...)
}

func createStructItems(items TypeItems, config *Config) []jen.Code {
	structItems := make([]jen.Code, 0)

	if config.Alphabetical {
		sort.Slice(items, func(i, ii int) bool {
			return items[i].Title() < items[ii].Title()
		})
	}

	for _, item := range items {
		item.OmitEmpty = config.OmitEmpty
		structItems = append(structItems, createStructItem(item))
	}

	return structItems
}

func createStructItem(item TypeItem) jen.Code {
	s := jen.Id(item.Title())

	switch item.Type {
	case "time":
		s.Qual("time", "Time")
	case "uuid":
		s.Qual("github.com/google/uuid", "UUID")
	default:
		s.Id(item.Type)
	}

	s.Tag(item.Tags())

	return s
}
