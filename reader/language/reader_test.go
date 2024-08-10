package language

import (
	"github.com/MichaelTheLi/go-hfcc-reader/reader"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLanguageItemsCountCorrect(t *testing.T) {
	fileReader := getReader()

	LanguageReader := fileReader.ProcessFile()
	rawData := LanguageReader.(*Reader).RawLanguageData
	assert.Len(t, rawData.Items, 3)
}

func TestLanguageMetadataCorrect(t *testing.T) {
	fileReader := getReader()
	LanguageReader := fileReader.ProcessFile()
	rawData := LanguageReader.(*Reader).RawLanguageData
	metadata := rawData.Metadata
	assert.NotEmpty(t, metadata)
	assert.Equal(t, "23-JAN-2014", metadata.Date)
	assert.Equal(t, "Reference Table Language", metadata.Name)
}

func TestLanguageItemIsCorrect(t *testing.T) {
	fileReader := getReader()
	LanguageReader := fileReader.ProcessFile()
	rawData := LanguageReader.(*Reader).RawLanguageData
	item := rawData.Items[0]

	assert.Equal(t, "Aaa", item.Code)
	assert.Equal(t, "1================================================1", item.EnglishName)
}

func getReader() reader.FileReader {
	path, _ := os.Getwd()
	LanguageFileProcessor := NewLanguageFileReader()
	return reader.NewFileReader(
		path+"/../../resources/language.txt",
		reader.NewLineReader(),
		&LanguageFileProcessor,
		nil,
	)
}
