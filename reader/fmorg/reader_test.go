package fmOrg

import (
	"github.com/MichaelTheLi/go-hfcc-reader/reader"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFmOrgItemsCountCorrect(t *testing.T) {
	fileReader := getReader()

	FmOrgReader := fileReader.ProcessFile()
	rawData := FmOrgReader.(*Reader).RawFmOrgData
	assert.Len(t, rawData.Items, 3)
}

func TestFmOrgMetadataCorrect(t *testing.T) {
	fileReader := getReader()
	FmOrgReader := fileReader.ProcessFile()
	rawData := FmOrgReader.(*Reader).RawFmOrgData
	metadata := rawData.Metadata
	assert.NotEmpty(t, metadata)
	assert.Equal(t, "25-JAN-2023", metadata.Date)
	assert.Equal(t, "Reference Table Freq. Management Org.", metadata.Name)
}

func TestFmOrgItemIsCorrect(t *testing.T) {
	fileReader := getReader()
	FmOrgReader := fileReader.ProcessFile()
	rawData := FmOrgReader.(*Reader).RawFmOrgData
	item := rawData.Items[0]

	assert.Equal(t, "ITU", item.Code)
	assert.Equal(t, "1================================================1", item.EnglishName)
	assert.Equal(t, "2==================2", item.ContactPerson)
	assert.Equal(t, "3===========3", item.Telephone)
	assert.Equal(t, "4===========4", item.Fax)
	assert.Equal(t, "5======================================5", item.Email)
	assert.Equal(t, "6===========6", item.Notes)
}

func getReader() reader.FileReader {
	path, _ := os.Getwd()
	FmOrgFileProcessor := NewFmOrgFileReader()
	return reader.NewFileReader(
		path+"/../../resources/fmorg.txt",
		reader.NewLineReader(),
		&FmOrgFileProcessor,
		nil,
	)
}
