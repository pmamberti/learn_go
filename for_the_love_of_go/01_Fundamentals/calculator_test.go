package calculator_test

import (
	"calculator"
	"fmt"
	"testing"
)

type testCase struct {
	a, b        float64
	want        float64
	description string
	errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4, description: "Adding 2 + 2 must return 4"},
		{a: 1, b: -1, want: 0, description: "Adding 1 and -1 must return 0"},
		{a: 5, b: 0, want: 5, description: "Adding 0 must not change a"},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			fmt.Printf("Expected test: %s\n", tc.description)
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 5, b: 3, want: 2, description: "5 - 3 = 2"},
		{a: 1, b: 2, want: -1, description: "1 - 2 = -1"},
		{a: 5, b: 0, want: 5, description: "5 - 0 = 5"},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			fmt.Printf("Expected test: %s\n", tc.description)
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	testCases := []testCase{
		{a: 5, b: 3, want: 15, description: "5 * 3 = 15"},
		{a: 1, b: 2, want: 2, description: "1 * 2 = 2"},
		{a: 5, b: 0, want: 0, description: "5 * 0 = 0"},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			fmt.Printf("Expected test: %s\n", tc.description)
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {

	testCases := []testCase{
		{a: 6, b: 0, want: 999, description: "6 / 0 = Nan", errExpected: false},
		{a: 5, b: 0, want: 0, description: "5 / 0 = NaN", errExpected: true},
		{a: 5, b: 3, want: 1.666, description: "5 / 3 = 1.666", errExpected: true},
		{a: 1, b: 2, want: 2, description: "1 / 2 = 0.5", errExpected: false},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Errorf("Divide(%f, %f: Expected Error Status: %v - Received Error status: %v",
				tc.a, tc.b, tc.errExpected, err)
		}

		fmt.Printf("%f", got)
	}
}
