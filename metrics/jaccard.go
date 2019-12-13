package metrics

import (
	"strings"

	"github.com/adrg/strutil/internal/util"
)

// Jaccard represents the Jaccard index for measuring the similarity
// between sequences.
//   https://en.wikipedia.org/wiki/Jaccard_index.
type Jaccard struct {
	// CaseSensitive specifies if the string comparison is case sensitive.
	CaseSensitive bool

	// NgramSize represents the size (in characters) of the tokens generated
	// when comparing the input sequences.
	NgramSize int
}

// NewJaccard returns a new Jaccard string metric.
//
// Default options:
//   CaseSensitive: true
//   NGramSize: 2
func NewJaccard() *Jaccard {
	return &Jaccard{
		CaseSensitive: true,
		NgramSize:     2,
	}
}

// Compare returns the Jaccard similarity coefficient of a and b. The
// returned similarity is a number between 0 and 1. Larger similarity numbers
// indicate closer matches.
// An n-gram size of 2 is used if the provided size is less than or equal to 0.
func (m *Jaccard) Compare(a, b string) float64 {
	// Lower terms if case insensitive comparison is specified.
	if !m.CaseSensitive {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}

	// Check if both terms are empty.
	runesA, runesB := []rune(a), []rune(b)
	if len(runesA) == 0 && len(runesB) == 0 {
		return 1
	}

	ngramSize := m.NgramSize
	if ngramSize <= 0 {
		ngramSize = 2
	}

	// Calculate n-gram intersection and union.
	_, intersection, total := util.NgramIntersection(runesA, runesB, ngramSize)
	if total == 0 {
		return 0
	}

	// Return similarity.
	return float64(intersection) / float64(total-intersection)
}
