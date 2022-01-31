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

func TestNgrams(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{0, len(util.Ngrams(nil, -1))},
		{0, len(util.Ngrams(nil, 0))},
		{0, len(util.Ngrams(nil, 1))},
		{0, len(util.Ngrams([]rune{}, -1))},
		{0, len(util.Ngrams([]rune{}, 0))},
		{0, len(util.Ngrams([]rune{}, 1))},
		{
			[]string{"a", "b", "c", "d", "e", "f"},
			util.Ngrams([]rune("abcdef"), -1),
		},
		{
			[]string{"a", "b", "c", "d", "e", "f"},
			util.Ngrams([]rune("abcdef"), 0),
		},
		{
			[]string{"a", "b", "c", "d", "e", "f"},
			util.Ngrams([]rune("abcdef"), 1),
		},
		{
			[]string{"ab", "bc", "cd", "de", "ef"},
			util.Ngrams([]rune("abcdef"), 2),
		},
		{
			[]string{"abc", "bcd", "cde", "def"},
			util.Ngrams([]rune("abcdef"), 3),
		},
		{
			[]string{"abcd", "bcde", "cdef"},
			util.Ngrams([]rune("abcdef"), 4),
		},
		{
			[]string{"abcde", "bcdef"},
			util.Ngrams([]rune("abcdef"), 5),
		},
		{
			[]string{"abcdef"},
			util.Ngrams([]rune("abcdef"), 6),
		},
		{
			0,
			len(util.Ngrams([]rune("abcdef"), 7)),
		},
	})
}

func TestNgramMap(t *testing.T) {
	inputs := []*struct {
		term     []rune
		size     int
		expMap   map[string]int
		expTotal int
	}{
		{
			term:   nil,
			size:   -1,
			expMap: map[string]int{},
		},
		{
			term:   nil,
			expMap: map[string]int{},
		},
		{
			term:   nil,
			size:   1,
			expMap: map[string]int{},
		},
		{
			term:   []rune{},
			size:   -1,
			expMap: map[string]int{},
		},
		{
			term:   []rune{},
			expMap: map[string]int{},
		},
		{
			term:   []rune{},
			size:   1,
			expMap: map[string]int{},
		},
		{
			term:     []rune("abbabb"),
			size:     -1,
			expMap:   map[string]int{"a": 2, "b": 4},
			expTotal: 6,
		},
		{
			term:     []rune("abbabb"),
			expMap:   map[string]int{"a": 2, "b": 4},
			expTotal: 6,
		},
		{
			term:     []rune("abbabb"),
			size:     1,
			expMap:   map[string]int{"a": 2, "b": 4},
			expTotal: 6,
		},
		{
			term:     []rune("abbabb"),
			size:     2,
			expMap:   map[string]int{"ab": 2, "bb": 2, "ba": 1},
			expTotal: 5,
		},
		{
			term:     []rune("abbabb"),
			size:     3,
			expMap:   map[string]int{"abb": 2, "bba": 1, "bab": 1},
			expTotal: 4,
		},
		{
			term:     []rune("abbabb"),
			size:     4,
			expMap:   map[string]int{"abba": 1, "bbab": 1, "babb": 1},
			expTotal: 3,
		},
		{
			term:     []rune("abbabb"),
			size:     5,
			expMap:   map[string]int{"abbab": 1, "bbabb": 1},
			expTotal: 2,
		},
		{
			term:     []rune("abbabb"),
			size:     6,
			expMap:   map[string]int{"abbabb": 1},
			expTotal: 1,
		},
		{
			term:     []rune("abbabb"),
			size:     7,
			expMap:   map[string]int{},
			expTotal: 0,
		},
	}

	for _, input := range inputs {
		actMap, actTotal := util.NgramMap(input.term, input.size)
		require.Equal(t, input.expMap, actMap)
		require.Equal(t, input.expTotal, actTotal)
	}
}

func TestNgramIntersection(t *testing.T) {
	inputs := []*struct {
		a    []rune
		b    []rune
		size int

		expMap    map[string]int
		expTotal  int
		expTotalA int
		expTotalB int
	}{
		{
			size:   1,
			expMap: map[string]int{},
		},
		{
			a:      []rune{},
			size:   1,
			expMap: map[string]int{},
		},
		{
			b:      []rune{},
			size:   1,
			expMap: map[string]int{},
		},
		{
			a:      []rune{},
			b:      []rune{},
			size:   1,
			expMap: map[string]int{},
		},
		{
			a:         []rune("ababbaa"),
			b:         []rune("aabbaa"),
			size:      -1,
			expMap:    map[string]int{"a": 4, "b": 2},
			expTotal:  6,
			expTotalA: 7,
			expTotalB: 6,
		},
		{
			a:         []rune("aabbaa"),
			b:         []rune("ababbaa"),
			expMap:    map[string]int{"a": 4, "b": 2},
			expTotal:  6,
			expTotalA: 6,
			expTotalB: 7,
		},
		{
			a:         []rune("ababbaa"),
			b:         []rune("aabbaa"),
			size:      1,
			expMap:    map[string]int{"a": 4, "b": 2},
			expTotal:  6,
			expTotalA: 7,
			expTotalB: 6,
		},
		{
			a:         []rune("aabbaa"),
			b:         []rune("ababbaa"),
			size:      2,
			expMap:    map[string]int{"aa": 1, "ab": 1, "ba": 1, "bb": 1},
			expTotal:  4,
			expTotalA: 5,
			expTotalB: 6,
		},
		{
			a:         []rune("ababbaa"),
			b:         []rune("aabbaa"),
			size:      3,
			expMap:    map[string]int{"abb": 1, "bba": 1, "baa": 1},
			expTotal:  3,
			expTotalA: 5,
			expTotalB: 4,
		},
		{
			a:         []rune("aabbaa"),
			b:         []rune("ababbaa"),
			size:      4,
			expMap:    map[string]int{"abba": 1, "bbaa": 1},
			expTotal:  2,
			expTotalA: 3,
			expTotalB: 4,
		},
		{
			a:         []rune("ababbaa"),
			b:         []rune("aabbaa"),
			size:      5,
			expMap:    map[string]int{"abbaa": 1},
			expTotal:  1,
			expTotalA: 3,
			expTotalB: 2,
		},
		{
			a:         []rune("aabbaa"),
			b:         []rune("ababbaa"),
			size:      6,
			expMap:    map[string]int{},
			expTotalA: 1,
			expTotalB: 2,
		},
		{
			a:         []rune("ababbaa"),
			b:         []rune("aabbaa"),
			size:      7,
			expMap:    map[string]int{},
			expTotalA: 1,
		},
		{
			a:         []rune("aabbaa"),
			b:         []rune("ababbaa"),
			size:      7,
			expMap:    map[string]int{},
			expTotalB: 1,
		},
		{
			a:      []rune("ababbaa"),
			b:      []rune("aabbaa"),
			size:   8,
			expMap: map[string]int{},
		},
		{
			a:      []rune("aabbaa"),
			b:      []rune("ababbaa"),
			size:   8,
			expMap: map[string]int{},
		},
	}

	for _, input := range inputs {
		actMap, actTotal, actTotalA, actTotalB := util.NgramIntersection(input.a, input.b, input.size)
		require.Equal(t, input.expMap, actMap)
		require.Equal(t, input.expTotal, actTotal)
		require.Equal(t, input.expTotalA, actTotalA)
		require.Equal(t, input.expTotalB, actTotalB)
	}
}

func requireEqual(t *testing.T, inputs [][2]interface{}) {
	for _, input := range inputs {
		require.Equal(t, input[0], input[1])
	}
}
