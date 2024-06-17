package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCSVFile(filePath string) ([][]string, error) {

	pwd, _ := os.Getwd()
	f, err := os.Open(pwd + filePath)
	if err != nil {
		log.Fatal("ReadCSVFile(): Unable to read input file ", err)
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("ReadCSVFile(): Unable to parse file as CSV for ", err)
		return nil, err
	}
	return records, nil
}
