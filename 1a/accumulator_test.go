package aoc

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// Assume the value of a new Accumulator is zero
func TestZero(t *testing.T) {
	// arrange
	var a Accumulator

	// act

	// assert
	if int64(a) != 0 {
		t.Errorf("Expected 0 got %d", a)
	}
}

// Tests the parser
func TestParse(t *testing.T) {
	tt := []struct {
		input    string
		expected int
		err      error
	}{
		{"+1", 1, nil},
		{"-1", -1, nil},
		{"3", 0, BadInputError("3")},
		{"+194", 194, nil},
		{"-194", -194, nil},
		{"+124a", 124, BadInputError("+124a")},
		{"-124a", 124, BadInputError("-124a")},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			actual, err := parse(tc.input)
			if tc.err == nil && err != nil {
				t.Fatalf("unexpected error %v", err)
				return
			}
			if tc.err != nil && err == nil {
				t.Fatalf("expected error %v, got none", tc.err)
				return
			}
			if tc.err != err {
				t.Fatalf("expected error %v, got %v", tc.err, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func TestAccumulate(t *testing.T) {
	var a Accumulator
	tt := []struct {
		input    string
		expected int
	}{
		{"+1", 1},
		{"-2", -1},
		{"+3", 2},
		{"+1", 3},
	}

	for _, tc := range tt {
		fmt.Printf("a is now %v\n", a)
		err := a.Acc(tc.input)
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		if int64(a) != int64(tc.expected) {
			t.Errorf("given %s, expected a to be %d, got %d", tc.input, tc.expected, a)
		}
	}
}

func TestInput(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	f, err := os.Open("testdata/testdata.txt")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	var a Accumulator
	s := bufio.NewScanner(f)
	for s.Scan() {
		err := a.Acc(s.Text())
		fmt.Println(s.Text())
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
	}

	t.Logf("Final value %v", a)
}
