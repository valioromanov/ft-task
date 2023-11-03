package calculatorRepo

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator Repository", func() {

	type CalculatorRepository interface {
		Calculate([]int, []string) int
		Validate(string) ([]int, []string, error)
	}

	var repository CalculatorRepository

	BeforeEach(func() {
		repository = &Calculator{}
	})

	Describe("Evaluate", func() {
		var exptectedResult int
		nums := make([]int, 2)
		ops := make([]string, 1)

		BeforeEach(func() {
			nums = []int{6, 2}
		})

		When("should sum two numbers", func() {
			BeforeEach(func() {
				exptectedResult = 8
				ops = []string{"+"}
			})

			It("should return the expected result", func() {
				actualResult := repository.Calculate(nums, ops)
				Expect(actualResult).To(Equal(exptectedResult))
			})
		})

		When("should diff two numbers", func() {
			BeforeEach(func() {
				ops[0] = "-"
				exptectedResult = 4
			})

			It("should return the expected result", func() {
				actualResult := repository.Calculate(nums, ops)
				Expect(actualResult).To(Equal(exptectedResult))
			})
		})

		When("should multiply two numbers", func() {
			BeforeEach(func() {
				ops[0] = "*"
				exptectedResult = 12
			})

			It("should return the expected result", func() {
				actualResult := repository.Calculate(nums, ops)
				Expect(actualResult).To(Equal(exptectedResult))
			})
		})

		When("should divide two numbers", func() {
			BeforeEach(func() {
				ops[0] = "/"
				exptectedResult = 3
			})

			It("should return the expected result", func() {
				actualResult := repository.Calculate(nums, ops)
				Expect(actualResult).To(Equal(exptectedResult))
			})
		})

		When("when only one number recieved in parameters", func() {
			var nums []int
			var ops []string
			BeforeEach(func() {
				nums = []int{6}
				ops = []string{}
			})

			It("should return the number", func() {
				actualResult := repository.Calculate(nums, ops)
				Expect(actualResult).To(Equal(nums[0]))
			})
		})
	})
})
