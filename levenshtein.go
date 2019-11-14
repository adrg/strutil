package strutil

import (
	"strings"
	"unicode/utf8"
)

// LevOpts defines options to use for calculating the Levenshtein distance.
type LevOpts struct {
	// CaseSensitive specifies if the string comparison is case sensitive.
	CaseSensitive bool

	// InsertCost represents the Levenshtein cost of a character insertion.
	InsertCost int

	// InsertCost represents the Levenshtein cost of a character deletion.
	DeleteCost int

	// InsertCost represents the Levenshtein cost of a character substitution.
	ReplaceCost int
}

// DefaultLevOpts contains default options used by the Levenshtein function
// if no options are specified.
var DefaultLevOpts = &LevOpts{
	CaseSensitive: true,
	InsertCost:    1,
	DeleteCost:    1,
	ReplaceCost:   1,
}

// Levenshtein returns the Levenshtein distance and similarity of first and
// second, based on the specified options. If the opts parameter is nil,
// DefaultLevOpts is used. The returned distance represents the number of
// operations needed in order to transform first into second. The returned
// similarity is a number between 0 and 1. Larger similarity numbers indicate
// closer matches.
//
// For more information see https://en.wikipedia.org/wiki/Levenshtein_distance.
func Levenshtein(first, second string, opts *LevOpts) (int, float64) {
	// Check if both terms are empty.
	fLen, sLen := utf8.RuneCountInString(first), utf8.RuneCountInString(second)
	if fLen == 0 && sLen == 0 {
		return 0, 1
	}

	// Use default options, if none are specified.
	if opts == nil {
		opts = DefaultLevOpts
	}

	// Check if one of the terms is empty.
	if fLen == 0 {
		return opts.InsertCost * sLen, 0
	}
	if sLen == 0 {
		return opts.DeleteCost * fLen, 0
	}

	// Lower terms if case insensitive comparison is specified.
	if !opts.CaseSensitive {
		first = strings.ToLower(first)
		second = strings.ToLower(second)
	}

	// Initialize cost slice.
	prevCol := make([]int, sLen+1)
	for i := 0; i <= sLen; i++ {
		prevCol[i] = i
	}

	// Calculate distance.
	col := make([]int, sLen+1)
	for i := 0; i < fLen; i++ {
		col[0] = i + 1
		for j := 0; j < sLen; j++ {
			delCost := prevCol[j+1] + opts.DeleteCost
			insCost := col[j] + opts.InsertCost

			subCost := prevCol[j]
			if first[i] != second[j] {
				subCost += opts.ReplaceCost
			}

			col[j+1] = min(delCost, insCost, subCost)
		}

		col, prevCol = prevCol, col
	}

	// Return distance and similarity.
	distance := prevCol[sLen]
	return distance, 1 - float64(distance)/float64(max(fLen, sLen))
}
