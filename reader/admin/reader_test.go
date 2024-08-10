package admin

import (
	"github.com/MichaelTheLi/go-hfcc-reader/reader"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/charmap"
	"os"
	"testing"
)

func TestAdminItemsCountCorrect(t *testing.T) {
	fileReader := getReader()

	AdminReader := fileReader.ProcessFile()
	rawData := AdminReader.(*Reader).RawAdminData
	assert.Len(t, rawData.Items, 3)
}

func TestAdminMetadataCorrect(t *testing.T) {
	fileReader := getReader()
	AdminReader := fileReader.ProcessFile()
	rawData := AdminReader.(*Reader).RawAdminData
	metadata := rawData.Metadata

	assert.NotEmpty(t, metadata)
	assert.Equal(t, "19-MAR-2012", metadata.Date)
	assert.Equal(t, "ADMIN.TXT REFERENCE TABLE", metadata.Name)
}

func TestAdminItemIsCorrect(t *testing.T) {
	fileReader := getReader()
	AdminReader := fileReader.ProcessFile()
	rawData := AdminReader.(*Reader).RawAdminData
	item := rawData.Items[0]

	assert.Equal(t, "AFG", item.Code)
	assert.Equal(t, "1================================================1", item.EnglishName)
	assert.Equal(t, "2================================================2", item.FrenchName)
	assert.Equal(t, "3===============================================3", item.SpanishName)
}

func getReader() reader.FileReader {
	path, _ := os.Getwd()
	AdminFileProcessor := NewAdminFileReader()
	return reader.NewFileReader(
		path+"/../../resources/admin.txt",
		reader.NewLineReader(),
		&AdminFileProcessor,
		charmap.ISO8859_1,
	)
}
