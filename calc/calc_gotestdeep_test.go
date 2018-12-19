package calc_test

import (
	. "github.com/bmuschko/go-testing-frameworks/calc"
	. "github.com/maxatome/go-testdeep"
	"testing"
)

func TestAddWithGoTestDeep(t *testing.T) {
	result := Add(1, 2)
	CmpNotZero(t, result)
	CmpDeeply(t, &result, Ptr(3))
	CmpDeeply(t, result, Code(func (r int) (bool, string) {
		if r == 3 {
			return true, ""
		}
		return false, "Result should be 3"
	}))
}

func TestSubtractWithGoTestDeep(t *testing.T) {
	result := Subtract(5, 3)
	CmpNotZero(t, result)
	CmpDeeply(t, &result, Ptr(2))
	CmpDeeply(t, result, Code(func (r int) (bool, string) {
		if r == 2 {
			return true, ""
		}
		return false, "Result should be 2"
	}))
}

func TestMultiplyWithGoTestDeep(t *testing.T) {
	result := Multiply(5, 6)
	CmpNotZero(t, result)
	CmpDeeply(t, &result, Ptr(30))
	CmpDeeply(t, result, Code(func (r int) (bool, string) {
		if r == 30 {
			return true, ""
		}
		return false, "Result should be 30"
	}))
}

func TestDivideWithGoTestDeep(t *testing.T) {
	result := Divide(10, 2)
	CmpNotZero(t, result)
	CmpDeeply(t, &result, Ptr(float64(5)))
	CmpDeeply(t, result, Code(func (r float64) (bool, string) {
		if r == float64(5) {
			return true, ""
		}
		return false, "Result should be 5"
	}))
}

