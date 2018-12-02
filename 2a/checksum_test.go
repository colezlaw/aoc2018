package aoc

import (
	"bufio"
	"os"
	"path"
	"testing"
)

func TestParseOne(t *testing.T) {
	tt := []struct {
		input    string
		expected Check
	}{
		{"abcdef", Check{three: false, two: false}},
		{"bababc", Check{three: true, two: true}},
		{"abbcde", Check{three: false, two: true}},
		{"abcccd", Check{three: true, two: false}},
		{"aabcdd", Check{three: false, two: true}},
		{"abcdee", Check{three: false, two: true}},
		{"ababab", Check{three: true, two: false}},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			actual, err := Parse(tc.input)
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}

			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestExample(t *testing.T) {
	input := []Check{
		Check{three: false, two: false},
		Check{three: true, two: true},
		Check{three: false, two: true},
		Check{three: true, two: false},
		Check{three: false, two: true},
		Check{three: false, two: true},
		Check{three: true, two: false},
	}
	actual := Checksum(input)

	if actual != 12 {
		t.Errorf("expected 12 got %d", actual)
	}
}

func TestQuiz(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	f, err := os.Open(path.Join("testdata", "testdata.txt"))
	if err != nil {
		t.Fatalf("unable to open testdata %v", err)
	}
	defer f.Close()

	checks := make([]Check, 10, 10)
	s := bufio.NewScanner(f)
	for s.Scan() {
		c, err := Parse(s.Text())
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		checks = append(checks, c)
	}

	t.Logf("Final answer %d", Checksum(checks))
}
