package plangenerator

import "testing"

func TestSquare(t *testing.T) {
	expect := 9.00
	sut := square(3.00, 2)
	if expect != sut {
		t.Fatal("Square function failed")
	}
}
