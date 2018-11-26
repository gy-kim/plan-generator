package plangenerator

import "math"

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
