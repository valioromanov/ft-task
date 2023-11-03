package calculator_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ft-calculator/cmd/calcurator-api/internal/calculator"
	"ft-calculator/cmd/calcurator-api/internal/calculator/mocks"
	"ft-calculator/helper/mockutil"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Presenter", mockutil.Mockable(func(helper *mockutil.Helper) {

	var (
		recorder    *httptest.ResponseRecorder
		controller  *mocks.MockCalculatorController
		mockContext *gin.Context
		presenter   *calculator.Presenter
	)

	BeforeEach(func() {
		recorder = httptest.NewRecorder()
		mockContext, _ = gin.CreateTestContext(recorder)
		controller = mocks.NewMockCalculatorController(helper.Controller())
		presenter = calculator.NewPresenter(controller)
	})

	Describe("Evaluate", func() {
		expression := "What is 3 plus 1"
		When("request is valid", func() {
			BeforeEach(func() {
				reqBody := map[string]interface{}{
					"expression": expression,
				}
				mockContext.Request = httptest.NewRequest("POST", "http://abc.com", createBody(reqBody))
				controller.EXPECT().Evaluate(expression).Return(4, nil)
			})

			It("returns StatusOK and the result", func() {
				presenter.Evaluate(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusOK))
			})
		})

		When("request body is empty", func() {
			BeforeEach(func() {
				mockContext.Request = httptest.NewRequest("POST", "http://abc.com", nil)
			})

			It("returns StatusBadRequest", func() {
				presenter.Evaluate(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusBadRequest))
			})
		})

		When("request body is invalid", func() {
			BeforeEach(func() {
				invalidExpression := "What is 3 plus plus 1"
				reqBody := map[string]interface{}{
					"expression": invalidExpression,
				}
				mockContext.Request = httptest.NewRequest("POST", "http://abc.com", createBody(reqBody))
				controller.EXPECT().Evaluate(invalidExpression).Return(0, fmt.Errorf("some-error"))
			})

			It("returns StatusBadRequest", func() {
				presenter.Evaluate(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusBadRequest))
			})
		})
	})

}))

func createBody(fields map[string]interface{}) *bytes.Buffer {
	b, err := json.Marshal(fields)
	Expect(err).ToNot(HaveOccurred())

	return bytes.NewBufferString(string(b))
}
