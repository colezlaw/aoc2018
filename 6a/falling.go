package aoc

import (
	"bufio"
	"fmt"
	"io"
)

type dest struct {
	label rune
	x     int
	y     int
}

type marker struct {
	label    rune
	distance int
}

func parse(s string) dest {
	var x, y int
	fmt.Sscanf(s, "%d, %d", &x, &y)
	return dest{x: x, y: y}
}

func load(r io.Reader) (dests map[rune]dest, minx, miny, maxx, maxy int) {
	alphabet := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	var i int
	dests = make(map[rune]dest)
	minx = 1000000
	miny = 1000000
	maxx = -1000000
	maxy = -1000000

	s := bufio.NewScanner(r)
	for s.Scan() {
		d := parse(s.Text())
		d.label = alphabet[i]
		i++
		dests[alphabet[i]] = d

		if d.x < minx {
			minx = d.x
		}
		if d.y < miny {
			miny = d.y
		}
		if d.x > maxx {
			maxx = d.x
		}
		if d.y > maxy {
			maxy = d.y
		}
	}

	return
}

func renderGrid(input map[rune]dest, minx, miny, maxx, maxy int) [][]marker {
	result := make([][]marker, maxy-miny+1, maxy-miny+1)
	for r := range result {
		result[r] = make([]marker, maxx-minx+1, maxx-minx+1)
		for c := range result[r] {
			result[r][c] = marker{label: '-', distance: 1000000}
		}
	}

	for k, v := range input {
		for r := 0; r < len(result); r++ {
			for c := 0; c < len(result[r]); c++ {
				dx := (v.x - minx) - c
				if dx < 0 {
					dx *= -1
				}

				dy := (v.y - miny) - r
				if dy < 0 {
					dy *= -1
				}

				dist := dx + dy
				if dist < result[r][c].distance {
					result[r][c].distance = dist
					result[r][c].label = k
				} else if dist == result[r][c].distance {
					result[r][c].label = '.'
				}
			}
		}
	}

	return result
}

// given a map and the targets, find the target with the largest
// gap around it
func areaSizes(grid [][]marker, targets map[rune]dest) (rune, int) {
	sizes := make(map[rune]int)

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c].label != '.' {
				sizes[grid[r][c].label]++
			}
		}
	}

	// Since the grid is infinite size, we have to remove
	// all the regions that are on a boundary
	for c := range grid[0] {
		delete(sizes, grid[0][c].label)
		delete(sizes, grid[len(grid)-1][c].label)
	}
	for r := range grid {
		delete(sizes, grid[r][0].label)
		delete(sizes, grid[r][len(grid[r])-1].label)
	}

	// Now find the biggest block remaining
	var maxlabel rune
	var maxsize int
	for k, v := range sizes {
		if v > maxsize {
			maxsize = v
			maxlabel = k
		}
	}

	return maxlabel, maxsize
}
