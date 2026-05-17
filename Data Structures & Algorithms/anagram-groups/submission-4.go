func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	var sol [][]string

	for i := 0; i < len(strs); i++ {
		key := toCountVector(strs[i])
		m[key] = append(m[key], strs[i]) 
	}

	for _, val := range m{
		sol = append(sol, val)
	}

	return sol
}

func toCountVector(str string) [26]int {
	var v [26]int
	for _, char := range str {
		idx := char - 'a'
		v[idx] += 1
	}
	return v	
}