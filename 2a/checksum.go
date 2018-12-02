package aoc

import "fmt"

// Check is a single box ID
type Check struct {
	three bool
	two   bool
}

// InvalidCharacter represents a single character input error
type InvalidCharacter rune

// Error returns the error as a string
func (i InvalidCharacter) Error() string {
	return fmt.Sprintf("%c", i)
}

// Parse takes a string and returns a check for it
func Parse(n string) (Check, error) {
	scratch := make(map[rune]int)

	c := Check{}
	for _, l := range n {
		if l < 'a' || l > 'z' {
			return c, InvalidCharacter(l)
		}

		// just add up the characters
		scratch[l] = scratch[l] + 1
	}

	for _, v := range scratch {
		if v == 3 {
			c.three = true
		}
		if v == 2 {
			c.two = true
		}
	}

	return c, nil
}

// Checksum returns the checksum given a set of ids
func Checksum(checks []Check) int {
	var threes, twos int

	for _, c := range checks {
		if c.three {
			threes++
		}
		if c.two {
			twos++
		}
	}

	return threes * twos
}
