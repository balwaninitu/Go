package calc

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AddTestSuite struct {
	suite.Suite
}
type SubtractTestSuite struct {
	suite.Suite
}
type MultiplyTestSuite struct {
	suite.Suite
}
type DivideTestSuite struct {
	suite.Suite
}

// All methods that begin with "Test" are run as tests within a suite.
func (suite *AddTestSuite) TestAdd() {
	assert.Equal(suite.T(), 3, Add(1, 2))
	assert.Equal(suite.T(), 1, Add(1, 0))
	assert.Equal(suite.T(), 0, Add(2, -2))
}
func (suite *SubtractTestSuite) TestSubtract() {
	assert.Equal(suite.T(), 0, Subtract(1, 1))
	assert.Equal(suite.T(), 5, Subtract(10, 5))
	assert.Equal(suite.T(), 0, Subtract(-5, -5))
}
func (suite *MultiplyTestSuite) TestMultiply() {
	assert.Equal(suite.T(), 4, Multiply(1, 4))
	assert.Equal(suite.T(), 25, Multiply(5, 5))
	assert.Equal(suite.T(), 0, Multiply(100, 0))
}
func (suite *DivideTestSuite) TestDivide() {
	assert.Equal(suite.T(), float64(20), Divide(100, 5))
	assert.Equal(suite.T(), float64(0.1), Divide(100, 1000))
	assert.Equal(suite.T(), math.Inf(1), Divide(100, 0))
}

// This setups the test
func TestSuite(t *testing.T) {
	suite.Run(t, new(AddTestSuite))
	suite.Run(t, new(SubtractTestSuite))
	suite.Run(t, new(MultiplyTestSuite))
	suite.Run(t, new(DivideTestSuite))
}
