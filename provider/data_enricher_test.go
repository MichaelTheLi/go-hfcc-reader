package provider

import (
	"github.com/MichaelTheLi/go-hfcc-reader/reader"
	"github.com/MichaelTheLi/go-hfcc-reader/reader/admin"
	"github.com/MichaelTheLi/go-hfcc-reader/reader/antenna"
	"github.com/MichaelTheLi/go-hfcc-reader/reader/broadcaster"
	fmOrg "github.com/MichaelTheLi/go-hfcc-reader/reader/fmorg"
	"github.com/MichaelTheLi/go-hfcc-reader/reader/language"
	"github.com/MichaelTheLi/go-hfcc-reader/reader/program"
	"github.com/MichaelTheLi/go-hfcc-reader/reader/site"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestEnricherItemsCountCorrect(t *testing.T) {
	enricher := getEnricher()
	items := enricher.GetEnrichedData()
	assert.Len(t, items, 5)
}

func TestEnricherItemIsCorrect(t *testing.T) {
	enricher := getEnricher()
	items := enricher.GetEnrichedData()
	item := items["1022"]

	assert.Equal(t, "1022", item.Item.Id)

	assert.Equal(t, string(item.Item.Administration), item.Admin.Code)
	assert.Equal(t, "Vanuatu", item.Admin.EnglishName)

	assert.Equal(t, strconv.Itoa(int(item.Item.Antenna)), item.Antenna.Code)
	assert.Equal(t, "CHR(S)4/1/0.3", item.Antenna.Definition)

	assert.Equal(t, string(item.Item.Broadcaster), item.Broadcaster.Code)
	assert.Equal(t, "Vanuatu Broadcasting and Television Corporation", item.Broadcaster.EnglishName)

	assert.Equal(t, string(item.Item.FmOrgId), item.FmOrg.Code)
	assert.Equal(t, "Radio New Zealand Ltd.", item.FmOrg.EnglishName)

	assert.Equal(t, string(item.Item.Language), item.Language.Code)
	assert.Equal(t, "Bislama", item.Language.EnglishName)

	assert.Equal(t, string(item.Item.Location), item.Site.Code)
	assert.Equal(t, "Port Vila", item.Site.EnglishName)
}

func getEnricher() DataEnricher {
	provider := NewDataProvider(
		getProgramReader(),
	)

	return NewDataEnricher(
		provider.GetData(),
		getAdminReader(),
		getAntennaReader(),
		getBroadcasterReader(),
		getFmOrgReader(),
		getLanguageReader(),
		getSiteReader(),
	)
}

func getAdminReader() reader.FileReader {
	processor := admin.NewAdminFileReader()

	return reader.NewFileReader(
		"resources/admin.txt",
		reader.NewLineReader(),
		&processor,
		nil,
	)
}

func getAntennaReader() reader.FileReader {
	processor := antenna.NewAntennaFileReader()

	return reader.NewFileReader(
		"resources/antenna.txt",
		reader.NewLineReader(),
		&processor,
		nil,
	)
}

func getBroadcasterReader() reader.FileReader {
	processor := broadcaster.NewBroadcasterFileReader()

	return reader.NewFileReader(
		"resources/broadcas.txt",
		reader.NewLineReader(),
		&processor,
		nil,
	)
}

func getFmOrgReader() reader.FileReader {
	processor := fmOrg.NewFmOrgFileReader()

	return reader.NewFileReader(
		"resources/fmorg.txt",
		reader.NewLineReader(),
		&processor,
		nil,
	)
}

func getLanguageReader() reader.FileReader {
	languageFileProcessor := language.NewLanguageFileReader()

	return reader.NewFileReader(
		"resources/language.txt",
		reader.NewLineReader(),
		&languageFileProcessor,
		nil,
	)
}

func getProgramReader() reader.FileReader {
	processor := program.NewProgramFileReader()

	return reader.NewFileReader(
		"resources/test_hfcc_format_file.txt",
		reader.NewLineReader(),
		&processor,
		nil,
	)
}

func getSiteReader() reader.FileReader {
	processor := site.NewSiteFileReader()

	return reader.NewFileReader(
		"resources/site.txt",
		reader.NewLineReader(),
		&processor,
		nil,
	)
}
