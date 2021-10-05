package helpers

import "sort"

func Contains(s []string, searchTerm string) bool {
	i := sort.SearchStrings(s, searchTerm)
	ret := i < len(s) && s[i] == searchTerm

	return ret
}
