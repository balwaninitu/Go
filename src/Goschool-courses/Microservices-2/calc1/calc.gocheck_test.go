package calc

import (
	"math"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestAdd(c *C) {
	result := Add(1, 2)
	c.Assert(result, Equals, 3)

	result = Add(1, 0)
	c.Assert(result, Equals, 1)

	result = Add(2, -2)
	c.Assert(result, Equals, 0)

}

func (s *MySuite) TestSubtract(c *C) {
	result := Subtract(1, 1)
	c.Assert(result, Equals, 0)

	result = Subtract(10, 5)
	c.Assert(result, Equals, 5)

	result = Subtract(-5, -5)
	c.Assert(result, Equals, 0)
}

func (s *MySuite) TestMultiply(c *C) {
	result := Multiply(1, 5)
	c.Assert(result, Equals, 5)

	result = Multiply(5, 5)
	c.Assert(result, Equals, 25)

	result = Multiply(100, 0)
	c.Assert(result, Equals, 0)
}

func (s *MySuite) TestDivide(c *C) {
	result := Divide(100, 5)
	c.Assert(result, Equals, float64(20))

	result = Divide(100, 1000)
	c.Assert(result, Equals, float64(0.1))

	result = Divide(100, 0)
	c.Assert(result, Equals, math.Inf(1))
}
