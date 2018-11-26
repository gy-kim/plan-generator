package plangenerator

import (
	"testing"
)

func TestCalInterest(t *testing.T) {
	expect := 20.83
	sut := calInterest(5.00, 5000.00)
	if expect != sut {
		t.Fatalf("interest doesn't match with expect. expect(%f), interest(%f)", expect, sut)
	}
}

func TestCalAnnuityPayment(t *testing.T) {
	expect := 219.36
	sut := calAnnuityPayment(5.00, 24, 5000.00)
	if expect != sut {
		t.Fatalf("Annuity doesn't match with expect. expect(%f), annuity(%f).", expect, sut)
	}
}
