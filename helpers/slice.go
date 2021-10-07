package helpers

func SliceToStrMap(s []string) map[string]struct{} {
	m := make(map[string]struct{})

	for _, i := range s {
		m[i] = struct{}{}
	}

	return m
}
