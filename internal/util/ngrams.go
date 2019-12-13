package util

// Ngrams returns all the n-grams of the specified size for the provided term.
// The n-grams in the output slice are in the order in which they occur in the
// input term. An n-gram size of 1 is used if the provided size is less than or
// equal to 0.
func Ngrams(runes []rune, ngramSize int) []string {
	// Use an n-gram size of 1 if the provided size is invalid.
	ngramSize = Max(ngramSize, 1)

	// Check if term length is too small.
	lenRunes := len(runes)
	if lenRunes == 0 || lenRunes < ngramSize {
		return nil
	}

	// Generate n-gram slice.
	limit := lenRunes - (ngramSize - 1)
	ngrams := make([]string, limit)

	for i, j := 0, 0; i < limit; i++ {
		ngrams[j] = string(runes[i : i+ngramSize])
		j++
	}

	return ngrams
}

// NgramMap returns a map of all n-grams of the specified size for the provided
// term, along with their frequency. The function also returns the total number
// of n-grams, which is the sum of all the values in the output map.
// An n-gram size of 1 is used if the provided size is less than or equal to 0.
func NgramMap(runes []rune, ngramSize int) (map[string]int, int) {
	// Use an n-gram size of 1 if the provided size is invalid.
	ngramSize = Max(ngramSize, 1)

	// Check if term length is too small.
	lenRunes := len(runes)
	if lenRunes == 0 || lenRunes < ngramSize {
		return map[string]int{}, 0
	}

	// Generate n-gram map.
	limit := lenRunes - (ngramSize - 1)
	ngrams := map[string]int{}
	var ngramCount int

	for i := 0; i < limit; i++ {
		ngram := string(runes[i : i+ngramSize])
		count, _ := ngrams[ngram]
		ngrams[ngram] = count + 1
		ngramCount++
	}

	return ngrams, ngramCount
}

// NgramIntersection returns a map of the n-grams of the specified size found
// in both terms, along with their frequency. The function also returns the
// number of common n-grams (the sum of all the values in the output map) and
// the total number of n-grams (the count of all n-grams in both terms, common
// or not). An n-gram size of 1 is used if the provided size is less than or
// equal to 0.
func NgramIntersection(a, b []rune, ngramSize int) (map[string]int, int, int) {
	// Use an n-gram size of 1 if the provided size is invalid.
	ngramSize = Max(ngramSize, 1)

	// Compute the n-grams of the first term.
	ngrams, total := NgramMap(a, ngramSize)

	// Calculate n-gram intersection with the second term.
	var intersection int
	commonNgrams := map[string]int{}
	limit := len(b) - (ngramSize - 1)

	for i := 0; i < limit; i++ {
		ngram := string(b[i : i+ngramSize])
		total++

		if count, ok := ngrams[ngram]; ok && count > 0 {
			// Decrease frequency of n-gram found in the first term each time
			// a successful match is found.
			intersection++
			ngrams[ngram] = count - 1

			// Update common n-grams map with the matched n-gram and its
			// frequency.
			count, _ = commonNgrams[ngram]
			commonNgrams[ngram] = count + 1
		}
	}

	return commonNgrams, intersection, total
}
