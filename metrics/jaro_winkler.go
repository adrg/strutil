package metrics

import (
	"strings"
	"unicode/utf8"

	"github.com/adrg/strutil/internal/stringutil"
)

// JaroWinkler represents the Jaro-Winkler metric for measuring the similarity
// between sequences.
//
// For more information see https://en.wikipedia.org/wiki/Jaro-Winkler_distance.
type JaroWinkler struct {
	// CaseSensitive specifies if the string comparison is case sensitive.
	CaseSensitive bool

	// Threshold specifies a Jaro similarity threshold (usually 0.7) at which
	// to add the Winkler prefix bonus. If the similarity of the compared terms
	// is less than the specified threshold, the Jaro similarity is returned
	// without applying any score boost to it.
	Threshold float64
}

// NewJaroWinkler returns a new Jaro-Winkler string metric.
//
// Default options:
//
//	CaseSensitive: true
//	Threshold: 0.7
func NewJaroWinkler() *JaroWinkler {
	return &JaroWinkler{
		CaseSensitive: true,
		Threshold:     0.7,
	}
}

// Compare returns the Jaro-Winkler similarity of a and b. The returned
// similarity is a number between 0 and 1. Larger similarity numbers indicate
// closer matches.
func (m *JaroWinkler) Compare(a, b string) float64 {
	// Lower terms if case insensitive comparison is specified.
	if !m.CaseSensitive {
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}

	// Calculate common prefix.
	lenPrefix := utf8.RuneCountInString(stringutil.CommonPrefix(a, b))
	if lenPrefix > 4 {
		lenPrefix = 4
	}

	// Calculate Jaro similarity.
	jaro := NewJaro()
	jaro.CaseSensitive = m.CaseSensitive

	// If the Jaro similarity value is less than the configured threshold,
	// do not apply the prefix bonus.
	similarity := jaro.Compare(a, b)
	if similarity < m.Threshold {
		return similarity
	}
	return similarity + (0.1 * float64(lenPrefix) * (1.0 - similarity))
}
