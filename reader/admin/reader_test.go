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
	assert.Len(t, rawData.Items, 190)
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
	item := rawData.Items[1]

	assert.Equal(t, "AFS", item.Code)
	assert.Equal(t, "South Africa", item.EnglishName)
	assert.Equal(t, "Sudafricaine (Rép.)", item.FrenchName)
	assert.Equal(t, "Sudafricana (Rep.)", item.SpanishName)
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
