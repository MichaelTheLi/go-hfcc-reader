package provider

import (
	"MichaelTheLi/hfcc/reader"
	"strconv"
	"strings"
	"time"
)

type DataProvider struct {
	fileReader   reader.FileReader
	programsList map[string]DataItem
}

type CIRAFZoneId string
type LocationId string
type AdministrationId string
type BroadcasterId string
type FMOrgId string
type LanguageCode string
type AntennaId int
type Modulation string

const (
	DSB Modulation = "D"
	SSB Modulation = "T"
)

type DataItem struct {
	Id                     string
	Frequency              int
	StartTime              string
	EndTime                string
	CIRAFZones             []CIRAFZoneId
	Location               LocationId
	Power                  int
	Azimuth                int
	AntennaSlewAngle       int
	Antenna                AntennaId
	DaysActive             []time.Weekday
	StartDate              time.Time
	EndDate                time.Time
	Modulation             Modulation
	AntennaDesignFrequency int
	Language               LanguageCode
	Administration         AdministrationId
	Broadcaster            BroadcasterId
	FmOrgId                FMOrgId
	AlternativeFrequencies []int
	Notes                  string
}

func (dataItem DataItem) FreqString() string {
	return strconv.Itoa(dataItem.Frequency)
}
func (dataItem DataItem) Name() string {
	return string(dataItem.Broadcaster) + " of " + string(dataItem.Administration) + " at " + dataItem.FreqString()
}

func NewDataProvider(fileReader reader.FileReader) DataProvider {
	return DataProvider{
		fileReader:   fileReader,
		programsList: make(map[string]DataItem),
	}
}

func (source DataProvider) PullRawData() map[string]DataItem {
	rawItems := source.fileReader.GetRawItems()

	for _, item := range rawItems {
		dataItem := source.getDataItem(item)
		source.programsList[dataItem.Id] = dataItem
	}
	return source.programsList
}

func (source DataProvider) getDataItem(item reader.RawDataItem) DataItem {
	// TODO Errors
	freq, _ := strconv.Atoi(item.Frequency)
	power, _ := strconv.Atoi(item.Power)
	azimuth, _ := strconv.Atoi(item.Azimuth)
	antennaSlewAngle, _ := strconv.Atoi(item.AntennaSlewAngle)
	antennaId, _ := strconv.Atoi(item.Antenna)

	var cirafZones []CIRAFZoneId
	var rawCirafZones = strings.Split(item.CIRAF, ",")
	for i := range rawCirafZones {
		cirafZones = append(cirafZones, CIRAFZoneId(rawCirafZones[i]))
	}
	var daysActive []time.Weekday
	for i := range item.DaysActive {
		dayNum, _ := strconv.Atoi(string(item.DaysActive[i]))
		daysActive = append(daysActive, time.Weekday(dayNum-1))
	}
	var alternativeFrequencies []int

	if item.Alt1 != "" {
		altFreq, _ := strconv.Atoi(item.Alt1)
		alternativeFrequencies = append(alternativeFrequencies, altFreq)
	}
	if item.Alt2 != "" {
		altFreq, _ := strconv.Atoi(item.Alt2)
		alternativeFrequencies = append(alternativeFrequencies, altFreq)
	}
	if item.Alt3 != "" {
		altFreq, _ := strconv.Atoi(item.Alt3)
		alternativeFrequencies = append(alternativeFrequencies, altFreq)
	}

	antennaDesignFrequency, _ := strconv.Atoi(item.AntennaDesignFrequency)

	startDate, _ := time.Parse("020106", item.StartDate)
	endDate, _ := time.Parse("020106", item.EndDate)
	dataItem := DataItem{
		Id:                     item.Id,
		Frequency:              freq,
		StartTime:              item.StartTime,
		EndTime:                item.EndTime,
		CIRAFZones:             cirafZones,
		Location:               LocationId(item.Location),
		Power:                  power,
		Azimuth:                azimuth,
		AntennaSlewAngle:       antennaSlewAngle,
		Antenna:                AntennaId(antennaId),
		DaysActive:             daysActive,
		StartDate:              startDate,
		EndDate:                endDate,
		Modulation:             Modulation(item.Modulation),
		AntennaDesignFrequency: antennaDesignFrequency,
		Language:               LanguageCode(item.Language),
		Administration:         AdministrationId(item.Administration),
		Broadcaster:            BroadcasterId(item.Broadcaster),
		FmOrgId:                FMOrgId(item.FmOrgId),
		AlternativeFrequencies: alternativeFrequencies,
		Notes:                  item.Notes,
	}
	return dataItem
}
