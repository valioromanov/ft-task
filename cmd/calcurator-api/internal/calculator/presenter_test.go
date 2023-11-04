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
	"github.com/golang/mock/gomock"
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
				mockContext.Request, _ = http.NewRequest("POST", gomock.Any().String(), createBody(reqBody))
				controller.EXPECT().Evaluate(expression).Return(4, nil)
			})

			It("returns StatusOK and the result", func() {
				evaluateRes := calculator.ExpressionResponse{}
				presenter.Evaluate(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusOK))
				Expect(json.Unmarshal(recorder.Body.Bytes(), &evaluateRes)).To(Succeed())
				Expect(evaluateRes.Result).To(Equal(4))
			})
		})

		When("request body is empty", func() {
			BeforeEach(func() {
				mockContext.Request, _ = http.NewRequest("POST", gomock.Any().String(), nil)
			})

			It("returns StatusBadRequest and error message", func() {
				errMsg := ""
				presenter.Evaluate(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusBadRequest))
				Expect(json.Unmarshal(recorder.Body.Bytes(), &errMsg)).To(Succeed())
				Expect(errMsg).To(ContainSubstring("invalid request"))
			})
		})

		When("request body is invalid", func() {
			BeforeEach(func() {
				invalidExpression := "What is 3 plus plus 1"
				reqBody := map[string]interface{}{
					"expression": invalidExpression,
				}
				mockContext.Request, _ = http.NewRequest("POST", gomock.Any().String(), createBody(reqBody))
				controller.EXPECT().Evaluate(invalidExpression).Return(0, fmt.Errorf("some-error"))
			})

			It("returns StatusBadRequest and an error message", func() {
				errMsg := ""
				presenter.Evaluate(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusBadRequest))
				Expect(json.Unmarshal(recorder.Body.Bytes(), &errMsg)).To(Succeed())
				Expect(errMsg).To(ContainSubstring("some-error"))
			})
		})
	})

	Describe("Validate", func() {
		When("request is valid", func() {
			expression := "What is 3 plus 1"
			BeforeEach(func() {
				reqBody := map[string]interface{}{
					"expression": expression,
				}
				mockContext.Request, _ = http.NewRequest("POST", gomock.Any().String(), createBody(reqBody))
				controller.EXPECT().Validate(expression).Return(nil)
			})

			It("it should return StatusOK", func() {
				presenter.Validate(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusOK))
			})
		})

		When("request body is empty", func() {
			BeforeEach(func() {
				mockContext.Request, _ = http.NewRequest("POST", gomock.Any().String(), nil)
			})

			It("it should return StatusBadRequest", func() {
				validateResp := calculator.ValidateResponse{}
				presenter.Validate(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusBadRequest))
				Expect(json.Unmarshal(recorder.Body.Bytes(), &validateResp)).To(Succeed())
				Expect(validateResp.Valid).To(BeFalse())
				Expect(validateResp.Reason).To(ContainSubstring("invalid request"))
			})
		})

		When("request is invalid", func() {
			BeforeEach(func() {
				invalidExpression := "What is 3 plus plus 1"
				reqBody := map[string]interface{}{
					"expression": invalidExpression,
				}

				mockContext.Request, _ = http.NewRequest("POST", gomock.Any().String(), createBody(reqBody))
				controller.EXPECT().Validate(invalidExpression).Return(fmt.Errorf("some-error"))
			})

			It("it should return StatusBadRequest and return proper object", func() {
				validateResp := calculator.ValidateResponse{}
				presenter.Validate(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusBadRequest))
				Expect(json.Unmarshal(recorder.Body.Bytes(), &validateResp)).To(Succeed())
				Expect(validateResp.Valid).To(BeFalse())
				Expect(validateResp.Reason).To(Equal("some-error"))
			})
		})
	})

	Describe("GetErrors", func() {

		When("request is valid", func() {
			var (
				expression string
				endpoint   string
				frequency  int
				typeErr    string
			)
			expression = "What is 3 plus 2"
			endpoint = "\\evaluate"
			frequency = 1
			typeErr = "some-type"
			exptectedControllerResponse := calculator.InvalidExpression{
				calculator.InvalidKey{endpoint, expression}: calculator.InvalidData{frequency, typeErr},
			}
			BeforeEach(func() {
				mockContext.Request, _ = http.NewRequest("GET", gomock.Any().String(), nil)
				controller.EXPECT().GetErrors().Return(exptectedControllerResponse)
			})

			It("should return StatusOK", func() {
				errResp := []calculator.GetErrorResponse{}
				presenter.GetErrors(mockContext)
				Expect(mockContext.Writer.Status()).To(Equal(http.StatusOK))
				Expect(json.Unmarshal(recorder.Body.Bytes(), &errResp)).To(Succeed())
				Expect(len(errResp)).To(Equal(1))
				Expect(errResp[0].Endpoint).To(Equal(endpoint))
				Expect(errResp[0].Expression).To(Equal(expression))
				Expect(errResp[0].Frequency).To(Equal(frequency))
				Expect(errResp[0].Type).To(Equal(typeErr))
			})
		})
	})
}))

func createBody(fields map[string]interface{}) *bytes.Buffer {
	b, err := json.Marshal(fields)
	Expect(err).ToNot(HaveOccurred())

	return bytes.NewBufferString(string(b))
}
