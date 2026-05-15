func isAnagram(s string, t string) bool {
	m:= make(map[rune]int)

	for _, val := range s{
		m[val] += 1
	}

	for _, val := range t{
		m[val] -= 1
	}

	for _, val := range m{
		if val != 0 {
			return false
		}
	}

	return true
}
