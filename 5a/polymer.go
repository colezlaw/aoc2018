package aoc

func react(in string) string {
	d := []rune(in)
	found := true

loop:
	for found {
		found = false
		for i := 0; i < len(d)-1; i++ {
			if d[i]-d[i+1] == 0x20 || d[i+1]-d[i] == 0x20 {
				found = true
				d = append(d[:i], d[i+2:]...)
				continue loop
			}
		}
	}

	return string(d)
}
