package handler

import (
	"net/http"
	"time"

	"github.com/gy-kim/plan-generator/applogger"
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

	return c.JSON(http.StatusCreated, p)
}
