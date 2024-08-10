package site

type Reader struct {
	RawSiteData RawSiteData
}

type RawSiteData struct {
	Items    []*RawSiteDataItem
	Metadata RawSiteMetadata
}

type RawSiteMetadata struct {
	Date string `start:"11" end:"22"`
	Name string `start:"23"`
}

// RawSiteDataItem Example:
// ;         11-May-2023   Global HF Transmitter Site Table
// ;
// ;--+------------------------------+---+-----+------
// ;Co Site Name                      ADM Lati  Longi
// ;de                                    tude  tude
// ;--+------------------------------+---+-----+------
// A-A Alma Ata                       KAZ 43N17 077E00

type RawSiteDataItem struct {
	Code           string `start:"1" end:"3"`
	EnglishName    string `start:"4" end:"35"`
	Administration string `start:"36" end:"39"`
	Latitude       string `start:"40" end:"44"`
	Longitude      string `start:"46"`
}

func NewSiteFileReader() Reader {
	return Reader{
		RawSiteData: RawSiteData{
			Items: []*RawSiteDataItem{},
		},
	}
}

func (source *Reader) ProcessLine(index int, text string) interface{} {
	if index == 0 {
		return &source.RawSiteData.Metadata
	} else if text[0] != ';' {
		Site := RawSiteDataItem{}
		source.RawSiteData.Items = append(source.RawSiteData.Items, &Site)
		return &Site
	}

	return nil
}
