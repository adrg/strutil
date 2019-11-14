package strutil

import "strings"

// SWGOpts defines options to use for calculating the Smith-Waterman-Gotoh
// similarity.
type SWGOpts struct {
	// CaseSensitive specifies if the string comparison is case sensitive.
	CaseSensitive bool

	// GapPenalty defines a score penalty for character insertions or deletions.
	// For relevant results, the gap penalty should be a non-positive number.
	GapPenalty float64

	// Substitution represents a substitution function which is used to
	// calculate a score for character substitutions.
	Substitution interface {
		// Compare returns the substitution score of characters
		// a[aIdx] and b[bIdx].
		Compare(a []rune, aIdx int, b []rune, bIdx int) float64

		// Returns the maximum value a gap can have.
		Max() float64

		// Returns the minimum value a gap can have.
		Min() float64
	}
}

// SWGMatchMismatch represents a substitution function which can be used with
// the Smith-Waterman-Gotoh algorithm. The function returns the match or
// mismatch value depeding on the equality of the compared characters. The
// match value must be greater than the mismatch value.
type SWGMatchMismatch struct {
	// Match represents the score of equal character substitutions.
	Match float64

	// Mismatch represents the score of unequal character substitutions.
	Mismatch float64
}

// Compare returns the match value if a[aIdx] is equal to b[bIdx] or the
// mismatch value otherwise.
func (m SWGMatchMismatch) Compare(a []rune, aIdx int, b []rune, bIdx int) float64 {
	if a[aIdx] == b[bIdx] {
		return m.Match
	}

	return m.Mismatch
}

// Max returns the match value.
func (m SWGMatchMismatch) Max() float64 {
	return m.Match
}

// Min returns the mismatch value.
func (m SWGMatchMismatch) Min() float64 {
	return m.Mismatch
}

// DefaultSWGOpts contains default options used by the SmithWatermanGotoh
// function if no options are specified.
var DefaultSWGOpts = &SWGOpts{
	CaseSensitive: true,
	GapPenalty:    -0.5,
	Substitution: SWGMatchMismatch{
		Match:    1,
		Mismatch: -2,
	},
}

// SmithWatermanGotoh returns the Smith-Waterman-Gotoh similarity of first and
// second. If the opts parameter is nil, DefaultSWGOpts is used. The returned
// similarity is a number between 0 and 1. Larger similarity numbers indicate
// closer matches.
//
// For more information see https://en.wikipedia.org/wiki/Smith-Waterman_algorithm.
func SmithWatermanGotoh(first, second string, opts *SWGOpts) float64 {
	// Use default options, if none are specified.
	if opts == nil {
		opts = DefaultSWGOpts
	}
	gap := opts.GapPenalty

	// Lower terms if case insensitive comparison is specified.
	if !opts.CaseSensitive {
		first = strings.ToLower(first)
		second = strings.ToLower(second)
	}
	fRunes, sRunes := []rune(first), []rune(second)

	// Check if both terms are empty.
	fLen, sLen := len(fRunes), len(sRunes)
	if fLen == 0 && sLen == 0 {
		return 1
	}

	// Check if one of the terms is empty.
	if fLen == 0 || sLen == 0 {
		return 0
	}

	// Use default substitution, if none is specified.
	subst := opts.Substitution
	if subst == nil {
		subst = DefaultSWGOpts.Substitution
	}

	// Calculate max distance.
	maxDistance := minf(float64(fLen), float64(sLen)) * maxf(subst.Max(), gap)

	// Calculate distance.
	v0 := make([]float64, sLen)
	v1 := make([]float64, sLen)

	distance := maxf(0, gap, subst.Compare(fRunes, 0, sRunes, 0))
	v0[0] = distance

	for i := 1; i < sLen; i++ {
		v0[i] = maxf(0, v0[i-1]+gap, subst.Compare(fRunes, 0, sRunes, i))
		distance = maxf(distance, v0[i])
	}

	for i := 1; i < fLen; i++ {
		v1[0] = maxf(0, v0[0]+gap, subst.Compare(fRunes, i, sRunes, 0))
		distance = maxf(distance, v1[0])

		for j := 1; j < sLen; j++ {
			v1[j] = maxf(0, v0[j]+gap, v1[j-1]+gap, v0[j-1]+subst.Compare(fRunes, i, sRunes, j))
			distance = maxf(distance, v1[j])
		}

		for j := 0; j < sLen; j++ {
			v0[j] = v1[j]
		}
	}

	// Compute similarity.
	return distance / maxDistance
}
