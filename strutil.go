/*
Package strutil provides string metrics for calculating string similarity as
well as other string utility functions.
*/
package strutil

import (
	"github.com/adrg/strutil/internal/util"
)

// StringMetric represents a metric for measuring the similarity between
// strings. The metrics package implements the following string metrics:
//  - Jaro
//  - Jaro-Winkler
//  - Levenshtein
//  - Smith-Waterman-Gotoh
//  - Sorensen-Dice
//
// For more information see https://godoc.org/github.com/adrg/strutil/metrics.
type StringMetric interface {
	Compare(a, b string) float64
}

// Similarity returns the similarity of a and b, computed using the specified
// string metric. The returned similarity is a number between 0 and 1. Larger
// similarity numbers indicate closer matches.
func Similarity(a, b string, metric StringMetric) float64 {
	return metric.Compare(a, b)
}

// CommonPrefix returns the common prefix of the specified strings. An empty
// string is returned if the parameters have no prefix in common.
func CommonPrefix(a, b string) string {
	return util.CommonPrefix(a, b)
}

// UniqueSlice returns a slice containing the unique items from the specified
// string slice. The items in the output slice are in the order in which they
// occur in the input slice.
func UniqueSlice(items []string) []string {
	return util.UniqueSlice(items)
}

// SliceContains returns true if terms contains q, or false otherwise.
func SliceContains(terms []string, q string) bool {
	return util.SliceContains(terms, q)
}
