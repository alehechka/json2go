package jenshared

import (
	"github.com/dave/jennifer/jen"
)

func addStructs(f *jen.File, itemMap TypeItemsMap) {
	for name, items := range itemMap {
		f.Add(createStruct(name, items))
		f.Line()
	}
}

func addStruct(f *jen.File, name string, items TypeItems) {
	f.Add(createStruct(name, items))
}

func createStruct(name string, items TypeItems) *jen.Statement {

	if len(items) == 1 && name == items[0].Name {
		return jen.Type().Id(name).Id(items[0].Type)
	}

	structItems := createStructItems(items)

	return jen.Type().Id(name).Struct(structItems...)
}

func createStructItems(items TypeItems) []jen.Code {
	structItems := make([]jen.Code, 0)

	for _, item := range items {
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

	if item.Name != "" {
		s.Tag(map[string]string{"json": item.Name})
	}
	return s
}
