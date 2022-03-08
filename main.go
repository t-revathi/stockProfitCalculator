package main

import (
	"context"
	"flag"

	"github.com/t-revathi/stockProfitCalculator/log"
)

func main() {

	var (
		inputFilePath       string
		skipCorporateAction bool
		startFinancialMonth string
		endFinancialMonth   string
		financialYear       string
		logFileName         string
	)
	{
		flag.StringVar(&inputFilePath, "input-csv", "TradeHistory.csv", "csv input file path")
		flag.BoolVar(&skipCorporateAction, "skip-corp-action", true, "skip corporate action")
		flag.StringVar(&startFinancialMonth, "start-financial-month", "july", "month when the financial year starts")
		flag.StringVar(&endFinancialMonth, "end-financial-month", "jun", "month when the financial year ends")
		flag.StringVar(&financialYear, "financial-year", "2021", "caluclulation for financial year")
		flag.StringVar(&logFileName, "log-file-name", "deletelog", "Log file name")

	}

	flag.Parse()

	ctx := context.Background()

	config := Config{
		InputFilePath:       inputFilePath,
		SkipCorporateAction: skipCorporateAction,
		startFinancialMonth: startFinancialMonth,
		endFinancialMonth:   endFinancialMonth,
		financialYear:       financialYear,
	}

	logger := log.Newlogger(0, logFileName)
	logger.Info("Calling csv and calculate profits")
	ctx = log.AddLoggerToContext(ctx, logger)

	csvreader := NewcsvReader()
	c := NewCalculator(csvreader, config)
	c.calculateProfits(ctx)
	//calculateProfits(ctx, config,csvreader)

}
