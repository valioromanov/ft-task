package calculator

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

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

	logrus.Info(fmt.Sprintf("response: %s", resp.ToString()))
	ctx.Data(http.StatusOK, "application/json", []byte(resp.ToString()))
}

func (p *Presenter) extractBody(ctx *gin.Context) (ExpressionRequest, error) {
	var requestBody ExpressionRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		return ExpressionRequest{}, fmt.Errorf("invalid request body: %w", err)
	}

	return requestBody, nil
}
