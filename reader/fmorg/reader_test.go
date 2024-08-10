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
	assert.Len(t, rawData.Items, 164)
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
	item := rawData.Items[12]

	assert.Equal(t, "BAB", item.Code)
	assert.Equal(t, "Babcock Communications", item.EnglishName)
	assert.Equal(t, "Mr. Gary Stanley", item.ContactPerson)
	assert.Equal(t, "+442073445781", item.Telephone)
	assert.Equal(t, "+442073966227", item.Fax)
	assert.Equal(t, "opssfm@babcock.co.uk", item.Email)
	assert.Equal(t, "HFCC/ABU-HFC", item.Notes)
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
