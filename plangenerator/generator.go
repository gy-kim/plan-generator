package plangenerator

import (
	"errors"
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

const (
	ANNUITY uint8 = iota
)

type Generator interface {
	Calculate(loanAmount float64, rate float64, months int, startDate time.Time) []Payment
}

var notExistType error = errors.New("Not Exist type")

func GetGenerator(genType uint8) (Generator, error) {
	switch genType {
	case ANNUITY:
		return NewAnnuityPayment(), nil
	default:
		return nil, notExistType
	}
}

type AnnuityPayment struct{}

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
		principal, remainPrincipal, amount := a.calculatePrincipal(annuity, interest, initPrincipal)

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
