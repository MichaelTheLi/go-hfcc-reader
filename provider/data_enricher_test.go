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
	enricher, _ := getEnricher()
	items, _ := enricher.GetEnrichedData()
	assert.Len(t, items, 5)
}

func TestEnricherItemIsCorrect(t *testing.T) {
	enricher, _ := getEnricher()
	items, _ := enricher.GetEnrichedData()
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

func TestEnricherFailsIfAdminFileFails(t *testing.T) {
	data, _ := getData()

	enricher := NewDataEnricher(
		data,
		getAdminReader("admin1"),
		getAntennaReader("antenna"),
		getBroadcasterReader("broadcas"),
		getFmOrgReader("fmorg"),
		getLanguageReader("language"),
		getSiteReader("site"),
	)

	items, err := enricher.GetEnrichedData()
	assert.NotNil(t, err)
	assert.Equal(t, "open resources/admin1.txt: no such file or directory", err.Error())
	assert.Nil(t, items)
}

func TestEnricherFailsIfAntennasFileFails(t *testing.T) {
	data, _ := getData()

	enricher := NewDataEnricher(
		data,
		getAdminReader("admin"),
		getAntennaReader("antenna1"),
		getBroadcasterReader("broadcas"),
		getFmOrgReader("fmorg"),
		getLanguageReader("language"),
		getSiteReader("site"),
	)

	items, err := enricher.GetEnrichedData()
	assert.NotNil(t, err)
	assert.Equal(t, "open resources/antenna1.txt: no such file or directory", err.Error())
	assert.Nil(t, items)
}

func TestEnricherFailsIfBroadcasFileFails(t *testing.T) {
	data, _ := getData()

	enricher := NewDataEnricher(
		data,
		getAdminReader("admin"),
		getAntennaReader("antenna"),
		getBroadcasterReader("broadcas1"),
		getFmOrgReader("fmorg"),
		getLanguageReader("language"),
		getSiteReader("site"),
	)

	items, err := enricher.GetEnrichedData()
	assert.NotNil(t, err)
	assert.Equal(t, "open resources/broadcas1.txt: no such file or directory", err.Error())
	assert.Nil(t, items)
}

func TestEnricherFailsIfFmOrgFileFails(t *testing.T) {
	data, _ := getData()

	enricher := NewDataEnricher(
		data,
		getAdminReader("admin"),
		getAntennaReader("antenna"),
		getBroadcasterReader("broadcas"),
		getFmOrgReader("fmorg1"),
		getLanguageReader("language"),
		getSiteReader("site"),
	)

	items, err := enricher.GetEnrichedData()
	assert.NotNil(t, err)
	assert.Equal(t, "open resources/fmorg1.txt: no such file or directory", err.Error())
	assert.Nil(t, items)
}
func TestEnricherFailsIfLanguagesFileFails(t *testing.T) {
	data, _ := getData()

	enricher := NewDataEnricher(
		data,
		getAdminReader("admin"),
		getAntennaReader("antenna"),
		getBroadcasterReader("broadcas"),
		getFmOrgReader("fmorg"),
		getLanguageReader("language1"),
		getSiteReader("site"),
	)

	items, err := enricher.GetEnrichedData()
	assert.NotNil(t, err)
	assert.Equal(t, "open resources/language1.txt: no such file or directory", err.Error())
	assert.Nil(t, items)
}
func TestEnricherFailsIfSiteFileFails(t *testing.T) {
	data, _ := getData()

	enricher := NewDataEnricher(
		data,
		getAdminReader("admin"),
		getAntennaReader("antenna"),
		getBroadcasterReader("broadcas"),
		getFmOrgReader("fmorg"),
		getLanguageReader("language"),
		getSiteReader("site1"),
	)

	items, err := enricher.GetEnrichedData()
	assert.NotNil(t, err)
	assert.Equal(t, "open resources/site1.txt: no such file or directory", err.Error())
	assert.Nil(t, items)
}

func getEnricher() (*DataEnricher, error) {
	data, err := getData()
	if err != nil {
		return nil, err
	}

	enricher := NewDataEnricher(
		data,
		getAdminReader("admin"),
		getAntennaReader("antenna"),
		getBroadcasterReader("broadcas"),
		getFmOrgReader("fmorg"),
		getLanguageReader("language"),
		getSiteReader("site"),
	)

	return &enricher, nil
}

func getData() (*Data, error) {
	provider := NewDataProvider(
		getProgramReader("test_hfcc_format_file"),
	)

	data, err := provider.GetData()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func getAdminReader(name string) reader.FileReader {
	processor := admin.NewAdminFileReader()

	return getFileReader(name, &processor)
}

func getAntennaReader(name string) reader.FileReader {
	processor := antenna.NewAntennaFileReader()

	return getFileReader(name, &processor)
}

func getBroadcasterReader(name string) reader.FileReader {
	processor := broadcaster.NewBroadcasterFileReader()

	return getFileReader(name, &processor)
}

func getFmOrgReader(name string) reader.FileReader {
	processor := fmOrg.NewFmOrgFileReader()

	return getFileReader(name, &processor)
}

func getLanguageReader(name string) reader.FileReader {
	processor := language.NewLanguageFileReader()

	return getFileReader(name, &processor)
}

func getProgramReader(name string) reader.FileReader {
	processor := program.NewProgramFileReader()

	return getFileReader(name, &processor)
}

func getSiteReader(name string) reader.FileReader {
	processor := site.NewSiteFileReader()

	return getFileReader(name, &processor)
}

func getFileReader(name string, processor reader.LineProcessor) reader.FileReader {
	return reader.NewFileReader(
		"resources/"+name+".txt",
		reader.NewLineReader(),
		processor,
		nil,
	)
}
