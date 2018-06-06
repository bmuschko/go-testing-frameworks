package calc_goblin_test

import (
	"testing"
	. "github.com/franela/goblin"
	. "github.com/bmuschko/go-testing-frameworks/calc"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Calculator", func() {
		g.It("should add two numbers ", func() {
			g.Assert(Add(1, 2)).Equal(int64(3))
		})

		g.It("should substract two numbers", func() {
			g.Assert(Substract(5, 3)).Equal(int64(2))
		})

		g.It("should multiply two numbers", func() {
			g.Assert(Multiply(5, 6)).Equal(int64(30))
		})

		g.It("should divide two numbers", func() {
			g.Assert(Divide(10, 2)).Equal(float64(5))
		})
	})
}