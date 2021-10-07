package common

import (
	"encoding/csv"
	"fmt"
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
func WriteToCsv(filename string, data [][]string) {
	// 0644, 066, 0755
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println(err)
	}

	w := csv.NewWriter(file)
	w.WriteAll(data)

	if err := w.Error(); err != nil {
		fmt.Println("Error", err)
	}
}
