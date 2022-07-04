package jenshared

import (
	"fmt"
	"strings"
	"time"

	"github.com/dave/jennifer/jen"
	"github.com/google/uuid"
)

func addStructsFromJSON(f *jen.File, data interface{}, config *Config) {
	typeItemsMap := createTypeItemsMapFromJSON(data, config)
	addStructs(f, typeItemsMap)
}

func createTypeItemsMapFromJSON(data interface{}, config *Config) TypeItemsMap {
	typeItemsMap := make(TypeItemsMap)

	return parseInterface(typeItemsMap, data, config)
}

func parseInterface(items TypeItemsMap, data interface{}, config *Config) TypeItemsMap {
	switch concreteVal := data.(type) {
	case bool, float64, string:
		items[config.RootName] = TypeItems{{Name: config.RootName, Type: inferDataType(concreteVal, config), OmitEmpty: config.OmitEmpty}}
	case map[string]interface{}:
		parseMap(items, concreteVal, config.RootName, config)
	case []interface{}:
		diveTopLevelArray(items, concreteVal, config, "[]")
	}

	return items
}

func diveTopLevelArray(items TypeItemsMap, data []interface{}, config *Config, acc string) TypeItemsMap {
	if len(data) > 0 {
		switch firstVal := data[0].(type) {
		case bool, float64, string:
			items[config.RootName] = TypeItems{{Name: config.RootName, Type: fmt.Sprintf("%s%s", acc, inferDataType(firstVal, config)), OmitEmpty: config.OmitEmpty}}
		case map[string]interface{}:
			arrTitle := fmt.Sprintf("%sArray", config.RootName)
			items[arrTitle] = TypeItems{{Name: arrTitle, Type: fmt.Sprintf("%s%s", acc, config.RootName), OmitEmpty: config.OmitEmpty}}
			parseMap(items, firstVal, config.RootName, config)

		case []interface{}:
			diveTopLevelArray(items, firstVal, config, fmt.Sprintf("%s[]", acc))
		}
	}

	return items
}

func parseMap(items TypeItemsMap, data map[string]interface{}, parent string, config *Config) TypeItemsMap {
	for key, val := range data {
		title := strings.Title(key)
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			items[title] = make(TypeItems, 0)
			items[parent] = append(items[parent], TypeItem{Name: key, Type: title})
			parseMap(items, concreteVal, title, config)
		case []interface{}:
			items[parent] = append(items[parent], TypeItem{Name: key, Type: fmt.Sprintf("[]%s", title), OmitEmpty: config.OmitEmpty})
			parseFirstIndexArray(items, concreteVal, title, config)
		default:
			items[parent] = append(items[parent], TypeItem{Name: key, Type: inferDataType(concreteVal, config), OmitEmpty: config.OmitEmpty})
		}
	}
	return items
}

func parseFirstIndexArray(items TypeItemsMap, array []interface{}, parent string, config *Config) TypeItemsMap {
	if len(array) > 0 {
		switch concreteVal := array[0].(type) {
		case map[string]interface{}:
			parseMap(items, concreteVal, parent, config)
		case []interface{}:
		InterfaceArrayOuter:
			for key, itemArray := range items {
				for index, item := range itemArray {
					if item.Title() == parent {
						items[key][index].Type = fmt.Sprintf("[]%s", item.Type)
						parseFirstIndexArray(items, concreteVal, parent, config)
						break InterfaceArrayOuter
					}
				}
			}
		default:
		DefaultOuter:
			for key, itemArray := range items {
				for index, item := range itemArray {
					if item.Title() == parent && strings.HasSuffix(item.Type, parent) {
						items[key][index].Type = fmt.Sprintf("%s%s", strings.TrimSuffix(item.Type, parent), inferDataType(concreteVal, config))
						break DefaultOuter
					}
				}
			}
		}
	}
	return items
}

func inferDataType(value interface{}, config *Config) string {
	valType := fmt.Sprintf("%T", value)
	if valType == "string" {
		if _, err := time.Parse(config.TimeFormat, value.(string)); err == nil {
			return "time"
		}
		if _, err := uuid.Parse(value.(string)); err == nil {
			return "uuid"
		}
	}
	return valType
}
