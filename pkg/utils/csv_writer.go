package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteCSVFile(filePath string, records [][]string) (err error) {

	file, err := os.OpenFile(filePath, os.O_APPEND, 2)
	if err != nil {
		log.Fatal("ReadCSVFile(): Unable to open file "+filePath, err)
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	w.WriteAll(records)

	return
}
