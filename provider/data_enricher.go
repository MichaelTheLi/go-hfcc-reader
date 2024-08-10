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
	Data         Data
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
	data Data,
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

func (source DataEnricher) GetEnrichedData() map[string]EnrichedDataItem {
	administration := source.prepareAdmin()
	antennas := source.prepareAntenna()
	broadcasters := source.prepareBroadcaster()
	fmOrgs := source.prepareFmOrg()
	languages := source.prepareLanguages()
	sites := source.prepareSite()

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

	return source.EnrichedData
}

func (source DataEnricher) prepareAdmin() map[string]*admin.RawAdminDataItem {
	data := source.admin.ProcessFile().(*admin.Reader).RawAdminData

	result := make(map[string]*admin.RawAdminDataItem)

	for _, item := range data.Items {
		result[item.Code] = item
	}

	return result
}

func (source DataEnricher) prepareAntenna() map[string]*antenna.RawAntennaDataItem {
	data := source.antenna.ProcessFile().(*antenna.Reader).RawAntennaData

	result := make(map[string]*antenna.RawAntennaDataItem)

	for _, item := range data.Items {
		result[item.Code] = item
	}

	return result
}

func (source DataEnricher) prepareBroadcaster() map[string]*broadcaster.RawBroadcasterDataItem {
	data := source.broadcaster.ProcessFile().(*broadcaster.Reader).RawBroadcasterData

	result := make(map[string]*broadcaster.RawBroadcasterDataItem)

	for _, item := range data.Items {
		result[item.Code] = item
	}

	return result
}

func (source DataEnricher) prepareFmOrg() map[string]*fmOrg.RawFmOrgDataItem {
	data := source.fmorg.ProcessFile().(*fmOrg.Reader).RawFmOrgData

	result := make(map[string]*fmOrg.RawFmOrgDataItem)

	for _, item := range data.Items {
		result[item.Code] = item
	}

	return result
}

func (source DataEnricher) prepareLanguages() map[string]*language.RawLanguageDataItem {
	data := source.language.ProcessFile().(*language.Reader).RawLanguageData

	result := make(map[string]*language.RawLanguageDataItem)

	for _, item := range data.Items {
		result[item.Code] = item
	}

	return result
}

func (source DataEnricher) prepareSite() map[string]*site.RawSiteDataItem {
	data := source.site.ProcessFile().(*site.Reader).RawSiteData

	result := make(map[string]*site.RawSiteDataItem)

	for _, item := range data.Items {
		result[item.Code] = item
	}

	return result
}
