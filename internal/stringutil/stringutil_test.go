package stringutil_test

import (
	"testing"

	"github.com/adrg/strutil/internal/stringutil"
	"github.com/stretchr/testify/require"
)

func TestCommonPrefix(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{"", stringutil.CommonPrefix("", "")},
		{"", stringutil.CommonPrefix("a", "")},
		{"", stringutil.CommonPrefix("", "b")},
		{"", stringutil.CommonPrefix("a", "b")},
		{"a", stringutil.CommonPrefix("ab", "aab")},
		{"a", stringutil.CommonPrefix("aab", "ab")},
		{"aa", stringutil.CommonPrefix("aab", "aaab")},
		{"aa", stringutil.CommonPrefix("aaab", "aab")},
	})
}

func TestUniqueSlice(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{0, len(stringutil.UniqueSlice(nil))},
		{0, len(stringutil.UniqueSlice([]string{}))},
		{[]string{"a"}, stringutil.UniqueSlice([]string{"a"})},
		{[]string{"a", "b"}, stringutil.UniqueSlice([]string{"a", "b"})},
		{[]string{"b", "a"}, stringutil.UniqueSlice([]string{"b", "a"})},
		{[]string{"a"}, stringutil.UniqueSlice([]string{"a", "a"})},
		{[]string{"b", "a"}, stringutil.UniqueSlice([]string{"b", "a", "a"})},
		{[]string{"a", "b"}, stringutil.UniqueSlice([]string{"a", "a", "b"})},
		{[]string{"a", "b"}, stringutil.UniqueSlice([]string{"a", "a", "a", "b"})},
		{[]string{"b", "a"}, stringutil.UniqueSlice([]string{"b", "a", "a", "a"})},
		{[]string{"a", "b"}, stringutil.UniqueSlice([]string{"a", "b", "b", "a"})},
		{[]string{"a", "b"}, stringutil.UniqueSlice([]string{"a", "b", "a", "b"})},
	})
}

func TestSliceContains(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{false, stringutil.SliceContains(nil, "")},
		{false, stringutil.SliceContains(nil, "a")},
		{false, stringutil.SliceContains([]string{}, "")},
		{false, stringutil.SliceContains([]string{}, "a")},
		{true, stringutil.SliceContains([]string{"a", "b"}, "a")},
		{true, stringutil.SliceContains([]string{"b", "a"}, "a")},
		{false, stringutil.SliceContains([]string{"b", "a"}, "c")},
	})
}

func requireEqual(t *testing.T, inputs [][2]interface{}) {
	t.Helper()

	for _, input := range inputs {
		require.Equal(t, input[0], input[1])
	}
}
