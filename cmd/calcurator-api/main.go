package main

import (
	"fmt"
	"ft-calculator/cmd/calcurator-api/internal/calculator"
	"ft-calculator/pkg/app"
	calculatorRepo "ft-calculator/pkg/calculator"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func handleError(err error, m string) {
	if err != nil {
		app.Crash(fmt.Errorf("%s: %w", m, err))
	}
}

func main() {

	calculatorRepo := calculatorRepo.Calculator{}

	invalids := make(calculator.InvalidExpression)
	calculatorController := calculator.NewCalculatorController(&calculatorRepo, invalids)
	presenter := calculator.NewPresenter(calculatorController)

	handler := gin.New()
	handler.POST("/evaluate", presenter.Evaluate)
	handler.POST("/validate", presenter.Validate)
	handler.GET("/errors", presenter.GetErrors)

	logrus.Info("starting http server...")
	httpServer := &http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			handleError(err, "server returned an error")
		}
	}()

	app.WaitExitSignal()
	logrus.Info("shutting down the application")

}
