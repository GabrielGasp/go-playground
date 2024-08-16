package compareslices

func CompareSlicesUsingMap(a, b []string) []string {
	var diff []string
	m := make(map[string]bool)

	for _, s := range b {
		m[s] = true
	}

	for _, s := range a {
		if _, ok := m[s]; !ok {
			diff = append(diff, s)
		}
	}

	return diff
}

func CompareSlicesUsingNestedLoop(a, b []string) []string {
	var diff []string

OuterLoop:
	for _, s1 := range a {
		for _, s2 := range b {
			if s1 == s2 {
				continue OuterLoop
			}
		}

		diff = append(diff, s1)
	}

	return diff
}
