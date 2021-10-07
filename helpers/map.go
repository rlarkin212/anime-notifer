package helpers

func Contains(s map[string]struct{}, searchTerm string) bool {
	ret := false

	if _, ok := s[searchTerm]; ok {
		ret = true
	}

	return ret
}
