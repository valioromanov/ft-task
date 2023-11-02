package main

import (
	"fmt"
	"ft-calculator/cmd/calcurator-api/internal/calculator"
	calculatorRepo "ft-calculator/pkg/calculator"
)

func main() {

	calculatorRepo := calculatorRepo.Calculator{}
	calculatorController := calculator.NewCalculatorController(&calculatorRepo)
	presenter := calculator.NewPresenter(calculatorController)
	exp := "3 plus 2 plus"
	fmt.Print(presenter.Evaluate(exp))
}
