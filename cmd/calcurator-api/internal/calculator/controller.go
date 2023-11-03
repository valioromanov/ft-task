package calculator

import (
	"github.com/sirupsen/logrus"
)

type CalculatorRepository interface {
	Calculate([]int, []string) int
	Validate(string) ([]int, []string, error)
}

type calculatorController struct {
	calculatorRepo CalculatorRepository
	invalids       InvalidExpression
}

func NewCalculatorController(calculatorRepo CalculatorRepository, invalids InvalidExpression) CalculatorController {
	return &calculatorController{
		calculatorRepo,
		invalids,
	}
}

func (c *calculatorController) Evaluate(expression string) (int, error) {
	nums, ops, err := c.calculatorRepo.Validate(expression)
	if err != nil {
		c.addInvalid(expression, "\\evaluate", err.Error())
		return 0, err
	}
	return c.calculatorRepo.Calculate(nums, ops), nil
}

func (c *calculatorController) Validate(expression string) error {
	_, _, err := c.calculatorRepo.Validate(expression)
	if err != nil {
		c.addInvalid(expression, "\\validate", err.Error())
	}
	return err
}

func (c *calculatorController) GetErrors() InvalidExpression {
	return c.invalids
}

func (c *calculatorController) addInvalid(expression, endpoint, err string) {
	invalidKey := InvalidKey{
		Expression: expression,
		Endpoint:   endpoint,
	}
	data, ok := c.invalids[invalidKey]
	if ok {
		logrus.Info("in add invalid: ", data.Frequency)
		data.Frequency = data.Frequency + 1
		invalidData := InvalidData{
			Type:      err,
			Frequency: data.Frequency,
		}
		c.invalids[invalidKey] = invalidData
		return
	}

	c.invalids[invalidKey] = InvalidData{
		Type:      err,
		Frequency: 1,
	}
}
