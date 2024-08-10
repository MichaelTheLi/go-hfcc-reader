package reader

import (
	"bufio"
	"fmt"
	"os"
)

type FileReader struct {
	filePath      string
	lineReader    LineReader
	lineProcessor LineProcessor
}

type LineProcessor interface {
	ProcessLine(index int, text string) interface{}
}

func NewFileReader(filePath string, lineReader LineReader, lineProcessor LineProcessor) FileReader {
	return FileReader{
		filePath:      filePath,
		lineReader:    lineReader,
		lineProcessor: lineProcessor,
	}
}

func (source FileReader) ProcessFile() LineProcessor {
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

	index := 0
	lineProcessor := source.lineProcessor
	for scanner.Scan() {
		var text = scanner.Text()
		dataItem := lineProcessor.ProcessLine(index, text)
		if dataItem != nil {
			source.lineReader.fillDataItem(text, dataItem)
		}

		index += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return lineProcessor
}
