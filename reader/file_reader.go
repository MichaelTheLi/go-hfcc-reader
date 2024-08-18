package reader

import (
	"bufio"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"os"
)

type FileReader struct {
	filePath      string
	enc           *charmap.Charmap
	lineReader    LineReader
	lineProcessor LineProcessor
}

type LineProcessor interface {
	ProcessLine(index int, text string) interface{}
}

func NewFileReader(filePath string, lineReader LineReader, lineProcessor LineProcessor, enc *charmap.Charmap) FileReader {
	return FileReader{
		filePath:      filePath,
		lineReader:    lineReader,
		lineProcessor: lineProcessor,
		enc:           enc,
	}
}

func (source FileReader) ProcessFile() (LineProcessor, error) {
	file, err := os.Open(source.filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	var scanner *bufio.Scanner
	if source.enc != nil {
		dec := transform.NewReader(file, source.enc.NewDecoder())
		scanner = bufio.NewScanner(dec)
	} else {
		scanner = bufio.NewScanner(file)
	}

	index := 0
	lineProcessor := source.lineProcessor
	for scanner.Scan() {
		var text = scanner.Text()
		dataItem := lineProcessor.ProcessLine(index, text)
		if dataItem != nil {
			lineErr := source.lineReader.fillDataItem(text, dataItem)
			if lineErr != nil {
				return nil, lineErr
			}
		}

		index += 1
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lineProcessor, nil
}
