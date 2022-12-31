package main

import (
	"context"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/t-revathi/stockProfitCalculator/log"
)

type Calculator struct {
	t      TransactionData
	config Config
}

func NewCalculator(t TransactionData, config Config) *Calculator {
	return &Calculator{
		t:      t,
		config: config,
	}
}

func (c *Calculator) calculateProfits(ctx context.Context) {

	config := c.config
	transactionData, err := c.t.getTransactions(ctx, config.InputFilePath)
	if err != nil {
		log.Error(ctx, "Error processing file:", err)
		return
	}
	//fmt.Printf("%+v \n", transactionData)
	transactions := mapToStruct(ctx, transactionData)
	income := processTrans(transactions, config)
	for idx, val := range income {

		//if (strconv.Itoa(val.Date.Year)) == config.financialYear {
		fmt.Printf("%v %+v \n\n", idx, val)
		var totalPandL float32 = 0.0
		for _, finalresult := range val {
			totalPandL += finalresult.PandL
		}
		fmt.Printf("********\n %v = %v \n\n", idx, totalPandL)
		//}

	}

}
func processTrans(transactions []Transaction, config Config) map[string][]Income {
	buyShares := getbuyShares(transactions, config)
	//fmt.Println(buyShares)
	sellShares := getsellShares(transactions, config)
	return calculatePandL(buyShares, sellShares, config)
}
func calculatePandL(buyshares []Transaction, sellShares []Transaction, config Config) map[string][]Income {
	//income := make(map[string][]float32)
	income := make(map[string][]Income)
	//income := make(map[int][]Income)

	for idx := range sellShares {
		pq := 0
		var pl float32 = 0.0
		var inc Income
		var currentSellRecordYear string
		currentSellRecord := sellShares[idx]
		if currentSellRecord.Date.Month() < 7 {
			currentSellRecordYear = strconv.Itoa((currentSellRecord.Date.Year() - 1)) + "-" + strconv.Itoa(currentSellRecord.Date.Year())
		} else {
			currentSellRecordYear = strconv.Itoa(currentSellRecord.Date.Year()) + "-" + strconv.Itoa((currentSellRecord.Date.Year() + 1))
		}

		inc.Date = currentSellRecord.Date
		inc.Market = currentSellRecord.Market
		inc.Quantity = currentSellRecord.Quantity
		inc.SellUnitPrice = currentSellRecord.UnitPrice
		fmt.Printf("Sell: %v \n", currentSellRecord)
		for currentSellRecord.Quantity > 0 {
			buyt := getearlierbuyShare(buyshares, currentSellRecord.Market)
			if buyt.Quantity >= currentSellRecord.Quantity {
				pq = currentSellRecord.Quantity
			} else {
				pq = buyt.Quantity
			}
			//if currentSellRecord.Market == "Pilbara Minerals Limited" {
			fmt.Printf("%v \n\n", buyt)
			//fmt.Printf("%d\n%d\n%d\n\n", buyt.Quantity, sellShares[idx].Quantity, currentSellRecord.Quantity)
			//}
			pl += (currentSellRecord.UnitPrice - buyt.UnitPrice) * float32(pq)

			buyt.Quantity, sellShares[idx].Quantity, currentSellRecord.Quantity = buyt.Quantity-pq, currentSellRecord.Quantity-pq, currentSellRecord.Quantity-pq
			/*if _, ok := income[currentSellRecord.Market]; !ok {
				income[currentSellRecord.Market] = make([]Income, 0)
			}*/

			if _, ok := income[currentSellRecordYear]; !ok {
				income[currentSellRecordYear] = make([]Income, 0)
			}

			//income = append(income,income[sellShares[idx].Market])
		}

		inc.PandL = pl

		//income[currentSellRecord.Market] = append(income[currentSellRecord.Market], inc)
		income[currentSellRecordYear] = append(income[currentSellRecordYear], inc)

	}
	//fmt.Printf("%v+ \n", income)
	return income
}

func getearlierbuyShare(buyshares []Transaction, market string) *Transaction {
	mindate := time.Now()
	//var earliershare Transaction
	earlierShareIdx := 0
	for idx := range buyshares {
		bshares := buyshares[idx]
		if bshares.Market == market {
			if bshares.Date.Before(mindate) && bshares.Quantity > 0 {
				mindate = bshares.Date
				earlierShareIdx = idx
			}
		}
	}
	return &buyshares[earlierShareIdx]
}

func getsellShares(transaction []Transaction, config Config) []Transaction {
	selltransaction := make([]Transaction, 0)
	for _, t := range transaction {
		if config.SkipCorporateAction {
			if strings.ToLower(t.Activity) != "trade" {
				continue
			}
		}
		if strings.ToLower(t.Direction) == "sell" {
			selltransaction = append(selltransaction, t)

		}
	}
	sort.Slice(selltransaction, func(i, j int) bool {
		return selltransaction[i].Date.Before(selltransaction[j].Date)
	})
	return selltransaction
}
func getbuyShares(transactions []Transaction, config Config) []Transaction {

	buytransaction := make([]Transaction, 0)

	for _, t := range transactions {
		if config.SkipCorporateAction {
			if strings.ToLower(t.Activity) != "trade" {
				continue
			}
		}
		if strings.ToLower(t.Direction) == "buy" {
			buytransaction = append(buytransaction, t)

		}

	}
	return buytransaction
}

func Abs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}

func mapToStruct(ctx context.Context, transactionData map[int]map[string]string) []Transaction {

	transactions := make([]Transaction, 0)

	for _, t := range transactionData {

		transaction := Transaction{}

		//fmt.Println(t["activity"])
		transaction.Date = getDate(ctx, t["date"])
		//fmt.Println(transaction.Date)
		transaction.Market = t["market"]
		transaction.Cost = getFloat(ctx, t["cost/proceeds"])
		transaction.Direction = t["direction"]
		transaction.Price = getFloat(ctx, t["price"])
		transaction.Activity = t["activity"]

		transaction.Quantity = (getInt(ctx, t["quantity"]))
		unitPrice := float64((transaction.Cost / float32(transaction.Quantity)))
		transaction.UnitPrice = float32(math.Abs(unitPrice))

		transactions = append(transactions, transaction)
	}

	return transactions
}
func getInt(ctx context.Context, data string) int {
	value, err := strconv.ParseInt(data, 10, 0)

	if err != nil {
		//panic(err)
		log.Fatal(ctx, "", err)
	}
	//return int(value)
	return int(math.Abs(float64(value)))
}

func getFloat(ctx context.Context, data string) float32 {
	value, err := strconv.ParseFloat(data, 32)

	if err != nil {
		log.Fatal(ctx, "", err)
	}
	return float32(value)
}

func getDate(ctx context.Context, data string) time.Time {
	str := strings.Split(data, "/")
	//convert date from dd/mm/yyyy to mm/dd/yyyy
	data = str[1] + "/" + str[0] + "/" + str[2]

	t, err := time.Parse("1/2/2006", data)

	if err != nil {
		log.Fatal(ctx, "", err)
	}

	return t
}
