package reader

import (
	"reflect"
	"strconv"
	"strings"
)

type LineReader struct {
}

func NewLineReader() LineReader {
	return LineReader{}
}

func (lineReader LineReader) fillDataItem(line string, dataItem interface{}) {
	interfaceReflection := reflect.ValueOf(dataItem)
	valueReflection := reflect.Indirect(interfaceReflection)

	for fieldInd := 0; fieldInd < valueReflection.Type().NumField(); fieldInd++ {
		field := valueReflection.Type().Field(fieldInd)

		start, startErr := strconv.Atoi(field.Tag.Get("start"))
		if startErr != nil {
			panic(startErr)
		}
		end, endErr := strconv.Atoi(field.Tag.Get("end"))
		if endErr != nil {
			end = len(line)
		}
		fieldValue := trimSubstr(line, start-1, end)
		valueField := valueReflection.Field(fieldInd)
		valueField.SetString(fieldValue)
	}
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
