package main

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ProcessTransactions", func() {
	var _ = Context("ProcessTransactions", func() {
		It("ProcessTransactions", func() {
			transactions := make([]Transaction, 0)
			buytime, _ := time.Parse("1/2/2006", "10/12/2021")
			selltime, _ := time.Parse("1/2/2006", "11/20/2021")
			transactions = append(transactions, Transaction{"aaa", "buy", 705, 11.75, 60, buytime, "TRADE", 11.83})
			transactions = append(transactions, Transaction{"aaa", "sell", 765, 11.75, 60, selltime, "TRADE", 12.83})

			result := processTrans(transactions, Config{
				financialYear:       "2021",
				SkipCorporateAction: true,
				startFinancialMonth: "july",
				endFinancialMonth:   "jun",
			})
			Expect(len(result)).To(Equal(1))
			Expect(len(result["2021-2022"])).To(Equal(1))
			Expect(result["2021-2022"][0].PandL).To(Equal(float32(60)))
			

		})

		It("ProcessTransactions",func ()  {
			transactions := make([]Transaction,0)
			selltime,_ := time.Parse("1/2/2006","2021-02-09")
			transactions = append(transactions, Transaction{"NEXTDC" ,"SELL" ,3781.01 ,12.62 ,300,selltime, "TRADE", 12.603367})
			buytime,_ := time.Parse("1/2/2006", "2020-09-04" )
			transactions = append(transactions, Transaction{"NEXTDC","BUY", -2293.65, 11.33, 1,buytime, "TRADE", 11.354703})
			buytime,_ = time.Parse("1/2/2006", "2020-11-11" )
			transactions = append(transactions, Transaction{"NEXTDC", "BUY", -2874.75, 12.16, 236,buytime , "TRADE", 12.181144})
			buytime,_ = time.Parse("1/2/2006","2020-11-19")
			transactions = append(transactions, Transaction{"NEXTDC", "BUY", -1492.06, 12.09, 123, buytime,"TRADE", 12.130569})

			result := processTrans(transactions, Config{
				financialYear:       "2021",
				SkipCorporateAction: true,
				startFinancialMonth: "july",
				endFinancialMonth:   "jun",
			})
			Expect(len(result)).To(Equal(1))
			Expect(len(result["2021-2022"])).NotTo(Equal(0))
			Expect(result["2021-2022"][0].PandL).To(Equal(float32(60)))


		})

	})
})
