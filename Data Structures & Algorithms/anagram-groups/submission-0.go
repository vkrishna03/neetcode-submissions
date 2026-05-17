import "slices"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	var sol [][]string

	for i := 0; i < len(strs); i++ {
		
		r := []rune(strs[i])
		slices.Sort(r)
		key := string(r)

		if _, ok := m[key]; ok {
			continue
		}

		var val []string
		val = append(val, strs[i])
		
		for j := i + 1; j < len(strs); j++ {
			if checkAnagram(strs[i], strs[j]) {
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

func checkAnagram(str1 string, str2 string) bool {
	m := make(map[rune]int)
	
	for _, char := range str1{
		m[char] += 1
	}

	for _, char := range str2{
		m[char] -=1
	}

	for _, val := range m {
		if val != 0 {
			return false
		}
	}

	return true	
}