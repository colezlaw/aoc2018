package aoc

// Accumulator is a type that can increment and decrement
// given an input
type Accumulator int64

// BadInputError is thrown when the input is bad
type BadInputError string

func (b BadInputError) Error() string {
	return string(b)
}

// Acc will add or subtract the specified value
func (a *Accumulator) Acc(n string) error {
	r, err := parse(n)
	if err != nil {
		return err
	}

	*a = Accumulator(int64(*a) + int64(r))

	return nil
}

// parse parses the input string. I realize I could do something like
// strconv.Atoi, but thought I'd do it by hand
func parse(n string) (int, error) {
	acc := 0
	b := []byte(n)
	if b[0] != '+' && b[0] != '-' {
		return acc, BadInputError(n)
	}

	if len(b) < 2 {
		return acc, BadInputError(n)
	}

	mul := 1
	if b[0] == '-' {
		mul = -1
	}

	for _, x := range b[1:] {
		if x < 0x30 || x > 0x39 {
			return acc, BadInputError(n)
		}
		acc = acc * 10
		acc = acc + int(x) - 0x30
	}

	return acc * mul, nil
}
