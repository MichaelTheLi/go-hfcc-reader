package reader

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type FileReader struct {
	filePath string
	items    []RawDataItem
}

// Example:
// ;----+----+----+------------------------------+---+----+-------+---+---+-------+------+------+-+-----+----------+---+---+---+-----+-+-----+-----+-----+-------
// ;FREQ STRT STOP CIRAF ZONES                    LOC POWR AZIMUTH SLW ANT DAYS    FDATE  TDATE MOD AFRQ LANGUAGE   ADM BRC FMO REQ# OLD ALT1 ALT2  ALT3  NOTES
// ;----+----+----+------------------------------+---+----+-------+---+---+-------+------+------+-+-----+----------+---+---+---+-----+-+-----+-----+-----+-------
//
//	2485 1000 1900 56,51                          PVL   10 0         0 400 1234567 310324 271024 D  9000 Bis        VUT VBT RNZ  1022                     NZL
type RawDataItem struct {
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

func NewFileReader(filePath string) FileReader {
	return FileReader{
		filePath: filePath,
		items:    []RawDataItem{},
	}
}

func (source FileReader) GetRawItems() []RawDataItem {
	file, err := os.Open(source.filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var started = false
	for scanner.Scan() {
		var text = scanner.Text()

		if strings.HasPrefix(text, ";FREQ") {
			// Skip the first line and continue
			scanner.Scan()
			started = true
			continue
		}

		if started {
			dataItem := source.getDataItem(text)
			source.items = append(source.items, dataItem)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return source.items
}

func (source FileReader) getDataItem(line string) RawDataItem {
	var dataItem RawDataItem
	valueReflection := reflect.ValueOf(&dataItem)

	for fieldInd := 0; fieldInd < valueReflection.Elem().Type().NumField(); fieldInd++ {
		field := valueReflection.Elem().Type().Field(fieldInd)

		start, startErr := strconv.Atoi(field.Tag.Get("start"))
		if startErr != nil {
			panic(startErr)
		}
		end, endErr := strconv.Atoi(field.Tag.Get("end"))
		if endErr != nil {
			panic(endErr)
		}
		fieldValue := trimSubstr(line, start-1, end)
		valueField := valueReflection.Elem().Field(fieldInd)
		valueField.SetString(fieldValue)
	}
	return dataItem
}

func trimSubstr(input string, start int, length int) string {
	return strings.Trim(
		substr(input, start, length),
		" \x00",
	)
}

func substr(input string, start int, end int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	return string(asRunes[start:end])
}
