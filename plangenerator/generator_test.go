package plangenerator

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestCalculate(t *testing.T) {
	g := NewAnnuityPayment()
	loan := 5000.00
	rate := 5.00
	months := 24
	st, _ := time.Parse(time.RFC3339, "2018-01-01T00:00:01Z")

	payments := g.Calculate(loan, rate, months, st)
	b, _ := json.Marshal(payments)

	fmt.Println(string(b))
}
