package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStockProfitCalculator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StockProfitCalculator Suite")
}
