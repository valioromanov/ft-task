package calculator

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//go:generate mockgen --source=presenter.go --destination mocks/presenter.go --package mocks

type Presenter struct {
	controller CalculatorController
}

func NewPresenter(controller CalculatorController) *Presenter {
	return &Presenter{
		controller,
	}
}

func (p *Presenter) Evaluate(ctx *gin.Context) {
	expression, err := p.extractBody(ctx)
	logrus.Info(fmt.Sprintf("new evaluate request recieved: %s", expression))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		logrus.Error(fmt.Sprintf("invalid body: %s", err.Error()))
		return
	}

	answer, err := p.controller.Evaluate(expression.Expression)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		logrus.Error(fmt.Sprintf("invalid body: %s", err.Error()))
		return
	}

	var resp ExpressionResponse
	resp, err = resp.ToExpressionResponse(answer)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		logrus.Error(fmt.Sprintf("internal server error: %s", err.Error()))
		return
	}

	logrus.Info(fmt.Sprintf("response: %d", resp.Result))
	ctx.JSON(http.StatusOK, resp)
}

func (p *Presenter) Validate(ctx *gin.Context) {
	expression, err := p.extractBody(ctx)
	logrus.Info(fmt.Sprintf("new validate request recieved: %s", expression))

	var validateResponse ValidateResponse
	if err != nil {
		validateResponse.Valid = false
		validateResponse.Reason = err.Error()
		logrus.Error(fmt.Sprintf("invalid body: %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, validateResponse)
		return
	}

	if err := p.controller.Validate(expression.Expression); err != nil {
		validateResponse.Valid = false
		validateResponse.Reason = err.Error()
		logrus.Info(fmt.Sprintf("Valid: %t with reason: %s", validateResponse.Valid, validateResponse.Reason))
		ctx.JSON(http.StatusBadRequest, validateResponse)
		return
	}
	validateResponse.Valid = true
	logrus.Info(fmt.Sprintf("Valid: %t", validateResponse.Valid))
	ctx.JSON(http.StatusOK, validateResponse)
}

func (p *Presenter) GetErrors(ctx *gin.Context) {
	logrus.Info("new getErrors request recieved")
	invalids := p.controller.GetErrors()

	invalidsResponse := invalids.ToGetErrorsResponse()

	ctx.JSON(http.StatusOK, invalidsResponse)
}

func (p *Presenter) extractBody(ctx *gin.Context) (ExpressionRequest, error) {
	var requestBody ExpressionRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		return ExpressionRequest{}, fmt.Errorf("invalid request body: %w", err)
	}

	return requestBody, nil
}
