package runtime

func split2(s string, ch byte) (s1, s2 string) {
	for i := 0; i < len(s); i++ {
		if s[i] == ch {
			return s[:i], s[i+1:]
		}
	}
	return s, ""
}

