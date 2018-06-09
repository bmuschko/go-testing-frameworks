package calc_testify_test

import (
	"testing"
	. "github.com/stretchr/testify/assert"
	. "github.com/bmuschko/go-testing-frameworks/calc"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	Equal(t, 3, result)
}

func TestSubstract(t *testing.T) {
	result := Substract(5, 3)
	Equal(t, 2, result)
}

func TestMultiply(t *testing.T) {
	result := Multiply(5, 6)
	Equal(t, 30, result)
}

func TestDivide(t *testing.T) {
	result := Divide(10, 2)
	Equal(t, float64(5), result)
}