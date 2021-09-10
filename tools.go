package SealRPC

func CamelCase(s string) string {
	var cs []rune
	var lowLine bool
	for i, c := range s {
		if (i == 0) && (c >= 97 && c <= 122) {
			cs = append(cs, c-32)
		} else if lowLine {
			if c >= 97 && c <= 122 {
				c = c - 32
			}
			if c == 32 || c == 9 || c == 10 || c == 45 || c == 95 {
				continue
			}
			cs = append(cs, c)
			lowLine = false
		} else if c == 32 || c == 9 || c == 10 || c == 45 || c == 95 {
			lowLine = true
		} else {
			cs = append(cs, c)
		}
	}
	return string(cs)
}
