/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	calculatorRepo "ft-calculator/pkg/calculator"

	"github.com/spf13/cobra"
)

var expressionStatement string

// evaluateCmd represents the evaluate command
var evaluateCmd = &cobra.Command{
	Use:   "evaluate",
	Short: "Returns the result of the expression",
	Long: `Returns the result of the expression:
		For example if you send What is 3 plus 1
		it returns 3`,
	Run: func(cmd *cobra.Command, args []string) {
		runEvaluate(expressionStatement)
	},
}

func runEvaluate(exression string) {
	calculatorRepo := calculatorRepo.Calculator{}
	controller := NewCalculatorController(&calculatorRepo)
	result, _ := controller.Evaluate(exression)
	fmt.Println(result)
}

func init() {
	rootCmd.AddCommand(evaluateCmd)
	evaluateCmd.Flags().StringVarP(&expressionStatement, "expression", "e", "", "expression to be evaluated")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// evaluateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// evaluateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
