package reader

type ProgramFileReader struct {
	RawProgramsData RawProgramsData
}

type RawProgramsData struct {
	Items    []*RawProgram
	Metadata RawProgramsMetadata
	Note1    Note
	Note2    Note
	Note3    Note
}

type Note struct {
	Value string `start:"3"`
}

type RawProgramsMetadata struct {
	Season         string `start:"3" end:"5"`
	Administration string `start:"7" end:"9"`
	Date           string `start:"11" end:"21"`
}

// Example:
// ;----+----+----+------------------------------+---+----+-------+---+---+-------+------+------+-+-----+----------+---+---+---+-----+-+-----+-----+-----+-------
// ;FREQ STRT STOP CIRAF ZONES                    LOC POWR AZIMUTH SLW ANT DAYS    FDATE  TDATE MOD AFRQ LANGUAGE   ADM BRC FMO REQ# OLD ALT1 ALT2  ALT3  NOTES
// ;----+----+----+------------------------------+---+----+-------+---+---+-------+------+------+-+-----+----------+---+---+---+-----+-+-----+-----+-----+-------
//
//	2485 1000 1900 56,51                          PVL   10 0         0 400 1234567 310324 271024 D  9000 Bis        VUT VBT RNZ  1022                     NZL
type RawProgram struct {
	Frequency              string `start:"1" end:"5"`
	StartTime              string `start:"7" end:"10"`
	EndTime                string `start:"12" end:"15"`
	CIRAF                  string `start:"17" end:"46"`
	Location               string `start:"48" end:"50"`
	Power                  string `start:"52" end:"55"`
	Azimuth                string `start:"57" end:"63"`
	AntennaSlewAngle       string `start:"65" end:"67"`
	Antenna                string `start:"69" end:"71"`
	DaysActive             string `start:"73" end:"79"`
	StartDate              string `start:"81" end:"86"`
	EndDate                string `start:"88" end:"93"`
	Modulation             string `start:"95" end:"95"`
	AntennaDesignFrequency string `start:"97" end:"101"`
	Language               string `start:"103" end:"112"`
	Administration         string `start:"114" end:"116"`
	Broadcaster            string `start:"118" end:"120"`
	FmOrgId                string `start:"122" end:"124"`
	Id                     string `start:"126" end:"130"`
	OldData                string `start:"132" end:"132"`
	Alt1                   string `start:"134" end:"138"`
	Alt2                   string `start:"140" end:"144"`
	Alt3                   string `start:"146" end:"150"`
	Notes                  string `start:"152" end:"158"`
}

func NewProgramFileReader() ProgramFileReader {
	return ProgramFileReader{
		RawProgramsData: RawProgramsData{
			Items:    []*RawProgram{},
			Metadata: RawProgramsMetadata{},
			Note1:    Note{},
			Note2:    Note{},
			Note3:    Note{},
		},
	}
}

func (source *ProgramFileReader) ProcessLine(index int, _ string) interface{} {
	if index == 0 {
		return &source.RawProgramsData.Metadata
	} else if index >= 1 && index <= 5 {
		if index == 1 {
			return &source.RawProgramsData.Note1
		} else if index == 2 {
			return &source.RawProgramsData.Note2
		} else if index == 3 {
			return &source.RawProgramsData.Note3
		}
	} else if index > 6 {
		program := RawProgram{}
		source.RawProgramsData.Items = append(source.RawProgramsData.Items, &program)
		return &program
	}

	return nil
}
