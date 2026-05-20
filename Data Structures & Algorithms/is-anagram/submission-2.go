import "slices"

func isAnagram(s string, t string) bool {

	a,b := []rune(s), []rune(t)

	slices.Sort(a)
	slices.Sort(b)

	return string(a) == string(b)
}
