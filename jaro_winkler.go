package strutil

import (
	"unicode/utf8"
)

// Jaro returns the Jaro similarity of first and second. The returned
// similarity is a number in between 0 and 1. Larger similarity numbers
// indicate closer matches.
//
// For more information see https://en.wikipedia.org/wiki/Jaro-Winkler_distance.
func Jaro(first, second string) float64 {
	// Check if both terms are empty.
	fLen, sLen := utf8.RuneCountInString(first), utf8.RuneCountInString(second)
	if fLen == 0 && sLen == 0 {
		return 1
	}

	// Check if one of the terms is empty.
	if fLen == 0 || sLen == 0 {
		return 0
	}

	// Get matching runes.
	halfLen := max(fLen, sLen)/2 - 1
	fm := matchingRunes(first, second, halfLen)
	sm := matchingRunes(second, first, halfLen)

	fmLen, smLen := len(fm), len(sm)
	if fmLen == 0 || smLen == 0 {
		return 0.0
	}

	// Return similarity.
	return (float64(fmLen)/float64(fLen) +
		float64(smLen)/(float64(sLen)) +
		float64(fmLen-transpositions(fm, sm)/2)/float64(fmLen)) / 3.0
}

// JaroWinkler returns the Jaro-Winkler similarity of first and second. The
// returned similarity is a number between 0 and 1. Larger similarity numbers
// indicate closer matches.
//
// For more information see https://en.wikipedia.org/wiki/Jaro-Winkler_distance.
func JaroWinkler(first, second string) float64 {
	// Calculate common prefix.
	lenPrefix := utf8.RuneCountInString(CommonPrefix(first, second))
	if lenPrefix > 4 {
		lenPrefix = 4
	}

	// Return similarity.
	similarity := Jaro(first, second)
	return similarity + (0.1 * float64(lenPrefix) * (1.0 - similarity))
}

func matchingRunes(first, second string, limit int) []rune {
	common := []rune{}
	sRunes := []rune(second)

	for i, r := range first {
		end := min(i+limit, len(sRunes))
		for j := max(0, i-limit); j < end; j++ {
			if r == sRunes[j] {
				common = append(common, sRunes[j])
				sRunes[j] = 0
				break
			}
		}
	}

	return common
}

func transpositions(first, second []rune) int {
	var count int

	minLen := min(len(first), len(second))
	for i := 0; i < minLen; i++ {
		if first[i] != second[i] {
			count++
		}
	}

	return count
}
