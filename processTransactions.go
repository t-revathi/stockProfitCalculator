package main

import (
	"context"
	"fmt"
)

func calculateProfits(ctx context.Context, config Config) {

	transactionData := loadCsvFile(config.InputFilePath)
	fmt.Printf("%+v \n", transactionData)

}
