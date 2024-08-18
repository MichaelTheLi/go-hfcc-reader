package site

import (
	"github.com/MichaelTheLi/go-hfcc-reader/reader"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSiteItemsCountCorrect(t *testing.T) {
	fileReader := getReader()

	SiteReader, err := fileReader.ProcessFile()
	assert.Nil(t, err)
	rawData := SiteReader.(*Reader).RawSiteData
	assert.Len(t, rawData.Items, 3)
}

func TestSiteMetadataCorrect(t *testing.T) {
	fileReader := getReader()
	SiteReader, err := fileReader.ProcessFile()
	assert.Nil(t, err)
	rawData := SiteReader.(*Reader).RawSiteData
	metadata := rawData.Metadata
	assert.NotEmpty(t, metadata)
	assert.Equal(t, "11-May-2023", metadata.Date)
	assert.Equal(t, "Global HF Transmitter Site Table", metadata.Name)
}

func TestSiteItemIsCorrect(t *testing.T) {
	fileReader := getReader()
	SiteReader, err := fileReader.ProcessFile()
	assert.Nil(t, err)
	rawData := SiteReader.(*Reader).RawSiteData
	item := rawData.Items[0]

	assert.Equal(t, "A-A", item.Code)
	assert.Equal(t, "1============================1", item.EnglishName)
	assert.Equal(t, "KAZ", item.Administration)
	assert.Equal(t, "43N17", item.Latitude)
	assert.Equal(t, "077E00", item.Longitude)
}

func getReader() reader.FileReader {
	path, _ := os.Getwd()
	SiteFileProcessor := NewSiteFileReader()
	return reader.NewFileReader(
		path+"/../../resources/site.txt",
		reader.NewLineReader(),
		&SiteFileProcessor,
		nil,
	)
}
