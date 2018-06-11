package calc_test

import (
	. "github.com/bmuschko/go-testing-frameworks/calc"
	"testing"
)

func TestAddWithTestingPackage(t *testing.T) {
	result := Add(1, 2)

	if result != 3 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 3)
	}
}

func TestSubtractWithTestingPackage(t *testing.T) {
	result := Subtract(5, 3)

	if result != 2 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 2)
	}
}

func TestMultiplyWithTestingPackage(t *testing.T) {
	result := Multiply(5, 6)

	if result != 30 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 30)
	}
}

func TestDivideWithTestingPackage(t *testing.T) {
	result := Divide(10, 2)

	if result != 5 {
		t.Errorf("Result was incorrect, got: %f, want: %f.", result, float64(5))
	}
}
