package calc_gocheck_test

import (
	"testing"
	. "github.com/go-check/check"
	. "github.com/bmuschko/go-testing-frameworks/calc"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestAdd(c *C) {
	result := Add(1, 2)
	c.Assert(result, Equals, int64(3))
}

func (s *MySuite) TestSubstract(c *C) {
	result := Substract(5, 3)
	c.Assert(result, Equals, int64(2))
}

func (s *MySuite) TestMultiply(c *C) {
	result := Multiply(5, 6)
	c.Assert(result, Equals, int64(30))
}

func (s *MySuite) TestDivide(c *C) {
	result := Divide(10, 2)
	c.Assert(result, Equals, float64(5))
}