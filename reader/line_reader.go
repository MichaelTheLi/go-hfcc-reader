package reader

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type LineReader struct {
}

func NewLineReader() LineReader {
	return LineReader{}
}

func (lineReader LineReader) fillDataItem(line string, dataItem interface{}) error {
	interfaceReflection := reflect.ValueOf(dataItem)
	valueReflection := reflect.Indirect(interfaceReflection)

	for fieldInd := 0; fieldInd < valueReflection.Type().NumField(); fieldInd++ {
		field := valueReflection.Type().Field(fieldInd)

		start, startErr := strconv.Atoi(field.Tag.Get("start"))
		if startErr != nil {
			return errors.New("invalid start tag for the " + field.Name + ": " + startErr.Error())
		}
		end, endErr := strconv.Atoi(field.Tag.Get("end"))
		if endErr != nil {
			end = len(line)
		}
		if start > end {
			return errors.New("invalid start or end for the " + field.Name + " field: start should be lower than end")
		}

		fieldValue := trimSubstr(line, start-1, end)
		valueField := valueReflection.Field(fieldInd)
		valueField.SetString(fieldValue)
	}

	return nil
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

	end = min(end, len(asRunes))

	return string(asRunes[start:end])
}
