package common

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsvFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return make([][]string, 0), err
	}

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func WriteToCsv(filename string, data [][]string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}

	w := csv.NewWriter(file)
	w.WriteAll(data)

	if err := w.Error(); err != nil {
		fmt.Println("Error", err)
	}
}
