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
