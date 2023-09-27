package stringutil

// CommonPrefix returns the common prefix of the specified strings. An empty
// string is returned if the parameters have no prefix in common.
func CommonPrefix(first, second string) string {
	fRunes, sRunes := []rune(first), []rune(second)
	if len(fRunes) > len(sRunes) {
		fRunes, sRunes = sRunes, fRunes
	}

	var commonLen int
	for i, r := range fRunes {
		if r != sRunes[i] {
			break
		}

		commonLen++
	}

	return string(sRunes[0:commonLen])
}

// UniqueSlice returns a slice containing the unique items from the specified
// string slice. The items in the output slice are in the order in which they
// occur in the input slice.
func UniqueSlice(items []string) []string {
	var uniq []string
	registry := map[string]struct{}{}

	for _, item := range items {
		if _, ok := registry[item]; ok {
			continue
		}

		registry[item] = struct{}{}
		uniq = append(uniq, item)
	}

	return uniq
}

// SliceContains returns true if terms contains q, or false otherwise.
func SliceContains(terms []string, q string) bool {
	for _, term := range terms {
		if q == term {
			return true
		}
	}

	return false
}
