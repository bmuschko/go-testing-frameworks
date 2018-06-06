package calc_goconvey_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/bmuschko/go-testing-frameworks/calc"
)

func TestAdd(t *testing.T) {
	Convey("Adding two numbers", t, func() {
		x := int64(1)
		y := int64(2)

		Convey("should produce the expected result", func() {
			So(Add(x, y), ShouldEqual, int64(3))
		})
	})
}

func TestSubstract(t *testing.T) {
	Convey("Substracting two numbers", t, func() {
		x := int64(5)
		y := int64(3)

		Convey("should produce the expected result", func() {
			So(Substract(x, y), ShouldEqual, int64(2))
		})
	})
}

func TestMultiply(t *testing.T) {
	Convey("Multiplying two numbers", t, func() {
		x := int64(5)
		y := int64(6)

		Convey("should produce the expected result", func() {
			So(Multiply(x, y), ShouldEqual, int64(30))
		})
	})
}

func TestDivide(t *testing.T) {
	Convey("Dividing two numbers", t, func() {
		x := int64(10)
		y := int64(2)

		Convey("should produce the expected result", func() {
			So(Divide(x, y), ShouldEqual, float64(5))
		})
	})
}

