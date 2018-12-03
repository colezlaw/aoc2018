package overlap

import (
	"bufio"
	"os"
	"path"
	"testing"
)

func TestParse(t *testing.T) {
	tt := []struct {
		in       string
		expected square
	}{
		{"#1 @ 1,3: 4x4", square{id: 1, left: 1, top: 3, w: 4, h: 4}},
		{"#2 @ 3,1: 4x4", square{id: 2, left: 3, top: 1, w: 4, h: 4}},
		{"#3 @ 5,5: 2x2", square{id: 3, left: 5, top: 5, w: 2, h: 2}},
	}

	for _, tc := range tt {
		t.Run(tc.in, func(t *testing.T) {
			actual, err := parse(tc.in)
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}
			if actual != tc.expected {
				t.Errorf("expected %#v, got %#v", tc.expected, actual)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tc := []square{
		square{id: 1, left: 1, top: 3, w: 4, h: 4},
		square{id: 2, left: 3, top: 1, w: 4, h: 4},
		square{id: 3, left: 5, top: 5, w: 2, h: 2},
	}

	x, y := maxDimensions(tc)
	if x != 8 || y != 8 {
		t.Errorf("expected (8, 8) got (%d, %d)", x, y)
	}
}

func TestRenderGrid(t *testing.T) {
	tc := []square{
		square{id: 1, left: 1, top: 3, w: 4, h: 4},
		square{id: 2, left: 3, top: 1, w: 4, h: 4},
		square{id: 3, left: 5, top: 5, w: 2, h: 2},
	}

	expected := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 2, 2, 2, 0},
		{0, 0, 0, 2, 2, 2, 2, 0},
		{0, 1, 1, -1, -1, 2, 2, 0},
		{0, 1, 1, -1, -1, 2, 2, 0},
		{0, 1, 1, 1, 1, 3, 3, 0},
		{0, 1, 1, 1, 1, 3, 3, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	result := renderGrid(8, 8, tc)

	if len(result) != len(expected) {
		t.Fatalf("expected len to be %d, got %d", len(expected), len(result))
	}

	if len(result[0]) != len(expected[0]) {
		t.Fatalf("expected len to be %d, got %d", len(expected[0]), len(result[0]))
	}

	for y, r := range result {
		for x, c := range r {
			if c != expected[y][x] {
				t.Errorf("At (%d,%d) expected %d got %d", x, y, expected[y][x], c)
			}
		}
	}
}

func TestGoodClaim(t *testing.T) {
	tc := []square{
		square{id: 1, left: 1, top: 3, w: 4, h: 4},
		square{id: 2, left: 3, top: 1, w: 4, h: 4},
		square{id: 3, left: 5, top: 5, w: 2, h: 2},
	}

	output := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 2, 2, 2, 0},
		{0, 0, 0, 2, 2, 2, 2, 0},
		{0, 1, 1, -1, -1, 2, 2, 0},
		{0, 1, 1, -1, -1, 2, 2, 0},
		{0, 1, 1, 1, 1, 3, 3, 0},
		{0, 1, 1, 1, 1, 3, 3, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	result := findGoodSquare(output, tc)
	if result != 3 {
		t.Errorf("expected 3 got %d", result)
	}
}

func TestExam(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	f, err := os.Open(path.Join("testdata", "input.txt"))
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	defer f.Close()

	grid := make([]square, 0, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		sq, err := parse(s.Text())
		if err != nil {
			t.Fatalf("unexpected error %v", err)
		}
		grid = append(grid, sq)
	}

	x, y := maxDimensions(grid)
	result := renderGrid(x, y, grid)
	goodClaim := findGoodSquare(result, grid)

	t.Logf("Final answer %d", goodClaim)
}
