/*
Package strutil provides string metrics for calculating string similarity as
well as other string utility functions.
*/
package strutil

import (
	"unicode/utf8"
)

// CommonPrefix returns the common prefix of the specified strings. An empty
// string is returned if the parameters have no prefix in common.
func CommonPrefix(first, second string) string {
	if utf8.RuneCountInString(first) > utf8.RuneCountInString(second) {
		first, second = second, first
	}

	var commonLen int
	sRunes := []rune(second)
	for i, r := range first {
		if r != sRunes[i] {
			break
		}

		commonLen++
	}

	return string(sRunes[0:commonLen])
}

// Unique returns a slice containing the unique items from the specified
// string slice.
func Unique(items []string) []string {
	uniq := make([]string, len(items))

	index := 0
	catalog := map[string]struct{}{}
	for _, item := range items {
		if _, ok := catalog[item]; ok {
			continue
		}

		catalog[item] = struct{}{}
		uniq[index] = item
		index++
	}

	return uniq[0:index]
}
