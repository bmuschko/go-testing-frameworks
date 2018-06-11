package calc_test

import (
	. "github.com/bmuschko/go-testing-frameworks/calc"
	. "github.com/go-check/check"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestAddWithGocheck(c *C) {
	result := Add(1, 2)
	c.Assert(result, Equals, 3)
}

func (s *MySuite) TestSubtractWithGocheck(c *C) {
	result := Subtract(5, 3)
	c.Assert(result, Equals, 2)
}

func (s *MySuite) TestMultiplyWithGocheck(c *C) {
	result := Multiply(5, 6)
	c.Assert(result, Equals, 30)
}

func (s *MySuite) TestDivideWithGocheck(c *C) {
	result := Divide(10, 2)
	c.Assert(result, Equals, float64(5))
}
