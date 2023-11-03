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

	Describe("Validate", func() {
		var (
			expression   string
			expectedNums []int
			exptectedOps []string
		)

		When("everything is okay", func() {
			BeforeEach(func() {
				expression = "What is 3 plus 5"
				expectedNums = []int{3, 5}
				exptectedOps = []string{"+"}
			})

			It("should return nil error and extracted numbers and operators", func() {
				nums, ops, err := repository.Validate(expression)
				Expect(err).ToNot(HaveOccurred())
				Expect(nums).To(Equal(expectedNums))
				Expect(ops).To(Equal(exptectedOps))
			})
		})

		When("not exptession does not start with 'What is'", func() {
			BeforeEach(func() {
				expression = "I am 3 plus 5"
			})

			It("should return an error", func() {
				_, _, err := repository.Validate(expression)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("expressions with invalid syntax"))
			})
		})

		When("non-math expression is recieved", func() {
			BeforeEach(func() {
				expression = "What is the capital of Bulgaria"
			})

			It("should return an error", func() {
				_, _, err := repository.Validate(expression)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("non-math questions"))
			})
		})

		When("when ivalid syntax is recieved", func() {
			BeforeEach(func() {
				expression = "What is 3 plus plus 2"
			})

			It("should return an error", func() {
				_, _, err := repository.Validate(expression)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("expressions with invalid syntax"))
			})
		})

		When("when there is an unsupported opperation", func() {
			BeforeEach(func() {
				expression = "What is 3 cubed"
			})

			It("should return an error", func() {
				_, _, err := repository.Validate(expression)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("unsupported operations"))
			})
		})
	})
})
