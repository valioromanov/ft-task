package calculator

type Presenter struct {
	controller CalculatorController
}

func NewPresenter(controller CalculatorController) *Presenter {
	return &Presenter{
		controller,
	}
}

func (p *Presenter) Evaluate(expression string) int {
	return p.controller.Evaluate(expression)
}
