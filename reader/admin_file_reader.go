package reader

type AdminFileReader struct {
	RawAdminData RawAdminData
}

type RawAdminData struct {
	Items    []*RawAdminDataItem
	Metadata RawAdminMetadata
}

type RawAdminMetadata struct {
	Date string `start:"11" end:"21"`
	Name string `start:"23" end:"31"`
	Note string `start:"33" end:"47"`
}

// RawAdminDataItem Example:
// ;-- -------------------------------------------------- -------------------------------------------------- -------------------------------------------------
// ;CO ADMINISTRATION ENGLISH NAME                        ADMINISTRATION FRENCH NAME                         ADMINISTRATION SPANISH NAME
// ;DE
// ;-- -------------------------------------------------- -------------------------------------------------- -------------------------------------------------
// ARG Argentina                                          Argentine                                          Argentina
type RawAdminDataItem struct {
	Code        string `start:"1" end:"3"`
	EnglishName string `start:"5" end:"54"`
	FrenchName  string `start:"56" end:"105"`
	SpanishName string `start:"107" end:"156"`
}

func NewAdminFileReader() AdminFileReader {
	return AdminFileReader{
		RawAdminData: RawAdminData{
			Items: []*RawAdminDataItem{},
		},
	}
}

func (source *AdminFileReader) ProcessLine(index int, _ string) interface{} {
	if index == 0 {
		return &source.RawAdminData.Metadata
	} else if index >= 6 {
		Admin := RawAdminDataItem{}
		source.RawAdminData.Items = append(source.RawAdminData.Items, &Admin)
		return &Admin
	}

	return nil
}
