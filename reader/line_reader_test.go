package reader

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testItemValid struct {
	One    string `start:"1" end:"2"`
	Two    string `start:"3" end:"5"`
	Single string `start:"6" end:"6"`
	Rest   string `start:"7"`
}

func TestNewLineReader(t *testing.T) {
	reader := NewLineReader()
	assert.IsType(t, LineReader{}, reader)
}

func TestLineReaderReadsValidStringWithoutError(t *testing.T) {
	reader := NewLineReader()
	testString := "123456789ABCDEF"
	testItemObj := testItemValid{}
	err := reader.fillDataItem(testString, &testItemObj)

	assert.Empty(t, err)
}

func TestLineReaderValidStringCorrectlyStructured(t *testing.T) {
	reader := NewLineReader()
	testString := "123456789ABCDEF"
	testItemObj := testItemValid{}
	_ = reader.fillDataItem(testString, &testItemObj)

	assert.Equal(t, "12", testItemObj.One)
	assert.Equal(t, "345", testItemObj.Two)
	assert.Equal(t, "6", testItemObj.Single)
	assert.Equal(t, "789ABCDEF", testItemObj.Rest)
}

type testItemInvalidStart struct {
	One string `end:"2"`
}

func TestLineReaderInvalidStartTagFailed(t *testing.T) {
	reader := NewLineReader()
	testString := "123456789ABCDEF"
	testItemObj := testItemInvalidStart{}
	err := reader.fillDataItem(testString, &testItemObj)

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("invalid start tag for the One: strconv.Atoi: parsing \"\": invalid syntax"), err)
	}
}

type testItemInvalidEndOrStart struct {
	One string `start:"5" end:"2"`
}

func TestLineReaderInvalidEndOrStartFails(t *testing.T) {
	reader := NewLineReader()
	testString := "123456789ABCDEF"
	testItemObj := testItemInvalidEndOrStart{}
	err := reader.fillDataItem(testString, &testItemObj)

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("invalid start or end for the One field: start should be lower than end"), err)
	}
}

type testItemInvalidFormat struct {
	One string `start:"asd" end:"2"`
}

func TestLineReaderInvalidFormatFails(t *testing.T) {
	reader := NewLineReader()
	testString := "123456789ABCDEF"
	testItemObj := testItemInvalidFormat{}
	err := reader.fillDataItem(testString, &testItemObj)

	if assert.Error(t, err) {
		assert.Equal(t, errors.New("invalid start tag for the One: strconv.Atoi: parsing \"asd\": invalid syntax"), err)
	}
}

func TestShortStringIgnored(t *testing.T) {
	reader := NewLineReader()
	testString := "1234"
	testItemObj := testItemValid{}
	_ = reader.fillDataItem(testString, &testItemObj)

	assert.Equal(t, "12", testItemObj.One)
	assert.Equal(t, "34", testItemObj.Two)
	assert.Equal(t, "", testItemObj.Single)
	assert.Equal(t, "", testItemObj.Rest)
}

type testItemTrimmedFormat struct {
	One string `start:"1" end:"5"`
	Two string `start:"6"`
}

func TestStringTrimmedIgnored(t *testing.T) {
	reader := NewLineReader()
	testString := "123    abc    "
	testItemObj := testItemTrimmedFormat{}
	_ = reader.fillDataItem(testString, &testItemObj)

	assert.Equal(t, "123", testItemObj.One)
	assert.Equal(t, "abc", testItemObj.Two)
}
