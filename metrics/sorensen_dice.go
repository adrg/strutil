package metrics

import (
	"strings"
)

// SorensenDice represents the Sorensen-Dice metric for measuring the
// similarity between sequences.
//   For more information see https://en.wikipedia.org/wiki/Sorensen–Dice_coefficient.
type SorensenDice struct {
	// CaseSensitive specifies if the string comparison is case sensitive.
	CaseSensitive bool

	// NgramSize represents the size (in characters) of the tokens generated
	// when comparing the input sequences.
	NgramSize int
}

// NewSorensenDice returns a new Sorensen-Dice string metric.
//
// Default options:
//   CaseSensitive: true
//   NGramSize: 2
func NewSorensenDice() *SorensenDice {
	return &SorensenDice{
		CaseSensitive: true,
		NgramSize:     2,
	}
}

// Compare returns the Sorensen-Dice similarity coefficient of a and b. The
// returned similarity is a number between 0 and 1. Larger similarity numbers
// indicate closer matches.
func (m *SorensenDice) Compare(a, b string) float64 {
	// Lower terms if case insensitive comparison is specified.
	if !m.CaseSensitive {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}

	runesA, runesB := []rune(a), []rune(b)
	lenA, lenB := len(runesA), len(runesB)

	// Check if both terms are empty.
	if lenA == 0 && lenB == 0 {
		return 1
	}

	// Check if one of the terms is shorter than the size of a n-gram.
	ngramSize := m.NgramSize
	if lenA < ngramSize || lenB < ngramSize {
		return 0
	}

	// Compute n-grams of the first term.
	ngrams := map[string]int{}
	var ngramCount int

	for i := 0; i < lenA-(ngramSize-1); i++ {
		ngram := string(runesA[i : i+ngramSize])
		count, _ := ngrams[ngram]
		ngrams[ngram] = count + 1
		ngramCount++
	}

	// Calculate n-gram intersection with the second term.
	var intersection int

	for i := 0; i < lenB-(ngramSize-1); i++ {
		ngram := string(runesB[i : i+ngramSize])
		ngramCount++

		if count, ok := ngrams[ngram]; ok && count > 0 {
			intersection++
			ngrams[ngram] = count - 1
		}
	}

	// Return similarity.
	return 2 * float64(intersection) / float64(ngramCount)
}
