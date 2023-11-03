package calculator_test

import (
	"ft-calculator/cmd/calcurator-api/internal/calculator"
	"ft-calculator/cmd/calcurator-api/internal/calculator/mocks"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator Controller", func(helper *mockutil.Helper) {

	var (
		mockRepository *mocks.MockCalculatorRepository
		controller     *calculator.CalculatorController
	)
})
