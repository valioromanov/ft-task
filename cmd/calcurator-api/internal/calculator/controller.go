package calculator

type CalculatorController interface {
	Evaluate(string) (int, error)
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

func (c *calculatorController) Evaluate(expression string) (int, error) {
	nums, ops, err := c.calculatorRepo.Validate(expression)
	//TODO add inserting in error in-memory
	if err != nil {
		return 0, err
	}
	return c.calculatorRepo.Calculate(nums, ops), nil
}
