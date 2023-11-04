package facade

import (
	"github.com/sirupsen/logrus"
)

//go:generate mockgen --source=facade.go --destination mocks/facade.go --package mocks

type CalculatorRepository interface {
	Calculate([]int, []string) int
	Validate(string) ([]int, []string, error)
}

type Facade struct {
	calculatorRepo CalculatorRepository
	invalids       InvalidExpression
}

func NewCalulatorFacade(calculatorRepo CalculatorRepository, invalids InvalidExpression) *Facade {
	return &Facade{
		calculatorRepo,
		invalids,
	}
}

func (f *Facade) Evaluate(expression string) (int, error) {
	nums, ops, err := f.calculatorRepo.Validate(expression)
	if err != nil {
		f.addInvalid(expression, "\\evaluate", err.Error())
		return 0, err
	}
	return f.calculatorRepo.Calculate(nums, ops), nil
}

func (f *Facade) Validate(expression string) error {
	_, _, err := f.calculatorRepo.Validate(expression)
	if err != nil {
		f.addInvalid(expression, "\\validate", err.Error())
	}
	return err
}

func (f *Facade) GetErrors() InvalidExpression {
	return f.invalids
}

func (f *Facade) addInvalid(expression, endpoint, err string) {
	invalidKey := InvalidKey{
		Expression: expression,
		Endpoint:   endpoint,
	}
	data, ok := f.invalids[invalidKey]
	if ok {
		logrus.Info("in add invalid: ", data.Frequency)
		data.Frequency = data.Frequency + 1
		invalidData := InvalidData{
			Type:      err,
			Frequency: data.Frequency,
		}
		f.invalids[invalidKey] = invalidData
		return
	}

	f.invalids[invalidKey] = InvalidData{
		Type:      err,
		Frequency: 1,
	}
}
