package calc_ginko_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/bmuschko/go-testing-frameworks/calc"
)

var _ = Describe("Calculator", func() {
	Describe("Add numbers", func() {
		Context("1 and 2", func() {
			It("should be 3", func() {
				Expect(Add(1, 2)).To(Equal(int64(3)))
			})
		})
	})

	Describe("Substract numbers", func() {
		Context("3 from 5", func() {
			It("should be 2", func() {
				Expect(Substract(5, 3)).To(Equal(int64(2)))
			})
		})
	})

	Describe("Multiply numbers", func() {
		Context("5 with 6", func() {
			It("should be 30", func() {
				Expect(Multiply(5, 6)).To(Equal(int64(30)))
			})
		})
	})

	Describe("Divide numbers", func() {
		Context("10 by 2", func() {
			It("should be 30", func() {
				Expect(Divide(10, 2)).To(Equal(float64(5)))
			})
		})
	})
})

