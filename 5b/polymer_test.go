package aoc

import (
	"io/ioutil"
	"path"
	"strings"
	"testing"
)

func TestExampleBest(t *testing.T) {
	in := "dabAcCaCBAcCcaDA"
	expected := "daDA"

	result := bestReaction(in)
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestExamples(t *testing.T) {
	tt := []struct {
		input    string
		expected string
	}{
		{"aA", ""},
		{"abBA", ""},
		{"abAB", "abAB"},
		{"aabAAB", "aabAAB"},
		{"dabAcCaCBAcCcaDA", "dabCBAcaDA"},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			result := react(tc.input)
			if result != tc.expected {
				t.Errorf("expected %q got %q", tc.expected, result)
			}
		})
	}
}

func TestExam(t *testing.T) {
	str, err := ioutil.ReadFile(path.Join("testdata", "input.txt"))
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	result := bestReaction(strings.TrimSuffix(string(str), "\n"))
	t.Logf("final answer is %d", len(result))
}
