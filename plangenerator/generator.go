package plangenerator

import (
	"time"
)

type Payment struct {
	Amount          float64
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
	annuity := calculateAnnuity(rate, months, loanAmount)

	initPrincipal := loanAmount

	for i := 0; i < months; i++ {
		interest := calculateInterest(rate, initPrincipal)
		principal, remainPrincipal, amount := calculatePrincipal(annuity, interest, initPrincipal)

		payments[i] = Payment{
			Amount:          amount,
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
