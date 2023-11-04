/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	calculatorRepo "ft-calculator/pkg/calculator"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runValidate(expressionStatement)
	},
}

func runValidate(exression string) {
	calculatorRepo := calculatorRepo.Calculator{}
	controller := NewCalculatorController(&calculatorRepo)
	err := controller.Validate(exression)

	if err != nil {
		fmt.Printf("The expression '%s' is not valid!\n", exression)
		fmt.Printf("Reason: %s \n", err.Error())
		return
	}

	fmt.Printf("The experession '%s' is valid! \n", exression)
}

func init() {
	rootCmd.AddCommand(validateCmd)
	validateCmd.Flags().StringVarP(&expressionStatement, "expression", "e", "", "expression to be evaluated")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
