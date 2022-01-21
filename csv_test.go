package main

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Csv", func() {
	It("Invalid path to csv file or no file found", func() {
		_, err := loadCsvFile("aa.csv")
		fmt.Println("Invalid path to csv file or no file found", err)

		Expect(err).ToNot(Equal(nil))
	})

	It("corrupted file or invalid file", func() {
		_, err := loadCsvFile("stock1.txt")
		Expect(err).ToNot(BeNil())
		fmt.Println(err)
	})
})
