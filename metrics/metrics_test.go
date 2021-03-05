package metrics_test

import (
	"fmt"
	"testing"

	"github.com/adrg/strutil/metrics"
	"github.com/stretchr/testify/require"
)

func sf(a float64) string {
	return fmt.Sprintf("%.2f", a)
}

func TestHamming(t *testing.T) {
	h := metrics.NewHamming()
	require.Equal(t, 0, h.Distance("", ""))
	require.Equal(t, "0.75", sf(h.Compare("text", "test")))
	require.Equal(t, "0.50", sf(h.Compare("once", "one")))
	h.CaseSensitive = false
	require.Equal(t, "0.50", sf(h.Compare("one", "ONCE")))
}

func TestJaccard(t *testing.T) {
	j := metrics.NewJaccard()
	require.Equal(t, "1.00", sf(j.Compare("", "")))
	require.Equal(t, "0.00", sf(j.Compare("a", "b")))
	require.Equal(t, "0.43", sf(j.Compare("night", "alright")))
	j.NgramSize = 0
	require.Equal(t, "0.43", sf(j.Compare("night", "alright")))
	j.CaseSensitive = false
	j.NgramSize = 3
	require.Equal(t, "0.33", sf(j.Compare("NIGHT", "alright")))
}

func TestJaro(t *testing.T) {
	j := metrics.NewJaro()
	require.Equal(t, "1.00", sf(j.Compare("", "")))
	require.Equal(t, "0.00", sf(j.Compare("test", "")))
	require.Equal(t, "0.00", sf(j.Compare("a", "b")))
	require.Equal(t, "0.78", sf(j.Compare("sort", "shirt")))
	require.Equal(t, "0.64", sf(j.Compare("sort", "report")))
	j.CaseSensitive = false
	require.Equal(t, "0.78", sf(j.Compare("sort", "SHIRT")))
}

func TestJaroWinkler(t *testing.T) {
	j := metrics.NewJaroWinkler()
	require.Equal(t, "1.00", sf(j.Compare("", "")))
	require.Equal(t, "0.00", sf(j.Compare("test", "")))
	require.Equal(t, "0.80", sf(j.Compare("sort", "shirt")))
	require.Equal(t, "0.94", sf(j.Compare("charm", "charmed")))
	j.CaseSensitive = false
	require.Equal(t, "0.80", sf(j.Compare("sort", "SHIRT")))
}

func TestLevenshtein(t *testing.T) {
	l := metrics.NewLevenshtein()
	require.Equal(t, 0, l.Distance("", ""))
	require.Equal(t, 4, l.Distance("test", ""))
	require.Equal(t, 4, l.Distance("", "test"))
	require.Equal(t, "0.40", sf(l.Compare("book", "brick")))
	l.CaseSensitive = false
	require.Equal(t, "0.80", sf(l.Compare("hello", "jello")))
	l.ReplaceCost = 2
	require.Equal(t, "0.60", sf(l.Compare("hello", "JELLO")))
}

func TestOperlapCoefficient(t *testing.T) {
	o := metrics.NewOverlapCoefficient()
	require.Equal(t, "1.00", sf(o.Compare("", "")))
	require.Equal(t, "0.75", sf(o.Compare("night", "alright")))
	require.Equal(t, "0.00", sf(o.Compare("aa", "")))
	require.Equal(t, "0.00", sf(o.Compare("bb", "")))
	o.NgramSize = 0
	require.Equal(t, "0.75", sf(o.Compare("night", "alright")))
	require.Equal(t, "1.00", sf(o.Compare("aa", "aaaa")))
	o.CaseSensitive = false
	require.Equal(t, "1.00", sf(o.Compare("aa", "AAAA")))
	o.NgramSize = 3
	require.Equal(t, "0.67", sf(o.Compare("night", "alright")))
}

func TestSmithWatermanGotoh(t *testing.T) {
	s := metrics.NewSmithWatermanGotoh()
	require.Equal(t, "1.00", sf(s.Compare("", "")))
	require.Equal(t, "0.00", sf(s.Compare("test", "")))
	require.Equal(t, "0.00", sf(s.Compare("", "test")))
	require.Equal(t, "0.88", sf(s.Compare("a pink kitten", "a kitten")))
	s.Substitution = nil
	require.Equal(t, "0.88", sf(s.Compare("a pink kitten", "a kitten")))
	s.CaseSensitive = false
	s.GapPenalty = -0.1
	s.Substitution = metrics.MatchMismatch{
		Match:    1,
		Mismatch: -0.5,
	}
	require.Equal(t, "0.94", sf(s.Compare("a pink kitten", "A KITTEN")))
}

func TestSorensenDice(t *testing.T) {
	s := metrics.NewSorensenDice()
	require.Equal(t, "1.00", sf(s.Compare("", "")))
	require.Equal(t, "0.00", sf(s.Compare("a", "b")))
	require.Equal(t, "0.60", sf(s.Compare("night", "alright")))
	s.NgramSize = 0
	require.Equal(t, "0.60", sf(s.Compare("night", "alright")))
	s.CaseSensitive = false
	require.Equal(t, "0.60", sf(s.Compare("night", "ALRIGHT")))
	s.NgramSize = 3
	require.Equal(t, "0.50", sf(s.Compare("night", "alright")))
}

func TestMatchMismatch(t *testing.T) {
	m := metrics.MatchMismatch{
		Match:    2,
		Mismatch: 1,
	}
	require.Equal(t, "1.00", sf(m.Compare([]rune{'a'}, 0, []rune{'b'}, 0)))
	require.Equal(t, "2.00", sf(m.Compare([]rune{'a'}, 0, []rune{'a'}, 0)))
	require.Equal(t, "1.00", sf(m.Min()))
	require.Equal(t, "2.00", sf(m.Max()))
}
