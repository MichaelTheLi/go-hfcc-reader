package fmOrg

type Reader struct {
	RawFmOrgData RawFmOrgData
}

type RawFmOrgData struct {
	Items    []*RawFmOrgDataItem
	Metadata RawFmOrgMetadata
}

type RawFmOrgMetadata struct {
	Date string `start:"11" end:"22"`
	Name string `start:"23"`
}

// RawFmOrgDataItem Example:
// ;         25-JAN-2023  Reference Table Freq. Management Org.
// ;
// ;--+--------------------------------------------------+--------------------+-------------+-------------+----------------------------------------+-------------
// ;Co Frequency Management Organisation                  Contact Person       Telephone     Fax           E-mail                                   Notes
// ;de (Full name)
// ;--+--------------------------------------------------+--------------------+-------------+-------------+----------------------------------------+-------------
// ITU ITU
// ABC Australian Broadcasting Corporation
// ABU ABU-HFC Regional Coordination Group

type RawFmOrgDataItem struct {
	Code          string `start:"1" end:"3"`
	EnglishName   string `start:"5" end:"54"`
	ContactPerson string `start:"56" end:"75"`
	Telephone     string `start:"77" end:"89"`
	Fax           string `start:"91" end:"103"`
	Email         string `start:"105" end:"144"`
	Notes         string `start:"146"`
}

func NewFmOrgFileReader() Reader {
	return Reader{
		RawFmOrgData: RawFmOrgData{
			Items: []*RawFmOrgDataItem{},
		},
	}
}

func (source *Reader) ProcessLine(index int, text string) interface{} {
	if index == 0 {
		return &source.RawFmOrgData.Metadata
	} else if text[0] != ';' {
		FmOrg := RawFmOrgDataItem{}
		source.RawFmOrgData.Items = append(source.RawFmOrgData.Items, &FmOrg)
		return &FmOrg
	}

	return nil
}
