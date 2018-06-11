package calc_goconvey_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/bmuschko/go-testing-frameworks/calc"
)

func TestAdd(t *testing.T) {
	Convey("Adding two numbers", t, func() {
		x := 1
		y := 2

		Convey("should produce the expected result", func() {
			So(Add(x, y), ShouldEqual, 3)
		})
	})
}

func TestSubtract(t *testing.T) {
	Convey("Subtracting two numbers", t, func() {
		x := 5
		y := 3

		Convey("should produce the expected result", func() {
			So(Subtract(x, y), ShouldEqual, 2)
		})
	})
}

func TestMultiply(t *testing.T) {
	Convey("Multiplying two numbers", t, func() {
		x := 5
		y := 6

		Convey("should produce the expected result", func() {
			So(Multiply(x, y), ShouldEqual, 30)
		})
	})
}

func TestDivide(t *testing.T) {
	Convey("Dividing two numbers", t, func() {
		x := 10
		y := 2

		Convey("should produce the expected result", func() {
			So(Divide(x, y), ShouldEqual, float64(5))
		})
	})
}

