package calc_test

import (
	. "github.com/bmuschko/go-testing-frameworks/calc"
	. "github.com/ToQoz/gopwt/assert"
	. "github.com/ToQoz/gopwt"
	"testing"
	"flag"
	"os"
)

func TestMain(m *testing.M) {
	flag.Parse()
	Empower()
	os.Exit(m.Run())
}

func TestAddWithGopwt(t *testing.T) {
	result := Add(1, 2)
	OK(t, 3 == result)
}

func TestSubtractWithGopwt(t *testing.T) {
	result := Subtract(5, 3)
	OK(t, 2 == result)
}

func TestMultiplyWithGopwt(t *testing.T) {
	result := Multiply(5, 6)
	OK(t, 30 == result)
}

func TestDivideWithGopwt(t *testing.T) {
	result := Divide(10, 2)
	OK(t, float64(5) == result)
}