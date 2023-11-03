package calculator_test

import (
	"fmt"
	"ft-calculator/cmd/calcurator-api/internal/calculator"
	"ft-calculator/cmd/calcurator-api/internal/calculator/mocks"
	"ft-calculator/helper/mockutil"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator Controller", mockutil.Mockable(func(helper *mockutil.Helper) {

	var (
		mockRepository *mocks.MockCalculatorRepository
		controller     calculator.CalculatorController
		invalids       calculator.InvalidExpression
	)

	BeforeEach(func() {
		mockRepository = mocks.NewMockCalculatorRepository(helper.Controller())
		invalids = make(calculator.InvalidExpression)
		controller = calculator.NewCalculatorController(mockRepository, invalids)
	})

	Context("Evaluate", func() {
		var exptectedResponse int
		When("evrything is okay", func() {

			BeforeEach(func() {
				exptectedResponse = 3
				gomock.InOrder(
					mockRepository.EXPECT().Validate(gomock.Any()).Return([]int{1, 2}, []string{"+"}, nil),
					mockRepository.EXPECT().Calculate(gomock.Any(), gomock.Any()).Return(exptectedResponse),
				)
			})

			It("should return the result", func() {
				actualResult, err := controller.Evaluate(gomock.Any().String())
				Expect(err).ToNot(HaveOccurred())
				Expect(actualResult).To(Equal(exptectedResponse))
			})
		})

		When("expression is not valid", func() {
			expErr := "error"
			BeforeEach(func() {
				gomock.InOrder(
					mockRepository.EXPECT().Validate(gomock.Any()).Return(nil, nil, fmt.Errorf(expErr)),
				)
				errors := controller.GetErrors()
				Expect(errors).To(BeEmpty())
			})

			It("should add in invalid in-memory map and return an error", func() {
				_, err := controller.Evaluate(gomock.Any().String())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(expErr))
				errors := controller.GetErrors()
				Expect(errors).ToNot(BeEmpty())
			})
		})

	})

	Context("Validate", func() {
		When("expression is valid", func() {
			BeforeEach(func() {
				gomock.InOrder(
					mockRepository.EXPECT().Validate(gomock.Any()).Return([]int{1, 2}, []string{"+"}, nil),
				)
			})

			It("should return no error", func() {
				err := controller.Validate(gomock.Any().String())
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("expression is not valid", func() {
			BeforeEach(func() {
				gomock.InOrder(
					mockRepository.EXPECT().Validate(gomock.Any()).Return(nil, nil, fmt.Errorf("some-error")),
				)
				errors := controller.GetErrors()
				Expect(errors).To(BeEmpty())
			})

			It("should add in invalid in-memory map return an error", func() {
				err := controller.Validate(gomock.Any().String())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("some-error"))
				errors := controller.GetErrors()
				Expect(errors).ToNot(BeEmpty())
			})
		})
	})
}))
