package calculator

import "fmt"

type CalculatorController interface {
	Evaluate(string) int
}

type CalculatorRepository interface {
	Calculate([]int, []string) int
	Validate(string) ([]int, []string, error)
}

type calculatorController struct {
	calculatorRepo CalculatorRepository
}

func NewCalculatorController(calculatorRepo CalculatorRepository) CalculatorController {
	return &calculatorController{
		calculatorRepo,
	}
}

func (c *calculatorController) Evaluate(expression string) int {
	nums, ops, err := c.calculatorRepo.Validate(expression)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c.calculatorRepo.Calculate(nums, ops)
}
