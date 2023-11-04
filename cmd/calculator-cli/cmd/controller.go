package cmd

type CalculatorRepository interface {
	Calculate([]int, []string) int
	Validate(string) ([]int, []string, error)
}

type CalculatorController struct {
	calculatorRepo CalculatorRepository
}

func NewCalculatorController(calculatorRepo CalculatorRepository) *CalculatorController {
	return &CalculatorController{
		calculatorRepo,
	}
}

func (c *CalculatorController) Evaluate(expression string) (int, error) {
	nums, ops, err := c.calculatorRepo.Validate(expression)
	if err != nil {
		return 0, err
	}
	return c.calculatorRepo.Calculate(nums, ops), nil
}

func (c *CalculatorController) Validate(expression string) error {
	_, _, err := c.calculatorRepo.Validate(expression)
	return err
}
