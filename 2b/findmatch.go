package aoc

func findMatch(ids []string) string {
	for i := 0; i < len(ids)-1; i++ {
		for j := i + 1; j < len(ids); j++ {
			ret := make([]byte, 0)
			for o := 0; o < len(ids[i]); o++ {
				if ids[i][o] == ids[j][o] {
					ret = append(ret, ids[i][o])
				}
			}

			if len(ret) == len(ids[j])-1 {
				return string(ret)
			}
		}
	}

	return ""
}
