package language

type Reader struct {
	RawLanguageData RawLanguageData
}

type RawLanguageData struct {
	Items    []*RawLanguageDataItem
	Metadata RawLanguageMetadata
}

type RawLanguageMetadata struct {
	Date string `start:"11" end:"22"`
	Name string `start:"24"`
}

// RawLanguageDataItem Example:
// ;         23-JAN-2014  Reference Table Language
// ;
// Aaa Ghotuo
// Aab Alumu-Tesu
type RawLanguageDataItem struct {
	Code        string `start:"1" end:"3"`
	EnglishName string `start:"5"`
}

func NewLanguageFileReader() Reader {
	return Reader{
		RawLanguageData: RawLanguageData{
			Items: []*RawLanguageDataItem{},
		},
	}
}

func (source *Reader) ProcessLine(index int, text string) interface{} {
	if index == 0 {
		return &source.RawLanguageData.Metadata
	} else if text[0] != ';' {
		Language := RawLanguageDataItem{}
		source.RawLanguageData.Items = append(source.RawLanguageData.Items, &Language)
		return &Language
	}

	return nil
}
