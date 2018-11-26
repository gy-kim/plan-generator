package plangenerator

import (
	"testing"
)

func TestGetGenerator(t *testing.T) {
	sut, _ := GetGenerator(ANNUITY)
	if _, ok := sut.(*AnnuityPayment); !ok {
		t.Error("It's not annuityPayment instanance")
	}
}
