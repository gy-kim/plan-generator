package plangenerator

import (
	"time"
)

type AnnuityPayment struct{}

const (
	daysInMonth = 30
	daysInYear  = 360
)

// NewAnnuityPayment returns payment pointer
func NewAnnuityPayment() *AnnuityPayment {
	return new(AnnuityPayment)
}

// Calculate returns Payment slice
func (a AnnuityPayment) Calculate(loanAmount float64, rate float64, months int, startDate time.Time) []Payment {
	payments := make([]Payment, months)
	annuity := a.calculateAnnuity(rate, months, loanAmount)

	initPrincipal := loanAmount

	for i := 0; i < months; i++ {
		interest := a.calculateInterest(rate, initPrincipal)
		principal, remainPrincipal, amount := a.calculate(annuity, interest, initPrincipal)

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

func (a AnnuityPayment) calculateInterest(rate float64, initPrincipal float64) float64 {
	r := rate / 100
	v := r * daysInMonth * initPrincipal / daysInYear
	return round(v)
}

func (a AnnuityPayment) calculateAnnuity(rate float64, months int, initPrincipal float64) float64 {
	r := (rate / 100) / 12 // rate per period (5.00 / 100 / 12)
	pv := (1 - 1/square(1+r, months)) / r
	annuity := initPrincipal / pv

	return round(annuity)
}

func (a AnnuityPayment) calculate(annuity, interest, initPrincipal float64) (principal, remainPrincipal, amount float64) {
	if annuity > initPrincipal {
		principal = initPrincipal
		remainPrincipal = 0.00
	} else {
		principal = round(annuity - interest)
		remainPrincipal = round(initPrincipal - principal)
	}
	amount = round(interest + principal)
	return
}
