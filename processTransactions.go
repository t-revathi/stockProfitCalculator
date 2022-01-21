package main

import (
	"context"
	"fmt"
)

func calculateProfits(ctx context.Context, config Config) {

	transactionData, err := loadCsvFile(config.InputFilePath)
	if err != nil {
		return
	}
	fmt.Printf("%+v \n", transactionData)

}
