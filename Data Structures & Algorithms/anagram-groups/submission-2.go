import "slices"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	var sol [][]string
	var sortedStr []string

	for _, str := range strs {
		sortedStr = append(sortedStr, sortStr(str))
	}

	for i := 0; i < len(strs); i++ {
		key := sortedStr[i]

		if _, ok := m[key]; ok {
			continue
		}

		var val []string
		val = append(val, strs[i])
		
		for j := i + 1; j < len(strs); j++ {
			if key == sortedStr[j] {
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