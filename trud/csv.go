package trud

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"reflect"
)

// ProcessCSV reads a comma-delimited file and calls a handler for a batch of rows.
func ProcessCSV(reader io.Reader, headings []string, batchSize int, logger *log.Logger, processFunc func(rows [][]string) error) error {
	r := csv.NewReader(bufio.NewReader(reader))
	if headings != nil {
		r.FieldsPerRecord = len(headings)
	}
	batch := make([][]string, 0, batchSize)
	for {
		record, err := r.Read()
		if err == io.EOF { // stop at EOF
			break
		}
		if headings != nil {
			if !reflect.DeepEqual(headings, record) {
				return fmt.Errorf("expecting column names: %v, got: %v", headings, record)
			}
			headings = nil
		} else {
			batch = append(batch, record)
			if len(batch) == batchSize {
				err = processFunc(batch)
				if err != nil {
					return err
				}
				batch = nil
			}
		}
	}
	if len(batch) > 0 {
		processFunc(batch)
	}
	return nil
}
