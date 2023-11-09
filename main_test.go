package main

import (
	"math"
	"testing"
)

const float64EqualityThreshold = 1e-9

func validateExpression(t *testing.T, expression string, expected float64) {
	result, err := calculate(expression)
	if !(math.Abs(result-expected) <= float64EqualityThreshold) || err != nil {
		t.Fatalf("Expected %f, got %f", expected, result)
	}
}

func TestAddExpression(t *testing.T) {
	validateExpression(t, "1 + 1", 2)
}

func TestDivideExpression(t *testing.T) {
	validateExpression(t, "10 / 2", 5)
}

func TestMultiplyExpression(t *testing.T) {
	validateExpression(t, "10 * 10", 100)
}
