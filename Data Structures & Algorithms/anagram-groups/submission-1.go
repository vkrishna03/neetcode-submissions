import "slices"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	var sol [][]string

	for i := 0; i < len(strs); i++ {
		key := sortStr(strs[i])

		if _, ok := m[key]; ok {
			continue
		}

		var val []string
		val = append(val, strs[i])
		
		for j := i + 1; j < len(strs); j++ {
			if key == sortStr(strs[j]) {
				val = append(val, strs[j])
			}
		}
		m[key] = val
	}

	for _, val := range m{
		sol = append(sol, val)
	}

	return sol
}

func sortStr(str string) string {
	r := []rune(str)
	slices.Sort(r)
	return string(r)	
}