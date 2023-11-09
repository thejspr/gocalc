package main

import (
	"math"
	"reflect"
	"testing"
)

const float64EqualityThreshold = 1e-9

func TestAddExpression(t *testing.T) {
	validateExpression(t, "1 + 1", 2)
}

func TestDivideExpression(t *testing.T) {
	validateExpression(t, "10 / 2", 5)
}

func TestMultiplyExpression(t *testing.T) {
	validateExpression(t, "10 * 10", 100)
}

func TestMultiLineExpression(t *testing.T) {
	validateExpression(t, "10 * 10\n", 100)
}

// multiple expressions

func TestMulticalc(t *testing.T) {
	var expected = []float64{2, 5}
	result, err := Multicalc([]string{"1 + 1", "10 / 2"})

	if !reflect.DeepEqual(result, expected) || err != nil {
		t.Fatal(err)
	}
}

// helpers

func validateExpression(t *testing.T, expression string, expected float64) {
	result, err := Calculate(expression)
	if !(math.Abs(result-expected) <= float64EqualityThreshold) || err != nil {
		t.Fatalf("Expected %f, got %f", expected, result)
	}
}
