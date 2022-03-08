package main

import (
	"context"
	"encoding/csv"
	"errors"
	"os"
	"strings"

	"github.com/t-revathi/stockProfitCalculator/log"
)

type csvReader struct{}

func NewcsvReader() *csvReader {
	return &csvReader{}
}

//func loadCsvFile(filePath string) (map[int]map[string]string, error)
func (t *csvReader) getTransactions(ctx context.Context, filePath string) (map[int]map[string]string, error) {
	//TODO: check for file exists

	log.Debug(ctx, "trying to process the file - "+filePath)
	csvFile, fileErr := os.Open(filePath)
	if fileErr != nil {
		return nil, errors.New("error reading csv file " + fileErr.Error())

	}

	reader := csv.NewReader(csvFile)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("files is corrupted,couldn't read a file")
	}
	if len(records) < 1 {
		return nil, errors.New("no records in the file.fiile may be corrupted")
	}
	columnNames := getColumnNames(records[0])

	transactionData := make(map[int]map[string]string)

	for i := 1; i < len(records); i++ {

		transactionRow := make(map[string]string)
		for j := 0; j < len(records[i]); j++ {
			transactionRow[columnNames[j]] = strings.TrimSpace(records[i][j])
		}
		transactionData[i] = transactionRow
		// fmt.Printf("%v", transactionRow)
		// panic("err")
	}

	return transactionData, nil
}

func getColumnNames(record []string) []string {

	columnNames := make([]string, 0)

	for _, item := range record {
		columnNames = append(columnNames, strings.ToLower(strings.TrimSpace(item)))
	}

	return columnNames
}
