package main

import (
	"fmt"

	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Csv", func() {
	ctx := context.TODO()

	It("Invalid path to csv file or no file found", func() {
		var csv = csvReader{}

		_, err := csv.getTransactions(ctx, "aa.csv")
		fmt.Println("Invalid path to csv file or no file found", err)

		Expect(err).ToNot(Equal(nil))
	})

	It("corrupted file", func() {

		var csv = csvReader{}
		_, err := csv.getTransactions(ctx, "stock1.csv")
		Expect(err).ToNot(BeNil())
		fmt.Println(err)
	})
})

type TestReader struct{}

func (t *TestReader) getTransactions(filePath string) (m map[int]map[string]string, err error) {
	m = make(map[int]map[string]string)

	tRecord := make(map[string]string)

	tRecord["Market"] = "PLS"

	m[1] = tRecord

	return
	//return 3 transaction
}
