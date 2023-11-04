package cmd

type CalculatorFacade interface {
	Evaluate(string) (int, error)
	Validate(string) error
}

type CalculatorController struct {
	facade CalculatorFacade
}

func NewCalculatorController(calculatorRepo CalculatorFacade) *CalculatorController {
	return &CalculatorController{
		calculatorRepo,
	}
}

func (c *CalculatorController) Evaluate(expression string) (int, error) {
	result, err := c.facade.Evaluate(expression)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (c *CalculatorController) Validate(expression string) error {
	err := c.facade.Validate(expression)
	return err
}
