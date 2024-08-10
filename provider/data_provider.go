package provider

import (
	"github.com/MichaelTheLi/go-hfcc-reader/reader"
	"strconv"
	"strings"
	"time"
)

type DataProvider struct {
	fileReader reader.FileReader
	Data       Data
}

type Data struct {
	ProgramsList map[string]DataItem
	Metadata     Metadata
}

type Metadata struct {
	Season         string
	Administration string
	Date           time.Time
	Notes          []string
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
		fileReader: fileReader,
		Data: Data{
			ProgramsList: make(map[string]DataItem),
		},
	}
}

func (source DataProvider) GetData() Data {
	processor := source.fileReader.ProcessFile()
	rawData := processor.(*reader.ProgramFileReader).RawProgramsData

	for _, item := range rawData.Items {
		dataItem := source.getDataItem(*item)
		source.Data.ProgramsList[dataItem.Id] = dataItem
	}

	source.Data.Metadata = source.getMetadata(rawData)

	return source.Data
}

func (source DataProvider) getMetadata(rawData reader.RawProgramsData) Metadata {
	metadata := rawData.Metadata
	date, _ := time.Parse("02-Jan-2006", metadata.Date)

	return Metadata{
		Season:         metadata.Season,
		Administration: metadata.Administration,
		Date:           date,
		Notes:          []string{rawData.Note1.Value, rawData.Note2.Value, rawData.Note3.Value},
	}
}

func (source DataProvider) getDataItem(rawProgram reader.RawProgram) DataItem {
	// TODO Errors
	freq, _ := strconv.Atoi(rawProgram.Frequency)
	power, _ := strconv.Atoi(rawProgram.Power)
	azimuth, _ := strconv.Atoi(rawProgram.Azimuth)
	antennaSlewAngle, _ := strconv.Atoi(rawProgram.AntennaSlewAngle)
	antennaId, _ := strconv.Atoi(rawProgram.Antenna)

	var cirafZones []CIRAFZoneId
	var rawCirafZones = strings.Split(rawProgram.CIRAF, ",")
	for i := range rawCirafZones {
		cirafZones = append(cirafZones, CIRAFZoneId(rawCirafZones[i]))
	}
	var daysActive []time.Weekday
	for i := range rawProgram.DaysActive {
		dayNum, _ := strconv.Atoi(string(rawProgram.DaysActive[i]))
		daysActive = append(daysActive, time.Weekday(dayNum-1))
	}
	var alternativeFrequencies []int

	if rawProgram.Alt1 != "" {
		altFreq, _ := strconv.Atoi(rawProgram.Alt1)
		alternativeFrequencies = append(alternativeFrequencies, altFreq)
	}
	if rawProgram.Alt2 != "" {
		altFreq, _ := strconv.Atoi(rawProgram.Alt2)
		alternativeFrequencies = append(alternativeFrequencies, altFreq)
	}
	if rawProgram.Alt3 != "" {
		altFreq, _ := strconv.Atoi(rawProgram.Alt3)
		alternativeFrequencies = append(alternativeFrequencies, altFreq)
	}

	antennaDesignFrequency, _ := strconv.Atoi(rawProgram.AntennaDesignFrequency)

	startDate, _ := time.Parse("020106", rawProgram.StartDate)
	endDate, _ := time.Parse("020106", rawProgram.EndDate)
	dataItem := DataItem{
		Id:                     rawProgram.Id,
		Frequency:              freq,
		StartTime:              rawProgram.StartTime,
		EndTime:                rawProgram.EndTime,
		CIRAFZones:             cirafZones,
		Location:               LocationId(rawProgram.Location),
		Power:                  power,
		Azimuth:                azimuth,
		AntennaSlewAngle:       antennaSlewAngle,
		Antenna:                AntennaId(antennaId),
		DaysActive:             daysActive,
		StartDate:              startDate,
		EndDate:                endDate,
		Modulation:             Modulation(rawProgram.Modulation),
		AntennaDesignFrequency: antennaDesignFrequency,
		Language:               LanguageCode(rawProgram.Language),
		Administration:         AdministrationId(rawProgram.Administration),
		Broadcaster:            BroadcasterId(rawProgram.Broadcaster),
		FmOrgId:                FMOrgId(rawProgram.FmOrgId),
		AlternativeFrequencies: alternativeFrequencies,
		Notes:                  rawProgram.Notes,
	}
	return dataItem
}
