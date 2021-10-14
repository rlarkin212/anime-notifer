package helpers

func Contains(s map[string]struct{}, searchTerm string) bool {
	if _, ok := s[searchTerm]; ok {
		return true
	}

	return false
}
