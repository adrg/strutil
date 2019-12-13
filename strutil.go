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
//  - Jaccard
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

// Ngrams returns all the n-grams of the specified size for the provided term.
// The n-grams in the output slice are in the order in which they occur in the
// input term. An n-gram size of 1 is used if the provided size is less than or
// equal to 0.
func Ngrams(term string, ngramSize int) []string {
	return util.Ngrams([]rune(term), ngramSize)
}

// NgramMap returns a map of all n-grams of the specified size for the provided
// term, along with their frequency. The function also returns the total number
// of n-grams, which is the sum of all the values in the output map.
// An n-gram size of 1 is used if the provided size is less than or equal to 0.
func NgramMap(term string, ngramSize int) (map[string]int, int) {
	return util.NgramMap([]rune(term), ngramSize)
}

// NgramIntersection returns a map of the n-grams of the specified size found
// in both terms, along with their frequency. The function also returns the
// number of common n-grams (the sum of all the values in the output map) and
// the total number of n-grams (the count of all n-grams in both terms, common
// or not). An n-gram size of 1 is used if the provided size is less than or
// equal to 0.
func NgramIntersection(a, b string, ngramSize int) (map[string]int, int, int) {
	return util.NgramIntersection([]rune(a), []rune(b), ngramSize)
}
