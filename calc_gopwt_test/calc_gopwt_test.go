package calc_gopwt_test

import (
	"testing"
	. "github.com/ToQoz/gopwt/assert"
	. "github.com/bmuschko/go-testing-frameworks/calc"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	OK(t, int64(3) == result)
}

func TestSubstract(t *testing.T) {
	result := Substract(5, 3)
	OK(t, int64(2) == result)
}

func TestMultiply(t *testing.T) {
	result := Multiply(5, 6)
	OK(t, int64(30) == result)
}

func TestDivide(t *testing.T) {
	result := Divide(10, 2)
	OK(t, float64(5) == result)
}