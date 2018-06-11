package calc_test

import (
	. "github.com/bmuschko/go-testing-frameworks/calc"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAddWithGoConvey(t *testing.T) {
	Convey("Adding two numbers", t, func() {
		x := 1
		y := 2

		Convey("should produce the expected result", func() {
			So(Add(x, y), ShouldEqual, 3)
		})
	})
}

func TestSubtractWithGoConvey(t *testing.T) {
	Convey("Subtracting two numbers", t, func() {
		x := 5
		y := 3

		Convey("should produce the expected result", func() {
			So(Subtract(x, y), ShouldEqual, 2)
		})
	})
}

func TestMultiplyWithGoConvey(t *testing.T) {
	Convey("Multiplying two numbers", t, func() {
		x := 5
		y := 6

		Convey("should produce the expected result", func() {
			So(Multiply(x, y), ShouldEqual, 30)
		})
	})
}

func TestDivideWithGoConvey(t *testing.T) {
	Convey("Dividing two numbers", t, func() {
		x := 10
		y := 2

		Convey("should produce the expected result", func() {
			So(Divide(x, y), ShouldEqual, float64(5))
		})
	})
}
