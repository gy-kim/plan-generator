package plangenerator

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestCalculate(t *testing.T) {
	g := NewGenerator()
	loanAmount := 5000.00
	rate := 5.00
	months := 24
	dt := "2018-01-01T00:00:01Z"
	tm, _ := time.Parse(time.RFC3339, dt)

	payments := g.Calculate(loanAmount, rate, months, tm)
	b, _ := json.Marshal(payments)

	fmt.Println(string(b))
}
