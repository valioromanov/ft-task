package facade_test

import (
	"fmt"
	"ft-calculator/helper/mockutil"
	"ft-calculator/pkg/facade"
	"ft-calculator/pkg/facade/mocks"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator Controller", mockutil.Mockable(func(helper *mockutil.Helper) {

	var (
		mockRepository *mocks.MockCalculatorRepository
		mockFacade     facade.Facade
		invalids       facade.InvalidExpression
	)

	BeforeEach(func() {
		mockRepository = mocks.NewMockCalculatorRepository(helper.Controller())
		invalids = make(facade.InvalidExpression)
		mockFacade = *facade.NewCalulatorFacade(mockRepository, invalids)
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
				actualResult, err := mockFacade.Evaluate(gomock.Any().String())
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
				errors := mockFacade.GetErrors()
				Expect(errors).To(BeEmpty())
			})

			It("should add in invalid in-memory map and return an error", func() {
				_, err := mockFacade.Evaluate(gomock.Any().String())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(expErr))
				errors := mockFacade.GetErrors()
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
				err := mockFacade.Validate(gomock.Any().String())
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("expression is not valid", func() {
			BeforeEach(func() {
				gomock.InOrder(
					mockRepository.EXPECT().Validate(gomock.Any()).Return(nil, nil, fmt.Errorf("some-error")),
				)
				errors := mockFacade.GetErrors()
				Expect(errors).To(BeEmpty())
			})

			It("should add in invalid in-memory map return an error", func() {
				err := mockFacade.Validate(gomock.Any().String())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("some-error"))
				errors := mockFacade.GetErrors()
				Expect(errors).ToNot(BeEmpty())
			})
		})
	})
}))
