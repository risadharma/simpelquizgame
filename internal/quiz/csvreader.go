package quiz

import (
	"encoding/csv"
	"io"
	"os"
)

// CSVReader implement reader.Reader
type CSVReader struct {
	filepath string
}

// NewCSVReader return new CSVReader
func NewCSVReader(filepath string) *CSVReader {
	return &CSVReader{
		filepath: filepath,
	}
}

// Read represent read
func (r *CSVReader) Read() ([]*Problem, error) {
	csvFile, err := os.Open(r.filepath)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(csvFile)

	var problems []*Problem

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		problems = append(problems, &Problem{Question: record[0], Answer: record[1]})
	}

	return problems, nil
}
