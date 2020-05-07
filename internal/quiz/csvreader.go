package quiz

import (
	"encoding/csv"
	"io"
	"log"
)

// CSVReader implement reader.Reader
type CSVReader struct {
}

// NewCSVReader return new CSVReader
func NewCSVReader() Reader {
	return &CSVReader{}
}

// Read represent read
func (r *CSVReader) Read(rdr io.Reader) []*Problem {
	csvReader := csv.NewReader(rdr)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("error when reading records, %v", err)
	}

	problems := make([]*Problem, len(records))
	for i, record := range records {
		problems[i] = &Problem{question: record[0], answer: record[1]}
	}

	return problems
}
