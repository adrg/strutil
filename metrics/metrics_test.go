package metrics_test

import (
	"testing"

	"github.com/adrg/strutil/metrics"
	"github.com/stretchr/testify/require"
)

func TestHamming(t *testing.T) {
	h := metrics.NewHamming()
	require.Zero(t, h.Distance("", ""))
	require.Equal(t, 0.75, h.Compare("text", "test"))
	require.Equal(t, 0.5, h.Compare("once", "one"))
	h.CaseSensitive = false
	require.Equal(t, 0.5, h.Compare("one", "ONCE"))
}
