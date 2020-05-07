package quiz

import (
	"encoding/csv"
	"fmt"
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
		return nil, fmt.Errorf("error when opening file, %v", err)
	}

	csvReader := csv.NewReader(csvFile)

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error when reading records, %v", err)
	}

	problems := make([]*Problem, len(records))
	for i, record := range records {
		problems[i] = &Problem{Question: record[0], Answer: record[1]}
	}

	return problems, nil
}
