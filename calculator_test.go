package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatorInitWithZeroDefault(t *testing.T) {
	testCalculator, err := New()

	assert.NoError(t, err)
	assert.Equal(t, testCalculator.GetLastCalculation(), 0.0)
}

func TestCalculatorInitWithPassedValue(t *testing.T) {
	testCalculator, err := New(1.0)

	assert.NoError(t, err)
	assert.Equal(t, testCalculator.GetLastCalculation(), 1.0)
}

func TestCalculatorInitFail(t *testing.T) {
	testCalculator, err := New(1.0, 23.3)
	assert.Error(t, err)
	assert.Nil(t, testCalculator)
}
func TestCalculatorAddPassPositiveInts(t *testing.T) {
	testCalculator, err := New()

	assert.NoError(t, err)

	expectedSum := 10.0
	actualSum, err := testCalculator.Add(4, 6)
	assert.NoError(t, err)

	assert.Equal(t, expectedSum, actualSum)
}

func TestCalculatorAddPassNegativeInts(t *testing.T) {
	testCalculator, err := New()

	assert.NoError(t, err)

	expectedSum := -5.0
	actualSum, err := testCalculator.Add(-3, -2)
	assert.NoError(t, err)

	assert.Equal(t, expectedSum, actualSum)
}

func TestCalculatorAddPassPositiveFloats(t *testing.T) {
	testCalculator, err := New()

	assert.NoError(t, err)
	expectedSum := 100.0
	actualSum, err := testCalculator.Add(64.5, 35.5)
	assert.NoError(t, err)

	assert.Equal(t, expectedSum, actualSum)

}
func TestCalculatorAddFailure(t *testing.T) {
	testCalculator, err := New()
	assert.NoError(t, err)

	sum1, err := testCalculator.Add(64.5, "FAILURE")

	assert.Equal(t, sum1, 0.0)
	assert.Error(t, err)

	sum2, err := testCalculator.Add("FAILURE", 64.5)

	assert.Equal(t, sum2, 0.0)
	assert.Error(t, err)
}
func TestCalculatorAddToLastValuePass(t *testing.T) {
	testCalculator, err := New(3.0)

	assert.NoError(t, err)
	expectedSum := 5.0
	actualSum, err := testCalculator.AddToLastCalculation(2.0)
	assert.NoError(t, err)

	assert.Equal(t, expectedSum, actualSum)
}
func TestCalculatorAddToLastValueFailure(t *testing.T) {
	testCalculator, err := New()
	assert.NoError(t, err)

	actualSum, err := testCalculator.AddToLastCalculation("FAILURE")

	assert.Equal(t, actualSum, 0.0)
	assert.Error(t, err)
}
func TestCalculatorSubtractPassPositiveInts(t *testing.T) {
	testCalculator, err := New()

	assert.NoError(t, err)

	expectedSum := 2.0
	actualSum, err := testCalculator.Subtract(6, 4)
	assert.NoError(t, err)

	assert.Equal(t, expectedSum, actualSum)
}

func TestCalculatorSubtractPassNegativeInts(t *testing.T) {
	testCalculator, err := New()

	assert.NoError(t, err)

	expectedSum := 1.0
	actualSum, err := testCalculator.Subtract(-3, -4)
	assert.NoError(t, err)

	assert.Equal(t, expectedSum, actualSum)
}

func TestCalculatorSubtractPassPositiveFloats(t *testing.T) {
	testCalculator, err := New()

	assert.NoError(t, err)
	expectedSum := 30.0
	actualSum, err := testCalculator.Subtract(64.5, 34.5)
	assert.NoError(t, err)

	assert.Equal(t, expectedSum, actualSum)

}

//Combined these to avoid things being a bit too robust
func TestCalculatorSubtractFailure(t *testing.T) {
	testCalculator, err := New()

	assert.NoError(t, err)

	val, err := testCalculator.Subtract(64.5, "FAILURE")
	assert.Equal(t, val, 0.0)
	assert.Error(t, err)

	val2, err := testCalculator.Subtract("FAILURE", 64.5)
	assert.Equal(t, val2, 0.0)
	assert.Error(t, err)

}

func TestCalculatorSubtractFromLastPass(t *testing.T) {
	testCalculator, err := New(0.0)

	assert.NoError(t, err)

	expectedSum := -64.5
	actualSum, err := testCalculator.SubtractFromLastCalculation(64.5)
	assert.NoError(t, err)

	assert.Equal(t, expectedSum, actualSum)

}

func TestCalculatorSubtractFromLastFail(t *testing.T) {
	testCalculator, err := New(0.0)

	assert.NoError(t, err)

	actualSum, err := testCalculator.SubtractFromLastCalculation("FAILURE")

	assert.Equal(t, actualSum, 0.0)
	assert.Error(t, err)
}
