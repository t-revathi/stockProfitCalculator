package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func loadCsvFile(filePath string) map[int]map[string]string {
	//TODO: check for file exists

	//mydir, _ := os.Getwd()
	//filePath = filePath

	fmt.Println("trying to process the file - " + filePath)

	csvFile, fileErr := os.Open(filePath)
	if fileErr != nil {
		panic("error reading csv file " + fileErr.Error())
	}

	reader := csv.NewReader(csvFile)

	records, _ := reader.ReadAll()

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

	return transactionData
}

func getColumnNames(record []string) []string {

	columnNames := make([]string, 0)

	for _, item := range record {
		columnNames = append(columnNames, strings.ToLower(strings.TrimSpace(item)))
	}

	return columnNames
}
