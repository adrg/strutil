package metrics

import (
	"strings"

	"github.com/adrg/strutil/internal/mathutil"
)

// Levenshtein represents the Levenshtein metric for measuring the similarity
// between sequences.
//   For more information see https://en.wikipedia.org/wiki/Levenshtein_distance.
type Levenshtein struct {
	// CaseSensitive specifies if the string comparison is case sensitive.
	CaseSensitive bool

	// InsertCost represents the Levenshtein cost of a character insertion.
	InsertCost int

	// InsertCost represents the Levenshtein cost of a character deletion.
	DeleteCost int

	// InsertCost represents the Levenshtein cost of a character substitution.
	ReplaceCost int
}

// NewLevenshtein returns a new Levenshtein string metric.
//
// Default options:
//   CaseSensitive: true
//   InsertCost: 1
//   DeleteCost: 1
//   ReplaceCost: 1
func NewLevenshtein() *Levenshtein {
	return &Levenshtein{
		CaseSensitive: true,
		InsertCost:    1,
		DeleteCost:    1,
		ReplaceCost:   1,
	}
}

// Compare returns the Levenshtein similarity of a and b. The returned
// similarity is a number between 0 and 1. Larger similarity numbers indicate
// closer matches.
func (m *Levenshtein) Compare(a, b string) float64 {
	distance, maxLen := m.distance(a, b)
	return 1 - float64(distance)/float64(maxLen)
}

// Distance returns the Levenshtein distance between a and b. Lower distances
// indicate closer matches. A distance of 0 means the strings are identical.
func (m *Levenshtein) Distance(a, b string) int {
	distance, _ := m.distance(a, b)
	return distance
}

func (m *Levenshtein) distance(a, b string) (int, int) {
	// Lower terms if case insensitive comparison is specified.
	if !m.CaseSensitive {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}
	runesA, runesB := []rune(a), []rune(b)

	// Check if both terms are empty.
	lenA, lenB := len(runesA), len(runesB)
	if lenA == 0 && lenB == 0 {
		return 0, 0
	}

	// Check if one of the terms is empty.
	maxLen := mathutil.Max(lenA, lenB)
	if lenA == 0 {
		return m.InsertCost * lenB, maxLen
	}
	if lenB == 0 {
		return m.DeleteCost * lenA, maxLen
	}

	// Initialize cost slice.
	prevCol := make([]int, lenB+1)
	for i := 0; i <= lenB; i++ {
		prevCol[i] = i
	}

	// Calculate distance.
	col := make([]int, lenB+1)
	for i := 0; i < lenA; i++ {
		col[0] = i + 1
		for j := 0; j < lenB; j++ {
			delCost := prevCol[j+1] + m.DeleteCost
			insCost := col[j] + m.InsertCost

			subCost := prevCol[j]
			if runesA[i] != runesB[j] {
				subCost += m.ReplaceCost
			}

			col[j+1] = mathutil.Min(delCost, insCost, subCost)
		}

		col, prevCol = prevCol, col
	}

	return prevCol[lenB], maxLen
}
