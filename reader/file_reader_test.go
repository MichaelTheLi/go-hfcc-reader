package reader

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/charmap"
	"os"
	"testing"
)

type Reader struct {
	RawTestData RawTestData
}

type RawTestData struct {
	Items     []*RawTestDataItem
	Metadata  RawTestMetadata
	Metadata2 RawTestMetadata2
}

type RawTestMetadata struct {
	Something1 string `start:"2" end:"12"`
	Something2 string `start:"14" end:"23"`
	Something3 string `start:"25" end:"34"`
}

type RawTestMetadata2 struct {
	Something1 string `start:"2" end:"19"`
	Something2 string `start:"21" end:"37"`
}

type RawTestDataItem struct {
	One  string `start:"1" end:"3"`
	Two  string `start:"5" end:"34"`
	Rest string `start:"35" end:"105"`
}

func (source *Reader) ProcessLine(index int, text string) interface{} {
	if index == 0 {
		return &source.RawTestData.Metadata
	} else if index == 1 {
		return &source.RawTestData.Metadata2
	} else if text[0] != ';' {
		Item := RawTestDataItem{}
		source.RawTestData.Items = append(source.RawTestData.Items, &Item)
		return &Item
	}
	return nil
}

func TestNewFileReader(t *testing.T) {
	reader := getFileReader()
	assert.IsType(t, FileReader{}, reader)
}

func TestFileReaderCanProcessFileMetadata1(t *testing.T) {
	reader := getFileReader()
	processor, _ := reader.ProcessFile()
	testData := processor.(*Reader).RawTestData

	assert.Equal(t, "Something1", testData.Metadata.Something1)
	assert.Equal(t, "Something2", testData.Metadata.Something2)
	assert.Equal(t, "Something3", testData.Metadata.Something3)
}

func TestFileReaderCanProcessFileMetadata2(t *testing.T) {
	reader := getFileReader()
	processor, _ := reader.ProcessFile()
	testData := processor.(*Reader).RawTestData

	assert.Equal(t, "AnotherSomething1", testData.Metadata2.Something1)
	assert.Equal(t, "AnotherSomething2", testData.Metadata2.Something2)
}

func TestFileReaderCanProcessFileItems(t *testing.T) {
	reader := getFileReader()
	processor, _ := reader.ProcessFile()
	testData := processor.(*Reader).RawTestData

	assert.Len(t, testData.Items, 3)

	assert.Equal(t, "A-A", testData.Items[0].One)
	assert.Equal(t, "Alma Ata", testData.Items[0].Two)
	assert.Equal(t, "Some notes", testData.Items[0].Rest)
}

func TestFileInEncodingReads(t *testing.T) {
	path, _ := os.Getwd()
	reader := NewFileReader(
		path+"/resources/test_file_iso-8859-1.txt",
		NewLineReader(),
		&Reader{},
		charmap.ISO8859_1,
	)
	processor, err := reader.ProcessFile()

	assert.Empty(t, err)
	assert.NotEmpty(t, processor)
	assert.Len(t, processor.(*Reader).RawTestData.Items, 3)
	assert.Equal(t, "Some notes é", processor.(*Reader).RawTestData.Items[0].Rest)
}

func TestInvalidFileFails(t *testing.T) {
	reader := NewFileReader(
		"non_existent_file.txt",
		NewLineReader(),
		&Reader{},
		nil,
	)
	processor, err := reader.ProcessFile()

	assert.NotEmpty(t, err)
	assert.Empty(t, processor)
}

type InvalidReader struct {
	InvalidRawTestData InvalidRawTestData
}

type InvalidRawTestData struct {
	InvalidMetadata InvalidRawTestMetadata
}

type InvalidRawTestMetadata struct {
	Something1 string `start:"absag"`
}

func (source *InvalidReader) ProcessLine(index int, _ string) interface{} {
	if index == 0 {
		return &source.InvalidRawTestData.InvalidMetadata
	}

	return nil
}

func TestInvalidDataItemFails(t *testing.T) {
	path, _ := os.Getwd()
	reader := NewFileReader(
		path+"/resources/test_file.txt",
		NewLineReader(),
		&InvalidReader{},
		nil,
	)
	_, err := reader.ProcessFile()

	assert.NotNil(t, err)
	assert.Equal(t, "invalid start tag for the Something1: strconv.Atoi: parsing \"absag\": invalid syntax", err.Error())
}

func getFileReader() FileReader {
	path, _ := os.Getwd()
	reader := NewFileReader(
		path+"/resources/test_file.txt",
		NewLineReader(),
		&Reader{},
		nil,
	)
	return reader
}
