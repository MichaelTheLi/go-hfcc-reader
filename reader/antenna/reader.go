package antenna

type Reader struct {
	RawAntennaData RawAntennaData
}

type RawAntennaData struct {
	Items    []*RawAntennaDataItem
	Metadata RawAntennaMetadata
}

type RawAntennaMetadata struct {
	Date string `start:"11" end:"22"`
	Name string `start:"23"`
}

// RawAntennaDataItem Example:
// ;         29-MAY-2019 ANTENNA.TXT REFERENCE TABLE
// ;--+-------------------------------------------------+--------------------+
// ;ANT     ANTENNA DEFINITION                                   REMARKS
// ;CODE
// ;--+-------------------------------------------------+--------------------+
// 100 AHR1/1/0.3
// 101 AHR1/1/0.5

type RawAntennaDataItem struct {
	Code       string `start:"1" end:"3"`
	Definition string `start:"5" end:"54"`
	Notes      string `start:"55"`
}

func NewAntennaFileReader() Reader {
	return Reader{
		RawAntennaData: RawAntennaData{
			Items: []*RawAntennaDataItem{},
		},
	}
}

func (source *Reader) ProcessLine(index int, text string) interface{} {
	if index == 0 {
		return &source.RawAntennaData.Metadata
	} else if text[0] != ';' {
		Antenna := RawAntennaDataItem{}
		source.RawAntennaData.Items = append(source.RawAntennaData.Items, &Antenna)
		return &Antenna
	}

	return nil
}
