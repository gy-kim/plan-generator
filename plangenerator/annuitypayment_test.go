package plangenerator

import (
	"sync"
	"testing"
)

var once sync.Once

var annuity *AnnuityPayment

func getAnnuity() {
	annuity = NewAnnuityPayment()
}

func TestCalculateInterest(t *testing.T) {
	once.Do(getAnnuity)

	expect := 20.83
	sut := annuity.calculateInterest(5.00, 5000.00)
	if expect != sut {
		t.Fatalf("interest doesn't match with expect. expect(%f), interest(%f)", expect, sut)
	}
}

func TestCalculateAnnuity(t *testing.T) {
	once.Do(getAnnuity)

	expect := 219.36
	sut := annuity.calculateAnnuity(5.00, 24, 5000.00)
	if expect != sut {
		t.Fatalf("Annuity doesn't match with expect. expect(%f), annuity(%f).", expect, sut)
	}
}
