package program

import (
	"github.com/MichaelTheLi/go-hfcc-reader/reader"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestItemsCountCorrect(t *testing.T) {
	fileReader := getReader()

	programReader := fileReader.ProcessFile()
	rawData := programReader.(*Reader).RawProgramsData
	assert.Len(t, rawData.Items, 5)
}

func TestMetadataCorrect(t *testing.T) {
	fileReader := getReader()
	programReader := fileReader.ProcessFile()
	rawData := programReader.(*Reader).RawProgramsData
	metadata := rawData.Metadata

	assert.NotEmpty(t, metadata)
	assert.Equal(t, "A24", metadata.Season)
	assert.Equal(t, "ALL", metadata.Administration)
	assert.Equal(t, "23-jul-2024", metadata.Date)
	assert.Equal(t, "Global HF Schedule", rawData.Note1.Value)
	assert.Equal(t, "Processed on 23-jul-2024 at 09:37UTC", rawData.Note2.Value)
	assert.Equal(t, "Timestamp: 1721727440", rawData.Note3.Value)
}

func TestItemIsCorrect(t *testing.T) {
	fileReader := getReader()
	programReader := fileReader.ProcessFile()
	rawData := programReader.(*Reader).RawProgramsData
	item := rawData.Items[0]

	assert.Equal(t, "2485", item.Frequency)
	assert.Equal(t, "1000", item.StartTime)
	assert.Equal(t, "1900", item.EndTime)
	assert.Equal(t, "56,51", item.CIRAF)
	assert.Equal(t, "PVL", item.Location)
	assert.Equal(t, "10", item.Power)
	assert.Equal(t, "46", item.Azimuth)
	assert.Equal(t, "400", item.Antenna)
	assert.Equal(t, "0", item.AntennaSlewAngle)
	assert.Equal(t, "1234567", item.DaysActive)
	assert.Equal(t, "310324", item.StartDate)
	assert.Equal(t, "271024", item.EndDate)
	assert.Equal(t, "D", item.Modulation)
	assert.Equal(t, "9000", item.AntennaDesignFrequency)
	assert.Equal(t, "Bis", item.Language)
	assert.Equal(t, "VUT", item.Administration)
	assert.Equal(t, "VBT", item.Broadcaster)
	assert.Equal(t, "RNZ", item.FmOrgId)
	assert.Equal(t, "1022", item.Id)
	assert.Equal(t, "1", item.OldData)
	assert.Equal(t, "1234", item.Alt1)
	assert.Equal(t, "2345", item.Alt2)
	assert.Equal(t, "3456", item.Alt3)
	assert.Equal(t, "NZL", item.Notes)
}

func getReader() reader.FileReader {
	path, _ := os.Getwd()
	programFileProcessor := NewProgramFileReader()
	return reader.NewFileReader(
		path+"/../../resources/test_hfcc_format_file.txt",
		reader.NewLineReader(),
		&programFileProcessor,
		nil,
	)
}
