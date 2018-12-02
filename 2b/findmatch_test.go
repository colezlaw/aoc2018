package aoc

import (
	"bufio"
	"os"
	"path"
	"testing"
)

func TestExample(t *testing.T) {
	in := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}

	match := findMatch(in)
	if match != "fgij" {
		t.Errorf("expected fgij got %q", match)
	}
}

func TestQuiz(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	var in []string

	f, err := os.Open(path.Join("testdata", "testdata.txt"))
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		in = append(in, s.Text())
	}

	result := findMatch(in)
	t.Logf("Found result %q", result)
}
