package provider

import (
	"github.com/MichaelTheLi/go-hfcc-reader/reader"
	"github.com/MichaelTheLi/go-hfcc-reader/reader/admin"
	"github.com/MichaelTheLi/go-hfcc-reader/reader/language"
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
	Item     DataItem
	Language Language
	Admin    Admin
}

type Language struct {
	language.RawLanguageDataItem
}
type Admin struct {
	admin.RawAdminDataItem
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
	languages := source.prepareLanguages()

	for _, item := range source.Data.ProgramsList {
		lng, _ := languages[string(item.Language)]
		administrationItem, _ := administration[string(item.Administration)]
		dataItem := EnrichedDataItem{
			Item:     item,
			Language: Language{*lng},
			Admin:    Admin{*administrationItem},
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

func (source DataEnricher) prepareLanguages() map[string]*language.RawLanguageDataItem {
	data := source.language.ProcessFile().(*language.Reader).RawLanguageData

	result := make(map[string]*language.RawLanguageDataItem)

	for _, item := range data.Items {
		result[item.Code] = item
	}

	return result
}
