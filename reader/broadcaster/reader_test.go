package broadcaster

import (
	"github.com/MichaelTheLi/go-hfcc-reader/reader"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBroadcasterItemsCountCorrect(t *testing.T) {
	fileReader := getReader()

	BroadcasterReader := fileReader.ProcessFile()
	rawData := BroadcasterReader.(*Reader).RawBroadcasterData
	assert.Len(t, rawData.Items, 303)
}

func TestBroadcasterMetadataCorrect(t *testing.T) {
	fileReader := getReader()
	BroadcasterReader := fileReader.ProcessFile()
	rawData := BroadcasterReader.(*Reader).RawBroadcasterData
	metadata := rawData.Metadata
	assert.NotEmpty(t, metadata)
	assert.Equal(t, "25-Jan-2023", metadata.Date)
	assert.Equal(t, "BROADCAST.TXT REFERENCE TABLE", metadata.Name)
}

func TestBroadcasterItemIsCorrect(t *testing.T) {
	fileReader := getReader()
	BroadcasterReader := fileReader.ProcessFile()
	rawData := BroadcasterReader.(*Reader).RawBroadcasterData
	item := rawData.Items[3]
	assert.Equal(t, "AGM", item.Code)
	assert.Equal(t, "United States Agency for Global Media (USAGM)", item.EnglishName)
}

func getReader() reader.FileReader {
	path, _ := os.Getwd()
	BroadcasterFileProcessor := NewBroadcasterFileReader()
	return reader.NewFileReader(
		path+"/../../resources/broadcas.txt",
		reader.NewLineReader(),
		&BroadcasterFileProcessor,
		nil,
	)
}
