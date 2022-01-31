package mathutil_test

import (
	"testing"

	"github.com/adrg/strutil/internal/mathutil"
	"github.com/stretchr/testify/require"
)

func TestMin(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{0, mathutil.Min()},
		{1, mathutil.Min(1)},
		{0, mathutil.Min(0, 1)},
		{1, mathutil.Min(1, 1)},
		{1, mathutil.Min(2, 1)},
		{1, mathutil.Min(1, 2)},
		{0, mathutil.Min(2, 1, 0)},
		{0, mathutil.Min(0, 1, 2)},
	})
}

func TestMax(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{0, mathutil.Max()},
		{1, mathutil.Max(1)},
		{1, mathutil.Max(0, 1)},
		{1, mathutil.Max(1, 1)},
		{2, mathutil.Max(2, 1)},
		{2, mathutil.Max(1, 2)},
		{3, mathutil.Max(2, 1, 3)},
		{3, mathutil.Max(3, 1, 2)},
	})
}

func TestMinf(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{0.0, mathutil.Minf()},
		{1.0, mathutil.Minf(1.0)},
		{0.0, mathutil.Minf(0.0, 1.0)},
		{1.0, mathutil.Minf(1.0, 1.0)},
		{1.0, mathutil.Minf(2.0, 1.0)},
		{1.0, mathutil.Minf(1.0, 2.0)},
		{0.0, mathutil.Minf(2.0, 1.0, 0.0)},
		{0.0, mathutil.Minf(0.0, 1.0, 2.0)},
	})
}

func TestMaxf(t *testing.T) {
	requireEqual(t, [][2]interface{}{
		{0.0, mathutil.Maxf()},
		{1.0, mathutil.Maxf(1.0)},
		{1.0, mathutil.Maxf(0.0, 1.0)},
		{1.0, mathutil.Maxf(1.0, 1.0)},
		{2.0, mathutil.Maxf(2.0, 1.1, 1.0)},
		{2.0, mathutil.Maxf(1.1, 1.0, 2.0)},
		{3.0, mathutil.Maxf(2.0, 1.0, 3.0)},
		{3.0, mathutil.Maxf(3.0, 1.0, 2.0)},
	})
}

func requireEqual(t *testing.T, inputs [][2]interface{}) {
	t.Helper()

	for _, input := range inputs {
		require.Equal(t, input[0], input[1])
	}
}
