package calculatorRepo

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	operatorsToSymbols = map[string]string{
		"plus":          "+",
		"minus":         "-",
		"multiplied by": "*",
		"divided by":    "/",
	}

	operatorsToFunctions = map[string]func(a, b int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
)

type Calculator struct{}

func (c *Calculator) Calculate(numbers []int, operators []string) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	result := 0
	for len(numbers) > 1 {
		result = operatorsToFunctions[operators[0]](numbers[0], numbers[1])
		numbers[1] = result
		numbers = numbers[1:]
		operators = operators[1:]
	}
	return result
}

func (c *Calculator) Validate(exp string) ([]int, []string, error) {
	splitExp := strings.Split(exp, "What is")
	if len(splitExp) <= 1 {
		return nil, nil, fmt.Errorf("expressions with invalid syntax")
	}

	mathEquation := replaceOperators(strings.TrimSpace(splitExp[1]))
	nums, ops := getNumbersAndOps(strings.Split(mathEquation, " "))
	if len(nums) == 0 {
		return nil, nil, fmt.Errorf("non-math questions")
	}

	if !checkUsupportedOpperators(ops) && len(ops) > 0 {
		return nil, nil, fmt.Errorf("unsupported operations")
	}

	if len(nums) <= len(ops) {
		return nil, nil, fmt.Errorf("expressions with invalid syntax")
	}

	return nums, ops, nil
}

func getNumbersAndOps(exp []string) (numbers []int, ops []string) {
	for _, val := range exp {
		if numb, err := strconv.Atoi(val); err == nil {
			numbers = append(numbers, numb)
		} else {
			ops = append(ops, val)
		}
	}
	return
}

func replaceOperators(s string) string {
	for key, val := range operatorsToSymbols {
		s = strings.ReplaceAll(s, key, val)
	}
	return s
}

func checkUsupportedOpperators(ops []string) bool {
	for _, operator := range ops {
		if _, ok := operatorsToFunctions[operator]; !ok {
			return false
		}
	}

	return true
}
