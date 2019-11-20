package util

import "unicode/utf8"

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

// Unique returns a slice containing the unique items from the specified string
// slice. The items in the output slice are in the order in which they occur in
// the input slice.
func Unique(items []string) []string {
	var uniq []string
	registry := map[string]struct{}{}

	for _, item := range items {
		if _, ok := registry[item]; ok {
			continue
		}

		registry[item] = struct{}{}
		uniq = append(uniq, item)
	}

	return uniq
}

// Min returns the value of the smallest argument.
func Min(args ...int) int {
	if len(args) == 0 {
		return 0
	}
	if len(args) == 1 {
		return args[0]
	}

	min := args[0]
	for _, arg := range args[1:] {
		if min > arg {
			min = arg
		}
	}

	return min
}

// Max returns the value of the largest argument.
func Max(args ...int) int {
	if len(args) == 0 {
		return 0
	}
	if len(args) == 1 {
		return args[0]
	}

	max := args[0]
	for _, arg := range args[1:] {
		if max < arg {
			max = arg
		}
	}

	return max
}

// Minf returns the value of the smallest argument.
func Minf(args ...float64) float64 {
	if len(args) == 0 {
		return 0
	}
	if len(args) == 1 {
		return args[0]
	}

	min := args[0]
	for _, arg := range args[1:] {
		if min > arg {
			min = arg
		}
	}

	return min
}

// Maxf returns the value of the largest argument.
func Maxf(args ...float64) float64 {
	if len(args) == 0 {
		return 0
	}
	if len(args) == 1 {
		return args[0]
	}

	max := args[0]
	for _, arg := range args[1:] {
		if max < arg {
			max = arg
		}
	}

	return max
}
