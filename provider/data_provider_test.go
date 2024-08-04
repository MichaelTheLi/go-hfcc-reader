package provider

import (
	"github.com/MichaelTheLi/go-hfcc-reader/reader"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestItemsCountCorrect(t *testing.T) {
	provider := NewDataProvider(
		reader.NewFileReader("../resources/test_hfcc_format_file.txt"),
	)
	items := provider.PullRawData()
	assert.NotEmpty(t, items)
}

func TestFileMetadataCorrect(t *testing.T) {
	assert.Fail(t, "Not implemented yet")
	//provider := NewDataProvider(
	//	reader.NewFileReader("../resources/test_hfcc_format_file.txt"),
	//)

	//metadata := provider.GetMetadata()
	//assert.NotEmpty(t, metadata)
}

func TestItemIsCorrect(t *testing.T) {
	provider := NewDataProvider(
		reader.NewFileReader("../resources/test_hfcc_format_file.txt"),
	)
	items := provider.PullRawData()
	item := items["1022"]

	assert.Equal(t, 2485, item.Frequency)
	assert.Equal(t, "1000", item.StartTime)
	assert.Equal(t, "1900", item.EndTime)
	assert.Equal(t, []CIRAFZoneId{"56", "51"}, item.CIRAFZones)
	assert.Equal(t, LocationId("PVL"), item.Location)
	assert.Equal(t, 10, item.Power)
	assert.Equal(t, 46, item.Azimuth)
	assert.Equal(t, AntennaId(400), item.Antenna)
	assert.Equal(t, 0, item.AntennaSlewAngle)
	assert.Equal(t, []time.Weekday{time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday}, item.DaysActive)

	startDate, err := time.Parse("02.01.06", "31.03.24")
	assert.Nil(t, err)
	assert.Equal(t, startDate, item.StartDate)

	endDate, _ := time.Parse("02.01.06", "27.10.24")
	assert.Nil(t, err)
	assert.Equal(t, endDate, item.EndDate)

	assert.Equal(t, DSB, item.Modulation)
	assert.Equal(t, 9000, item.AntennaDesignFrequency)
	assert.Equal(t, LanguageCode("Bis"), item.Language)
	assert.Equal(t, AdministrationId("VUT"), item.Administration)
	assert.Equal(t, BroadcasterId("VBT"), item.Broadcaster)
	assert.Equal(t, FMOrgId("RNZ"), item.FmOrgId)
	assert.Equal(t, "1022", item.Id)
	assert.Equal(t, []int{1234, 2345, 3456}, item.AlternativeFrequencies)
	assert.Equal(t, "NZL", item.Notes)
}
