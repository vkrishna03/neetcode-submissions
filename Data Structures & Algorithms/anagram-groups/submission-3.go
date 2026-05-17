import "slices"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	var sol [][]string

	for i := 0; i < len(strs); i++ {
		key := sortStr(strs[i])
		m[key] = append(m[key], strs[i]) 
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