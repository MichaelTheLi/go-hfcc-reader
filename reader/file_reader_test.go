package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItemsCountCorrect(t *testing.T) {
	reader := NewFileReader("../resources/test_hfcc_format_file.txt")

	items := reader.GetRawItems()
	assert.Len(t, items, 5)
}

func TestMetadataCorrect(t *testing.T) {
	assert.Fail(t, "Not implemented yet")
	//reader := go-hfcc-reader.NewReader("./test_hfcc_format_file.txt")

	//metadata := reader.GetMetadata()
	//assert.NotEmpty(t, metadata)
}

func TestItemIsCorrect(t *testing.T) {
	reader := NewFileReader("../resources/test_hfcc_format_file.txt")
	items := reader.GetRawItems()
	item := items[0]

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
