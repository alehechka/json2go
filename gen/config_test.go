package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toJensharedConfig_Base(t *testing.T) {
	config := &Config{
		File:           "data.json",
		OutputFileName: "output.go",
		RootName:       "Response",
		PackageName:    "gql",
	}

	jenConfig := config.toJensharedConfig()

	assert.Equal(t, config.RootName, jenConfig.RootName)
	assert.Equal(t, config.PackageName, jenConfig.PackageName)
	assert.Equal(t, config.OutputFileName, jenConfig.OutputFileName)
}

func Test_prepareOutputFileName_File(t *testing.T) {
	config := &Config{
		File:           "data.json",
		OutputFileName: DefaultOutputFile,
	}

	jenConfig := config.toJensharedConfig()

	assert.Equal(t, "data.go", jenConfig.OutputFileName)
}

func Test_prepareOutputFileName_FileNoExtension(t *testing.T) {
	config := &Config{
		File:           "data",
		OutputFileName: DefaultOutputFile,
	}

	jenConfig := config.toJensharedConfig()

	assert.Equal(t, "data.go", jenConfig.OutputFileName)
}

func Test_prepareOutputFileName_NestedFile(t *testing.T) {
	config := &Config{
		File:           "/path/to/file/data.json",
		OutputFileName: DefaultOutputFile,
	}

	jenConfig := config.toJensharedConfig()

	assert.Equal(t, "data.go", jenConfig.OutputFileName)
}

func Test_prepareOutputFileName_FileEndsWithSlashes(t *testing.T) {
	config := &Config{
		File:           "/path/to/file/data//",
		OutputFileName: DefaultOutputFile,
	}

	jenConfig := config.toJensharedConfig()

	assert.Equal(t, "data.go", jenConfig.OutputFileName)
}

func Test_prepareOutputFileName_RelativeFile(t *testing.T) {
	config := &Config{
		File:           "../data.json",
		OutputFileName: DefaultOutputFile,
	}

	jenConfig := config.toJensharedConfig()

	assert.Equal(t, "data.go", jenConfig.OutputFileName)
}

func Test_prepareOutputFileName_BadInput_ForwardSlash(t *testing.T) {
	config := &Config{
		File:           "/",
		OutputFileName: DefaultOutputFile,
	}

	jenConfig := config.toJensharedConfig()

	assert.Equal(t, DefaultOutputFile, jenConfig.OutputFileName)
}

func Test_prepareOutputFileName_BadInput_ForwardSlashes(t *testing.T) {
	config := &Config{
		File:           "//",
		OutputFileName: DefaultOutputFile,
	}

	jenConfig := config.toJensharedConfig()

	assert.Equal(t, DefaultOutputFile, jenConfig.OutputFileName)
}

func Test_prepareOutputFileName_BadInput_Empty(t *testing.T) {
	config := &Config{
		File:           "",
		OutputFileName: DefaultOutputFile,
	}

	jenConfig := config.toJensharedConfig()

	assert.Equal(t, DefaultOutputFile, jenConfig.OutputFileName)
}

func Test_prepareOutputFileName_BadInput_Dot(t *testing.T) {
	config := &Config{
		File:           ".",
		OutputFileName: DefaultOutputFile,
	}

	jenConfig := config.toJensharedConfig()

	assert.Equal(t, DefaultOutputFile, jenConfig.OutputFileName)
}

func Test_prepareOutputFileName_BadInput_Dots(t *testing.T) {
	config := &Config{
		File:           "..",
		OutputFileName: DefaultOutputFile,
	}

	jenConfig := config.toJensharedConfig()

	assert.Equal(t, DefaultOutputFile, jenConfig.OutputFileName)
}

func Test_prepareOutputFileName_Output(t *testing.T) {
	config := &Config{
		File:           "data.json",
		OutputFileName: "output",
	}

	jenConfig := config.toJensharedConfig()

	assert.Equal(t, "output.go", jenConfig.OutputFileName)
}

func Test_prepareOutputFileName_OutputWithExt(t *testing.T) {
	config := &Config{
		File:           "data.json",
		OutputFileName: "output.go",
	}

	jenConfig := config.toJensharedConfig()

	assert.Equal(t, "output.go", jenConfig.OutputFileName)
}
