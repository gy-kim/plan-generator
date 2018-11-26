package plangenerator

import (
	"math"
	"time"
)

type Payment struct {
	Amount          float64 `json:"borrowerPaymentAmount"`
	Date            time.Time
	InitPrincipal   float64
	Interest        float64
	Principal       float64
	RemainPrincipal float64
}

type Generator struct{}

// NewGenerator returns Generator pointer
func NewGenerator() *Generator {
	return new(Generator)
}

// Calculate returns Payment slice
func (g *Generator) Calculate(loanAmount float64, rate float64, months int, startDate time.Time) []Payment {
	payments := make([]Payment, months)
	annuity := calAnnuityPayment(rate, months, loanAmount)

	initPrincipal := loanAmount
	remainPrincipal := 0.00

	for i := 0; i < months; i++ {
		interest := calInterest(rate, initPrincipal)
		principal := 0.00

		if i == (months - 1) {
			principal = remainPrincipal
			remainPrincipal = 0.00
		} else {
			principal = calPrincipal(annuity, interest)
			remainPrincipal = calRemainPrincipal(initPrincipal, principal)
		}

		payments[i] = Payment{
			Amount:          math.Round((interest+principal)*100) / 100,
			InitPrincipal:   initPrincipal,
			Date:            startDate.AddDate(0, i, 0),
			Interest:        interest,
			Principal:       principal,
			RemainPrincipal: remainPrincipal,
		}

		initPrincipal = remainPrincipal

	}

	return payments
}
