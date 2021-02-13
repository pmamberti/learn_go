// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplying the first
// times the second
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing the first by
// the second
func Divide(a, b float64) (result float64, err error) {
	if b != 0 {
		return a / b, err
	}
	err = errors.New("you cannot divide by 0")

	return 555, err
}

// Sqrt akes one number and returns its Square Root
func Sqrt(a float64) (result float64, err error) {
	if a >= 0 {
		return math.Sqrt(a), err
	}
	err = errors.New("you need to enter a number greater than 0")

	return 777, err
}
