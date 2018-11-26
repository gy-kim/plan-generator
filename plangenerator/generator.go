package plangenerator

import (
	"errors"
	"time"
)

type Payment struct {
	Amount          float64   `json:"borrowerPaymentAmount"`
	Date            time.Time `json:"date"`
	InitPrincipal   float64   `json:"initialOutstandingPrincipal"`
	Interest        float64   `json:"interest"`
	Principal       float64   `json:"principal"`
	RemainPrincipal float64   `json:"remainingOutstandingPrincipal"`
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
