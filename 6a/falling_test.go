package aoc

import (
	"os"
	"path"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tt := []struct {
		input    string
		expected dest
	}{
		{"1, 1", dest{x: 1, y: 1}},
		{"1, 6", dest{x: 1, y: 6}},
		{"8, 3", dest{x: 8, y: 3}},
		{"3, 4", dest{x: 3, y: 4}},
		{"5, 5", dest{x: 5, y: 5}},
		{"8, 9", dest{x: 8, y: 9}},
		{"213, 967", dest{x: 213, y: 967}},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			r := parse(tc.input)
			if r.x != tc.expected.x {
				t.Errorf("expected x of %d got %d", tc.expected.x, r.x)
			}
			if r.y != tc.expected.y {
				t.Errorf("expected t of %d got %d", tc.expected.y, r.y)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	input := `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`
	expected := []dest{
		{label: 'A', x: 1, y: 1},
		{label: 'B', x: 1, y: 6},
		{label: 'C', x: 8, y: 3},
		{label: 'D', x: 3, y: 4},
		{label: 'E', x: 5, y: 5},
		{label: 'F', x: 8, y: 9},
	}
	xMinx := 1
	xMiny := 1
	xMaxx := 8
	xMaxy := 9

	dests, minx, miny, maxx, maxy := load(strings.NewReader(input))

	if minx != xMinx {
		t.Errorf("expected X min of %d got %d", xMinx, minx)
	}
	if miny != xMiny {
		t.Errorf("expected Y min of %d got %d", xMiny, miny)
	}
	if maxx != xMaxx {
		t.Errorf("expected X max of %d got %d", xMaxx, maxx)
	}
	if maxy != xMaxy {
		t.Errorf("expected Y max of %d got %d", xMaxy, maxy)
	}
	if len(dests) != len(expected) {
		t.Errorf("expected len to be %d, got %d", len(expected), len(dests))
	}
}

func TestGrid(t *testing.T) {
	expected := [][]marker{
		[]marker{
			{distance: 0, label: 'A'},
			{distance: 1, label: 'A'},
			{distance: 2, label: 'A'},
			{distance: 3, label: 'A'},
			{distance: 4, label: '.'},
			{distance: 4, label: 'C'},
			{distance: 3, label: 'C'},
			{distance: 2, label: 'C'},
		},
		[]marker{
			{distance: 1, label: 'A'},
			{distance: 2, label: 'A'},
			{distance: 2, label: 'D'},
			{distance: 3, label: 'D'},
			{distance: 3, label: 'E'},
			{distance: 3, label: 'C'},
			{distance: 2, label: 'C'},
			{distance: 1, label: 'C'},
		},
		[]marker{
			{distance: 2, label: 'A'},
			{distance: 2, label: 'D'},
			{distance: 1, label: 'D'},
			{distance: 2, label: 'D'},
			{distance: 2, label: 'E'},
			{distance: 2, label: 'C'},
			{distance: 1, label: 'C'},
			{distance: 0, label: 'C'},
		},
		[]marker{
			{distance: 2, label: '.'},
			{distance: 1, label: 'D'},
			{distance: 0, label: 'D'},
			{distance: 1, label: 'D'},
			{distance: 1, label: 'E'},
			{distance: 2, label: 'E'},
			{distance: 2, label: 'C'},
			{distance: 1, label: 'C'},
		},
	}

	result := renderGrid(map[rune]dest{
		'A': {label: 'A', x: 1, y: 1},
		'B': {label: 'B', x: 1, y: 6},
		'C': {label: 'C', x: 8, y: 3},
		'D': {label: 'D', x: 3, y: 4},
		'E': {label: 'E', x: 5, y: 5},
		'F': {label: 'F', x: 8, y: 9},
	}, 1, 1, 8, 9)

	for r := range expected {
		for c := range expected[r] {
			if result[r][c] != expected[r][c] {
				t.Errorf("at (%d, %d) expected %#v, got %#v", c, r, expected[r][c], result[r][c])
			}
		}
	}
}

func TestExam(t *testing.T) {
	f, err := os.Open(path.Join("testdata", "input.txt"))
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	defer f.Close()

	targets, minx, miny, maxx, maxy := load(f)
	grid := renderGrid(targets, minx, miny, maxx, maxy)
	label, size := areaSizes(grid, targets)
	t.Logf("Largest area is %c with a size of %d", label, size)
}
