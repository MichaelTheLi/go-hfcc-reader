package reader

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAdminItemsCountCorrect(t *testing.T) {
	reader := getAdminReader()

	AdminReader := reader.ProcessFile()
	rawData := AdminReader.(*AdminFileReader).RawAdminData
	assert.Len(t, rawData.Items, 190)
}

func TestAdminMetadataCorrect(t *testing.T) {
	reader := getAdminReader()
	AdminReader := reader.ProcessFile()
	rawData := AdminReader.(*AdminFileReader).RawAdminData
	metadata := rawData.Metadata

	assert.NotEmpty(t, metadata)
	assert.Equal(t, "19-MAR-2012", metadata.Date)
	assert.Equal(t, "ADMIN.TXT", metadata.Name)
	assert.Equal(t, "REFERENCE TABLE", metadata.Note)
}

func TestAdminItemIsCorrect(t *testing.T) {
	reader := getAdminReader()
	AdminReader := reader.ProcessFile()
	rawData := AdminReader.(*AdminFileReader).RawAdminData
	item := rawData.Items[1]

	assert.Equal(t, "AFS", item.Code)
	assert.Equal(t, "South Africa", item.EnglishName)
	assert.Equal(t, "Sudafricaine (Rép.)", item.FrenchName)
	assert.Equal(t, "Sudafricana (Rep.)", item.SpanishName)
}

func getAdminReader() FileReader {
	path, _ := os.Getwd()
	AdminFileProcessor := NewAdminFileReader()
	reader := NewFileReader(
		path+"/../resources/admin.txt",
		LineReader{},
		&AdminFileProcessor,
	)
	return reader
}
