package common

import (
	"encoding/csv"
	"os"
)

// ReadCsvFile - Returns data from `filename`
func ReadCsvFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return make([][]string, 0), err
	}

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

// WriteToCsv - Writes `data` to `filename`
func WriteToCsv(filename string, data []string) error {
	// perms that work => 0644, 066, 0755
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return err
	}

	w := csv.NewWriter(file)
	err = w.WriteAll([][]string{data})
	if err != nil {
		return err
	}

	return nil
}
