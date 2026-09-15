package metrics_test

import (
	"testing"

	"github.com/adrg/strutil/metrics"
	"github.com/stretchr/testify/require"
)

// TestLevenshteinCustomCostWeights verifies that the documented InsertCost,
// DeleteCost and ReplaceCost fields actually weight the Wagner-Fischer border
// base cases. Pre-fix the DP border used the bare row index (i / i+1) instead
// of the cost-weighted value, so Distance silently undercounted whenever a
// non-unit insert/delete cost was set. The unit-cost (1/1/1) results are
// unchanged because 1*c == c.
func TestLevenshteinCustomCostWeights(t *testing.T) {
	cases := []struct {
		name            string
		a, b            string
		ins, del, sub   int
		wantDistance    int
		wantComparePrec string // first 2 decimals of Compare, "" to skip
	}{
		// --- smoking gun: cost-bound impossibility -------------------------
		// "x" -> "abcx" needs 3 insertions; with InsertCost=10 the true minimum
		// is 30. Pre-fix adrg returned 3, which is less than the cost of a SINGLE
		// insert (10) -> logically impossible under cost-weighted semantics.
		// Compare = 1 - dist/maxLen; with dist=30, maxLen=4 -> 1 - 7.5 = -6.50
		// (Compare can go negative with non-unit costs; pre-fix it wrongly read 0.25).
		{"smoking_gun_insert_x_abcx", "x", "abcx", 10, 1, 1, 30, "-6.50"},
		{"insert_x_abcx_cost7", "x", "abcx", 7, 1, 1, 21, "-4.25"},

		// mirror: heavy delete cost
		{"delete_abcx_x_cost10", "abcx", "x", 1, 10, 1, 30, "-6.50"},
		{"delete_abcx_x_cost7", "abcx", "x", 1, 7, 1, 21, "-4.25"},

		// pure-insert / pure-delete paths
		{"insert_a_abcde_cost7", "a", "abcde", 7, 1, 1, 28, "-4.60"},
		{"delete_abcde_a_cost7", "abcde", "a", 1, 7, 1, 28, "-4.60"},

		// mixed-cost classic edit-distance examples (Compare skipped where
		// 1 - dist/maxLen does not round to an exact 2-decimal value).
		{"kitten_sitting_235", "kitten", "sitting", 2, 3, 5, 12, ""},
		{"sitting_kitten_235", "sitting", "kitten", 2, 3, 5, 13, ""},
		{"flaw_lawn_448", "flaw", "lawn", 4, 4, 8, 8, "-1.00"},
		{"intention_execution_224", "intention", "execution", 2, 2, 4, 16, ""},
		{"GGACT_GAGCG_325", "GGACT", "GAGCG", 3, 2, 5, 10, "-1.00"},
		{"book_back_223", "book", "back", 2, 2, 3, 6, "-0.50"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			l := metrics.NewLevenshtein()
			l.InsertCost = c.ins
			l.DeleteCost = c.del
			l.ReplaceCost = c.sub
			got := l.Distance(c.a, c.b)
			require.Equal(t, c.wantDistance, got,
				"Distance(%q,%q) ins=%d del=%d sub=%d", c.a, c.b, c.ins, c.del, c.sub)
			if c.wantComparePrec != "" {
				require.Equal(t, c.wantComparePrec, sf(l.Compare(c.a, c.b)),
					"Compare(%q,%q) ins=%d del=%d sub=%d", c.a, c.b, c.ins, c.del, c.sub)
			}
		})
	}
}

// TestLevenshteinUnitCostUnchanged guards that the fix does not change any
// default-cost (1/1/1) result: with unit costs the bare index equals the
// cost-weighted value, so the old and new code must agree exactly.
func TestLevenshteinUnitCostUnchanged(t *testing.T) {
	cases := []struct {
		a, b         string
		wantDistance int
		wantCompare  string
	}{
		{"kitten", "sitting", 3, "0.57"},
		{"flaw", "lawn", 2, "0.50"},
		{"intention", "execution", 5, "0.44"},
		{"GGACT", "GAGCG", 3, "0.40"},
		{"book", "back", 2, "0.50"},
		{"text", "test", 1, "0.75"},
		{"ab\u2019d", "ab\u2019c", 1, "0.75"},
		{"hello", "jello", 1, "0.80"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.a+"_"+c.b, func(t *testing.T) {
			l := metrics.NewLevenshtein()
			require.Equal(t, c.wantDistance, l.Distance(c.a, c.b))
			require.Equal(t, c.wantCompare, sf(l.Compare(c.a, c.b)))
		})
	}
}

// TestLevenshteinEmptyStringFastPaths covers the early-return paths that
// already multiplied by the cost; these were correct pre-fix and must remain
// correct post-fix, for both unit and non-unit costs.
func TestLevenshteinEmptyStringFastPaths(t *testing.T) {
	// both empty
	require.Equal(t, 0, metrics.NewLevenshtein().Distance("", ""))

	// a empty -> b filled: cost = InsertCost * len(b)
	l := metrics.NewLevenshtein()
	l.InsertCost = 5
	l.DeleteCost = 1
	l.ReplaceCost = 1
	require.Equal(t, 15, l.Distance("", "abc"))

	// b empty -> a filled: cost = DeleteCost * len(a)
	l2 := metrics.NewLevenshtein()
	l2.InsertCost = 1
	l2.DeleteCost = 5
	l2.ReplaceCost = 1
	require.Equal(t, 15, l2.Distance("abc", ""))

	// Compare with empty string. distance("", "abc") = InsertCost*3 = 15,
	// maxLen = 3, so Compare = 1 - 15/3 = -4.00 (negative with non-unit costs).
	// (Compare("","") is NaN = 0/0, an unrelated pre-existing edge case, not
	// asserted here.)
	l3 := metrics.NewLevenshtein()
	l3.InsertCost = 5
	require.Equal(t, "-4.00", sf(l3.Compare("", "abc")))
}

// TestLevenshteinCostBoundSmokingGun is the explicit cost-bound assertion:
// transforming "x" into "abcx" requires at least one insertion, so with
// InsertCost=10 any correct edit-script cost must be >= 10. Pre-fix the
// library returned 3, violating this lower bound.
func TestLevenshteinCostBoundSmokingGun(t *testing.T) {
	l := metrics.NewLevenshtein()
	l.InsertCost = 10
	l.DeleteCost = 1
	l.ReplaceCost = 1
	got := l.Distance("x", "abcx")
	// "abcx" is 4 chars, "x" is 1; the only difference is the prefix "abc" (3
	// insertions), so the exact cost is 3 * InsertCost = 30.
	require.Equal(t, 30, got)
	require.GreaterOrEqual(t, got, l.InsertCost,
		"any edit script from %q to %q needs >=1 insert; cost must be >= InsertCost", "x", "abcx")
}
