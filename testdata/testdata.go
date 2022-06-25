package main

import (
	"github.com/alehechka/json2go/gen"
)

var configs = []*gen.Config{
	{File: "testdata/bool.json", PackageName: "testdata/out", RootName: "RootBoolean", OutputFileName: "bool.go"},
	{File: "testdata/float64.json", PackageName: "testdata/out", RootName: "RootFloat64", OutputFileName: "float64.go"},
	{File: "testdata/object.json", PackageName: "testdata/out", RootName: "RootObject", OutputFileName: "object.go"},
	{File: "testdata/nestedObjectArray.json", PackageName: "testdata/out", RootName: "RootNestedObject", OutputFileName: "nestedObjectArray.go"},
	{File: "testdata/objectArray.json", PackageName: "testdata/out", RootName: "RootObjectArray", OutputFileName: "objectArray.go"},
	{File: "testdata/string.json", PackageName: "testdata/out", RootName: "RootString", OutputFileName: "string.go"},
	{File: "testdata/stringArray.json", PackageName: "testdata/out", RootName: "RootStringArray", OutputFileName: "stringArray.go"},
	{File: "testdata/emptyArray.json", PackageName: "testdata/out", RootName: "RootEmptyArray", OutputFileName: "emptyArray.go"},
	{File: "testdata/emptyObject.json", PackageName: "testdata/out", RootName: "RootEmptyObject", OutputFileName: "emptyObject.go"},
	{File: "testdata/nestedBoolArray.json", PackageName: "testdata/out", RootName: "RootNestedBoolArray", OutputFileName: "nestedBoolArray.go"},
}

func main() {
	generator := gen.New()

	for _, config := range configs {
		generator.Build(config)
	}

}
