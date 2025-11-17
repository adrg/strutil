package metrics

import (
	"strings"
	"unicode/utf8"

	"github.com/dorzzz/strutil/internal/mathutil"
)

// Jaro represents the Jaro metric for measuring the similarity
// between sequences.
//   For more information see https://en.wikipedia.org/wiki/Jaro-Winkler_distance.
type Jaro struct {
	// CaseSensitive specifies if the string comparison is case sensitive.
	CaseSensitive     bool
	UseStandardWindow int
}

// NewJaro returns a new Jaro string metric.
//
// Default options:
//   CaseSensitive: true
//   UseStandardWindow: 0 (uses original strutil algorithm)
func NewJaro() *Jaro {
	return &Jaro{
		CaseSensitive:     true,
		UseStandardWindow: 0,
	}
}

// Compare returns the Jaro similarity of a and b. The returned similarity is
// a number between 0 and 1. Larger similarity numbers indicate closer matches.
func (m *Jaro) Compare(a, b string) float64 {
	// Use rune counts (UTF-8 code points) for lengths.
	lenA, lenB := utf8.RuneCountInString(a), utf8.RuneCountInString(b)

	// Check if both terms are empty.
	if lenA == 0 && lenB == 0 {
		return 1.0
	}

	// Check if one of the terms is empty.
	if lenA == 0 || lenB == 0 {
		return 0.0
	}

	// Lower terms if case insensitive comparison is specified.
	if !m.CaseSensitive {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}

	// Choose algorithm based on UseStandardWindow
	if m.UseStandardWindow == 1 {
		// Apache Commons implementation
		if a == b {
			return 1.0
		}

		ra := []rune(a)
		rb := []rune(b)

		matches, halfTranspositions := jaroMatches(ra, rb, m.UseStandardWindow)
		if matches == 0 {
			return 0.0
		}

		mFloat := float64(matches)
		return (mFloat/float64(lenA) +
			mFloat/float64(lenB) +
			(mFloat-float64(halfTranspositions)/2.0)/mFloat) / 3.0
	}

	// Original strutil implementation (default)
	halfLen := mathutil.Max(0, mathutil.Max(lenA, lenB)/2)
	mrA := matchingRunes(a, b, halfLen)
	mrB := matchingRunes(b, a, halfLen)

	fmLen, smLen := len(mrA), len(mrB)
	if fmLen == 0 || smLen == 0 {
		return 0.0
	}

	return (float64(fmLen)/float64(lenA) +
		float64(smLen)/float64(lenB) +
		float64(fmLen-transpositions(mrA, mrB)/2)/float64(fmLen)) / 3.0
}

// jaroMatches mirrors Apache's JaroWinklerSimilarity.matches(...) logic,
// but operating on rune slices instead of Java chars.
func jaroMatches(first, second []rune, useStandardWindow int) (matches int, halfTranspositions int) {
	var maxRunes, minRunes []rune
	if len(first) > len(second) {
		maxRunes = first
		minRunes = second
	} else {
		maxRunes = second
		minRunes = first
	}

	// range = Math.max(max.length()/2 - 1, 0)
	rng := maxInt(len(maxRunes)/2-useStandardWindow, 0)

	matchIndexes := make([]int, len(minRunes))
	for i := range matchIndexes {
		matchIndexes[i] = -1
	}
	matchFlags := make([]bool, len(maxRunes))

	// Find matches
	for mi, c1 := range minRunes {
		start := maxInt(mi-rng, 0)
		end := minInt(mi+rng+1, len(maxRunes))
		for xi := start; xi < end; xi++ {
			if !matchFlags[xi] && c1 == maxRunes[xi] {
				matchIndexes[mi] = xi
				matchFlags[xi] = true
				matches++
				break
			}
		}
	}

	// Build the two matched sequences ms1, ms2
	ms1 := make([]rune, matches)
	ms2 := make([]rune, matches)

	si := 0
	for i := 0; i < len(minRunes); i++ {
		if matchIndexes[i] != -1 {
			ms1[si] = minRunes[i]
			si++
		}
	}

	si = 0
	for i := 0; i < len(maxRunes); i++ {
		if matchFlags[i] {
			ms2[si] = maxRunes[i]
			si++
		}
	}

	// Count half-transpositions
	for i := 0; i < len(ms1); i++ {
		if ms1[i] != ms2[i] {
			halfTranspositions++
		}
	}

	return matches, halfTranspositions
}

// matchingRunes returns the matching runes between a and b within the specified limit.
// This is the original strutil implementation.
func matchingRunes(a, b string, limit int) []rune {
	var (
		runesA      = []rune(a)
		runesB      = []rune(b)
		runesCommon = []rune{}
		lenB        = len(runesB)
	)

	for i, r := range runesA {
		end := mathutil.Min(i+limit+1, lenB)
		for j := mathutil.Max(0, i-limit); j < end; j++ {
			if r == runesB[j] && runesB[j] != -1 {
				runesCommon = append(runesCommon, runesB[j])
				runesB[j] = -1
				break
			}
		}
	}

	return runesCommon
}

// transpositions counts the number of transpositions between two rune slices.
// This is the original strutil implementation.
func transpositions(a, b []rune) int {
	var count int

	minLen := mathutil.Min(len(a), len(b))
	for i := 0; i < minLen; i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count
}

// local int helpers
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
