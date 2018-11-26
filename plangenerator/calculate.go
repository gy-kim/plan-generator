package plangenerator

import (
	"math"
)

const (
	daysInMonth = 30
	daysInYear  = 360
)

func calInterest(rate float64, initPrincipal float64) float64 {
	r := rate / 100
	v := r * daysInMonth * initPrincipal / daysInYear
	return round(v)
}

func calAnnuityPayment(rate float64, months int, initPrincipal float64) float64 {
	r := (rate / 100) / 12 // rate per period (5.00 / 100 / 12)
	pv := (1 - 1/square(1+r, months)) / r
	annuity := initPrincipal / pv

	return round(annuity)
}

func calPrincipalRemainPrincipal(annuity, interest, initPrincipal float64) (principal, remainPrincipal float64) {
	if annuity > initPrincipal {
		principal = initPrincipal
		remainPrincipal = 0.00
	} else {
		principal = round(annuity - interest)
		remainPrincipal = round(initPrincipal - principal)
	}
	return
}

func round(v float64) float64 {
	return math.Round(v*100) / 100
}

func square(v float64, times int) float64 {
	r := v
	for i := 0; i < times-1; i++ {
		r *= v
	}
	return r
}
