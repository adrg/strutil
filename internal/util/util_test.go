package util_test

import (
	"testing"

	"github.com/adrg/strutil/internal/util"
	"github.com/stretchr/testify/require"
)

func TestMin(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{0, util.Min()},
		{1, util.Min(1)},
		{0, util.Min(0, 1)},
		{1, util.Min(1, 1)},
		{1, util.Min(2, 1)},
		{1, util.Min(1, 2)},
		{0, util.Min(2, 1, 0)},
		{0, util.Min(0, 1, 2)},
	})
}

func TestMax(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{0, util.Max()},
		{1, util.Max(1)},
		{1, util.Max(0, 1)},
		{1, util.Max(1, 1)},
		{2, util.Max(2, 1)},
		{2, util.Max(1, 2)},
		{3, util.Max(2, 1, 3)},
		{3, util.Max(3, 1, 2)},
	})
}

func TestMinf(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{0.0, util.Minf()},
		{1.0, util.Minf(1.0)},
		{0.0, util.Minf(0.0, 1.0)},
		{1.0, util.Minf(1.0, 1.0)},
		{1.0, util.Minf(2.0, 1.0)},
		{1.0, util.Minf(1.0, 2.0)},
		{0.0, util.Minf(2.0, 1.0, 0.0)},
		{0.0, util.Minf(0.0, 1.0, 2.0)},
	})
}

func TestMaxf(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{0.0, util.Maxf()},
		{1.0, util.Maxf(1.0)},
		{1.0, util.Maxf(0.0, 1.0)},
		{1.0, util.Maxf(1.0, 1.0)},
		{2.0, util.Maxf(2.0, 1.1, 1.0)},
		{2.0, util.Maxf(1.1, 1.0, 2.0)},
		{3.0, util.Maxf(2.0, 1.0, 3.0)},
		{3.0, util.Maxf(3.0, 1.0, 2.0)},
	})
}

func TestCommonPrefix(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{"", util.CommonPrefix("", "")},
		{"", util.CommonPrefix("a", "")},
		{"", util.CommonPrefix("", "b")},
		{"", util.CommonPrefix("a", "b")},
		{"a", util.CommonPrefix("ab", "aab")},
		{"a", util.CommonPrefix("aab", "ab")},
		{"aa", util.CommonPrefix("aab", "aaab")},
		{"aa", util.CommonPrefix("aaab", "aab")},
	})
}

func TestUniqueSlice(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{0, len(util.UniqueSlice(nil))},
		{0, len(util.UniqueSlice([]string{}))},
		{[]string{"a"}, util.UniqueSlice([]string{"a"})},
		{[]string{"a", "b"}, util.UniqueSlice([]string{"a", "b"})},
		{[]string{"b", "a"}, util.UniqueSlice([]string{"b", "a"})},
		{[]string{"a"}, util.UniqueSlice([]string{"a", "a"})},
		{[]string{"b", "a"}, util.UniqueSlice([]string{"b", "a", "a"})},
		{[]string{"a", "b"}, util.UniqueSlice([]string{"a", "a", "b"})},
		{[]string{"a", "b"}, util.UniqueSlice([]string{"a", "a", "a", "b"})},
		{[]string{"b", "a"}, util.UniqueSlice([]string{"b", "a", "a", "a"})},
		{[]string{"a", "b"}, util.UniqueSlice([]string{"a", "b", "b", "a"})},
		{[]string{"a", "b"}, util.UniqueSlice([]string{"a", "b", "a", "b"})},
	})
}

func TestSliceContains(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{false, util.SliceContains(nil, "")},
		{false, util.SliceContains(nil, "a")},
		{false, util.SliceContains([]string{}, "")},
		{false, util.SliceContains([]string{}, "a")},
		{true, util.SliceContains([]string{"a", "b"}, "a")},
		{true, util.SliceContains([]string{"b", "a"}, "a")},
		{false, util.SliceContains([]string{"b", "a"}, "c")},
	})
}

func requireEqual(t *testing.T, inputs [][2]interface{}) {
	t.Helper()

	for _, input := range inputs {
		require.Equal(t, input[0], input[1])
	}
}
