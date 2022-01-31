package ngram_test

import (
	"testing"

	"github.com/adrg/strutil/internal/ngram"
	"github.com/stretchr/testify/require"
)

func TestNgramCount(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{0, ngram.Count(nil, -1)},
		{0, ngram.Count(nil, 0)},
		{0, ngram.Count(nil, 1)},
		{0, ngram.Count([]rune{}, -1)},
		{0, ngram.Count([]rune{}, 0)},
		{0, ngram.Count([]rune{}, 1)},
		{6, ngram.Count([]rune("abbabb"), -1)},
		{6, ngram.Count([]rune("abbabb"), 0)},
		{6, ngram.Count([]rune("abbabb"), 1)},
		{5, ngram.Count([]rune("abbabb"), 2)},
		{4, ngram.Count([]rune("abbabb"), 3)},
		{3, ngram.Count([]rune("abbabb"), 4)},
		{2, ngram.Count([]rune("abbabb"), 5)},
		{1, ngram.Count([]rune("abbabb"), 6)},
		{0, ngram.Count([]rune("abbabb"), 7)},
		{0, ngram.Count([]rune("abbabb"), 8)},
	})
}

func TestNgrams(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{0, len(ngram.Slice(nil, -1))},
		{0, len(ngram.Slice(nil, 0))},
		{0, len(ngram.Slice(nil, 1))},
		{0, len(ngram.Slice([]rune{}, -1))},
		{0, len(ngram.Slice([]rune{}, 0))},
		{0, len(ngram.Slice([]rune{}, 1))},
		{
			[]string{"a", "b", "c", "d", "e", "f"},
			ngram.Slice([]rune("abcdef"), -1),
		},
		{
			[]string{"a", "b", "c", "d", "e", "f"},
			ngram.Slice([]rune("abcdef"), 0),
		},
		{
			[]string{"a", "b", "c", "d", "e", "f"},
			ngram.Slice([]rune("abcdef"), 1),
		},
		{
			[]string{"ab", "bc", "cd", "de", "ef"},
			ngram.Slice([]rune("abcdef"), 2),
		},
		{
			[]string{"abc", "bcd", "cde", "def"},
			ngram.Slice([]rune("abcdef"), 3),
		},
		{
			[]string{"abcd", "bcde", "cdef"},
			ngram.Slice([]rune("abcdef"), 4),
		},
		{
			[]string{"abcde", "bcdef"},
			ngram.Slice([]rune("abcdef"), 5),
		},
		{
			[]string{"abcdef"},
			ngram.Slice([]rune("abcdef"), 6),
		},
		{
			0,
			len(ngram.Slice([]rune("abcdef"), 7)),
		},
		{
			0,
			len(ngram.Slice([]rune("abcdef"), 8)),
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
		{
			term:     []rune("abbabb"),
			size:     8,
			expMap:   map[string]int{},
			expTotal: 0,
		},
	}

	for _, input := range inputs {
		actMap, actTotal := ngram.Map(input.term, input.size)
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
		{
			a:      []rune("ababbaa"),
			b:      []rune("aabbaa"),
			size:   9,
			expMap: map[string]int{},
		},
		{
			a:      []rune("aabbaa"),
			b:      []rune("ababbaa"),
			size:   9,
			expMap: map[string]int{},
		},
	}

	for _, input := range inputs {
		actMap, actTotal, actTotalA, actTotalB := ngram.Intersection(input.a, input.b, input.size)
		require.Equal(t, input.expMap, actMap)
		require.Equal(t, input.expTotal, actTotal)
		require.Equal(t, input.expTotalA, actTotalA)
		require.Equal(t, input.expTotalB, actTotalB)
	}
}

func requireEqual(t *testing.T, inputs [][2]interface{}) {
	t.Helper()

	for _, input := range inputs {
		require.Equal(t, input[0], input[1])
	}
}
