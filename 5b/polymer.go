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

func bestReaction(n string) string {
	in := []rune(n)
	best := n
	for i := rune('A'); i <= 'Z'; i++ {
		cnt := 0
		str := make([]rune, len(in), len(in))
		for j := 0; j < len(in); j++ {
			if in[j] != i && in[j] != i+0x20 {
				str[cnt] = in[j]
				cnt++
			}
		}

		result := react(string(str[:cnt]))
		if len(result) < len(best) {
			best = result
		}
	}

	return best
}
