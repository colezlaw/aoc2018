package overlap

import "fmt"

var format = "#%d @ %d,%d: %dx%d"

type square struct {
	id   int
	left int
	top  int
	w    int
	h    int
}

type formatError struct {
	in    string
	cause error
}

func (fe formatError) Error() string {
	if fe.cause != nil {
		return fmt.Sprintf("%s: %s", fe.in, fe.cause)
	}
	return fe.in
}

func findGoodSquare(grid [][]int, squares []square) int {
outer:
	for _, sq := range squares {
		found := true

		for y := sq.top; y < sq.top+sq.h; y++ {
			for x := sq.left; x < sq.left+sq.w; x++ {
				if grid[y][x] != sq.id {
					found = false
					continue outer
				}
			}
		}

		if found {
			return sq.id
		}
	}

	return 0
}

func renderGrid(x, y int, sq []square) [][]int {
	ret := make([][]int, y)
	for r := range ret {
		ret[r] = make([]int, x)
	}

	for _, s := range sq {
		for y := s.top; y < s.top+s.h; y++ {
			for x := s.left; x < s.left+s.w; x++ {
				if ret[y][x] == 0 {
					ret[y][x] = s.id
				} else {
					ret[y][x] = -1
				}
			}
		}
	}

	return ret
}

func maxDimensions(in []square) (x, y int) {
	for _, sq := range in {
		h := sq.left + sq.w + 1
		if h > x {
			x = h
		}
		v := sq.top + sq.h + 1
		if v > y {
			y = v
		}
	}

	return x, y
}

func parse(in string) (square, error) {
	ret := square{}
	n, err := fmt.Sscanf(in, format, &ret.id, &ret.left, &ret.top, &ret.w, &ret.h)
	if err != nil {
		return ret, formatError{in: in, cause: err}
	}
	if n != 5 {
		return ret, formatError{in: in, cause: fmt.Errorf("read %d items", n)}
	}

	return ret, nil
}
