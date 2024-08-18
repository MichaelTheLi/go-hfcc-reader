package antenna

import (
	"github.com/MichaelTheLi/go-hfcc-reader/reader"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAntennaItemsCountCorrect(t *testing.T) {
	fileReader := getReader()

	AntennaReader, err := fileReader.ProcessFile()
	assert.Nil(t, err)
	rawData := AntennaReader.(*Reader).RawAntennaData
	assert.Len(t, rawData.Items, 3)
}

func TestAntennaMetadataCorrect(t *testing.T) {
	fileReader := getReader()
	AntennaReader, err := fileReader.ProcessFile()
	assert.Nil(t, err)
	rawData := AntennaReader.(*Reader).RawAntennaData
	metadata := rawData.Metadata
	assert.NotEmpty(t, metadata)
	assert.Equal(t, "29-MAY-2019", metadata.Date)
	assert.Equal(t, "ANTENNA.TXT REFERENCE TABLE", metadata.Name)
}

func TestAntennaItemIsCorrect(t *testing.T) {
	fileReader := getReader()
	AntennaReader, err := fileReader.ProcessFile()
	assert.Nil(t, err)
	rawData := AntennaReader.(*Reader).RawAntennaData
	item := rawData.Items[1]
	assert.Equal(t, "101", item.Code)
	assert.Equal(t, "AHR1/1/0.5", item.Definition)
	assert.Equal(t, "", item.Notes)
}

func TestAntennaItemWithNotesIsCorrect(t *testing.T) {
	fileReader := getReader()
	AntennaReader, err := fileReader.ProcessFile()
	assert.Nil(t, err)
	rawData := AntennaReader.(*Reader).RawAntennaData
	item := rawData.Items[0]

	assert.Equal(t, "100", item.Code)
	assert.Equal(t, "1===============================================1", item.Definition)
	assert.Equal(t, "2==================2", item.Notes)
}

func getReader() reader.FileReader {
	path, _ := os.Getwd()
	AntennaFileProcessor := NewAntennaFileReader()
	return reader.NewFileReader(
		path+"/../../resources/antenna.txt",
		reader.NewLineReader(),
		&AntennaFileProcessor,
		nil,
	)
}
