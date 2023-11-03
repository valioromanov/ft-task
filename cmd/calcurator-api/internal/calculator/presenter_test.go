package calculator_test

import (
	"ft-calculator/cmd/calcurator-api/internal/calculator"
	"ft-calculator/cmd/calcurator-api/internal/calculator/mocks"
	"ft-calculator/helper/mockutil"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Presenter", mockutil.Mockable(func(helper *mockutil.Helper) {

	var (
		recorder    *httptest.ResponseRecorder
		controller  mocks.MockCalculatorController
		mockContext *gin.Context
		presenter   *calculator.CalculatorRepository
	)

	BeforeEach(func() {
		recorder = httptest.NewRecorder()
		mockContext, _ = gin.CreateTestContext(recorder)
		controllerPointer := mocks.NewMockCalculatorController(helper.Controller())
		presenter = calculator.NewPresenter(controllerPointer)
	})

}))
