package handler

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gy-kim/plan-generator/applogger"
	"github.com/gy-kim/plan-generator/plangenerator"
	"github.com/labstack/echo"
)

var logger = applogger.GetInstance()

type Payload struct {
	LoanAmount string    `json:"loanAmount"`
	Rate       string    `json:"nominalRate"`
	Duration   int       `json:"duration"`
	StartDate  time.Time `json:"startDate"`
}

func GeneratePlan(c echo.Context) error {
	p := new(Payload)
	if err := c.Bind(p); err != nil {
		logger.Println(err)
		return err
	}

	gen, err := plangenerator.GetGenerator(plangenerator.ANNUITY)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("internal server error"))
	}

	loanAmount, _ := strconv.ParseFloat(p.LoanAmount, 64)
	rate, _ := strconv.ParseFloat(p.Rate, 64)
	payments := gen.Calculate(loanAmount, rate, p.Duration, p.StartDate)

	return c.JSON(http.StatusCreated, payments)
}
