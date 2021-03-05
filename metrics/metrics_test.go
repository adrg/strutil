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
