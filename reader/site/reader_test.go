package site

import (
	"github.com/MichaelTheLi/go-hfcc-reader/reader"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSiteItemsCountCorrect(t *testing.T) {
	fileReader := getReader()

	SiteReader := fileReader.ProcessFile()
	rawData := SiteReader.(*Reader).RawSiteData
	assert.Len(t, rawData.Items, 700)
}

func TestSiteMetadataCorrect(t *testing.T) {
	fileReader := getReader()
	SiteReader := fileReader.ProcessFile()
	rawData := SiteReader.(*Reader).RawSiteData
	metadata := rawData.Metadata
	assert.NotEmpty(t, metadata)
	assert.Equal(t, "11-May-2023", metadata.Date)
	assert.Equal(t, "Global HF Transmitter Site Table", metadata.Name)
}

func TestSiteItemIsCorrect(t *testing.T) {
	fileReader := getReader()
	SiteReader := fileReader.ProcessFile()
	rawData := SiteReader.(*Reader).RawSiteData
	item := rawData.Items[2]

	assert.Equal(t, "ABG", item.Code)
	assert.Equal(t, "Abu Ghraib (Bagdadh)", item.EnglishName)
	assert.Equal(t, "IRQ", item.Administration)
	assert.Equal(t, "33N19", item.Latitude)
	assert.Equal(t, "044E15", item.Longitude)
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
