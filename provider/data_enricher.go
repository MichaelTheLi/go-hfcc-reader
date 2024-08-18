package provider

import (
	"github.com/MichaelTheLi/go-hfcc-reader/reader"
	"github.com/MichaelTheLi/go-hfcc-reader/reader/admin"
	"github.com/MichaelTheLi/go-hfcc-reader/reader/antenna"
	"github.com/MichaelTheLi/go-hfcc-reader/reader/broadcaster"
	fmOrg "github.com/MichaelTheLi/go-hfcc-reader/reader/fmorg"
	"github.com/MichaelTheLi/go-hfcc-reader/reader/language"
	"github.com/MichaelTheLi/go-hfcc-reader/reader/site"
	"strconv"
)

type DataEnricher struct {
	Data         *Data
	admin        reader.FileReader
	antenna      reader.FileReader
	broadcaster  reader.FileReader
	fmorg        reader.FileReader
	language     reader.FileReader
	site         reader.FileReader
	EnrichedData map[string]EnrichedDataItem
}

type EnrichedDataItem struct {
	Item        DataItem
	Admin       Admin
	Antenna     Antenna
	Broadcaster Broadcaster
	FmOrg       FmOrg
	Language    Language
	Site        Site
}

type Admin struct {
	admin.RawAdminDataItem
}
type Antenna struct {
	antenna.RawAntennaDataItem
}
type Broadcaster struct {
	broadcaster.RawBroadcasterDataItem
}
type FmOrg struct {
	fmOrg.RawFmOrgDataItem
}
type Language struct {
	language.RawLanguageDataItem
}
type Site struct {
	site.RawSiteDataItem
}

func NewDataEnricher(
	data *Data,
	admin reader.FileReader,
	antenna reader.FileReader,
	broadcaster reader.FileReader,
	fmorg reader.FileReader,
	language reader.FileReader,
	site reader.FileReader,
) DataEnricher {
	return DataEnricher{
		Data:         data,
		admin:        admin,
		antenna:      antenna,
		broadcaster:  broadcaster,
		fmorg:        fmorg,
		language:     language,
		site:         site,
		EnrichedData: make(map[string]EnrichedDataItem),
	}
}

func (source DataEnricher) GetEnrichedData() (map[string]EnrichedDataItem, error) {
	administration, adminErr := source.prepareAdmin()
	if adminErr != nil {
		return nil, adminErr
	}
	antennas, antennasErr := source.prepareAntenna()
	if antennasErr != nil {
		return nil, antennasErr
	}
	broadcasters, broadcastersErr := source.prepareBroadcaster()
	if broadcastersErr != nil {
		return nil, broadcastersErr
	}
	fmOrgs, fmOrgsErr := source.prepareFmOrg()
	if fmOrgsErr != nil {
		return nil, fmOrgsErr
	}
	languages, languagesErr := source.prepareLanguages()
	if languagesErr != nil {
		return nil, languagesErr
	}
	sites, sitesErr := source.prepareSite()
	if sitesErr != nil {
		return nil, sitesErr
	}

	for _, item := range source.Data.ProgramsList {
		administrationItem, _ := administration[string(item.Administration)]
		antennaItem, _ := antennas[strconv.Itoa(int(item.Antenna))]
		broadcasterItem, _ := broadcasters[string(item.Broadcaster)]
		fmOrgItem, _ := fmOrgs[string(item.FmOrgId)]
		lngItem, _ := languages[string(item.Language)]
		siteItem, _ := sites[string(item.Location)]

		dataItem := EnrichedDataItem{
			Item:        item,
			Admin:       Admin{*administrationItem},
			Antenna:     Antenna{*antennaItem},
			Broadcaster: Broadcaster{*broadcasterItem},
			FmOrg:       FmOrg{*fmOrgItem},
			Language:    Language{*lngItem},
			Site:        Site{*siteItem},
		}
		source.EnrichedData[dataItem.Item.Id] = dataItem
	}

	return source.EnrichedData, nil
}

func (source DataEnricher) prepareAdmin() (map[string]*admin.RawAdminDataItem, error) {
	rawData, err := source.admin.ProcessFile()
	if err != nil {
		return nil, err
	}
	data := rawData.(*admin.Reader).RawAdminData

	result := make(map[string]*admin.RawAdminDataItem)

	for _, item := range data.Items {
		result[item.Code] = item
	}

	return result, nil
}

func (source DataEnricher) prepareAntenna() (map[string]*antenna.RawAntennaDataItem, error) {
	rawData, err := source.antenna.ProcessFile()
	if err != nil {
		return nil, err
	}
	data := rawData.(*antenna.Reader).RawAntennaData

	result := make(map[string]*antenna.RawAntennaDataItem)

	for _, item := range data.Items {
		result[item.Code] = item
	}

	return result, nil
}

func (source DataEnricher) prepareBroadcaster() (map[string]*broadcaster.RawBroadcasterDataItem, error) {
	rawData, err := source.broadcaster.ProcessFile()
	if err != nil {
		return nil, err
	}
	data := rawData.(*broadcaster.Reader).RawBroadcasterData

	result := make(map[string]*broadcaster.RawBroadcasterDataItem)

	for _, item := range data.Items {
		result[item.Code] = item
	}

	return result, nil
}

func (source DataEnricher) prepareFmOrg() (map[string]*fmOrg.RawFmOrgDataItem, error) {
	rawData, err := source.fmorg.ProcessFile()
	if err != nil {
		return nil, err
	}
	data := rawData.(*fmOrg.Reader).RawFmOrgData

	result := make(map[string]*fmOrg.RawFmOrgDataItem)

	for _, item := range data.Items {
		result[item.Code] = item
	}

	return result, nil
}

func (source DataEnricher) prepareLanguages() (map[string]*language.RawLanguageDataItem, error) {
	rawData, err := source.language.ProcessFile()
	if err != nil {
		return nil, err
	}
	data := rawData.(*language.Reader).RawLanguageData

	result := make(map[string]*language.RawLanguageDataItem)

	for _, item := range data.Items {
		result[item.Code] = item
	}

	return result, nil
}

func (source DataEnricher) prepareSite() (map[string]*site.RawSiteDataItem, error) {
	rawData, err := source.site.ProcessFile()
	if err != nil {
		return nil, err
	}
	data := rawData.(*site.Reader).RawSiteData

	result := make(map[string]*site.RawSiteDataItem)

	for _, item := range data.Items {
		result[item.Code] = item
	}

	return result, nil
}
