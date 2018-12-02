package aoc

// Accumulator is a type that can increment and decrement
// given an input
type Accumulator struct {
	freq  int64
	hist  map[int64]struct{}
	found bool
}

// New returns a new Accumulator
func New() *Accumulator {
	a := new(Accumulator)
	a.hist = make(map[int64]struct{})

	return a
}

// BadInputError is thrown when the input is bad
type BadInputError string

func (b BadInputError) Error() string {
	return string(b)
}

// Acc will add or subtract the specified value
func (a *Accumulator) Acc(n string) error {
	// We've already found the frequency, bail
	if a.found {
		return nil
	}
	a.hist[a.freq] = struct{}{}

	r, err := parse(n)
	if err != nil {
		return err
	}

	a.freq += int64(r)
	if _, ok := a.hist[a.freq]; ok {
		a.found = true
	}

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
