package main

import (
	"context"
	"time"
)

type Config struct {
	InputFilePath       string
	SkipCorporateAction bool
	financialYear       string
	startFinancialMonth string
	endFinancialMonth   string
}

type Transaction struct {
	Market    string
	Direction string
	Cost      float32
	Price     float32
	Quantity  int
	Date      time.Time
	Activity  string
	UnitPrice float32
}
type Income struct {
	Date          time.Time
	Market        string
	Quantity      int
	PandL         float32
	SellUnitPrice float32
}

type TransactionData interface {
	getTransactions(ctx context.Context,filepath string) (map[int]map[string]string, error)
}

