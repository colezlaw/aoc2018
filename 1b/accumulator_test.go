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
	a := New()

	// act

	// assert
	if a.freq != 0 {
		t.Errorf("Expected 0 got %d", a.freq)
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
	a := New()
	tt := []struct {
		input    string
		expected int64
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
		if a.freq != tc.expected {
			t.Errorf("given %s, expected a to be %d, got %d", tc.input, tc.expected, a.freq)
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

	a := New()
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

func TestExamples(t *testing.T) {
	tt := []struct {
		input    []string
		expected int64
	}{
		{[]string{"+1", "-2", "+3", "+1"}, 2},
		{[]string{"+1", "-1"}, 0},
		{[]string{"+3", "+3", "+4", "-2", "-4"}, 10},
		{[]string{"-6", "+3", "+8", "+5", "-6"}, 5},
		{[]string{"+7", "+7", "-2", "-7", "-4"}, 14},
	}

	for x, tc := range tt {
		t.Run(fmt.Sprintf("%d", x), func(t *testing.T) {
			a := New()
			t.Logf("Got here")
			for n := 0; !a.found && n < 100; n++ {
				a.Acc(tc.input[n%len(tc.input)])
			}

			if tc.expected != a.freq {
				t.Errorf("expected %d, got %d", tc.expected, a.freq)
			}
		})
	}
}

func TestFinal(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	f, err := os.Open("testdata/testdata.txt")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	a := New()

	for !a.found {
		fmt.Println("opening")
		s := bufio.NewScanner(f)
		for s.Scan() {
			err := a.Acc(s.Text())
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}
			if a.found {
				break
			}
		}
		f.Seek(0, 0)
	}

	t.Logf("Final value %v", a.freq)
}
