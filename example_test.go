package strutil_test

import (
	"fmt"

	"github.com/adrg/strutil"
)

func ExampleLevenshtein() {
	// Default options.
	dist, sim := strutil.Levenshtein("book", "brick", nil)
	fmt.Printf("(book, brick): distance: %d, similarity: %.2f\n", dist, sim)

	// Custom options.
	opts := &strutil.LevOpts{
		CaseSensitive: false,
		InsertCost:    1,
		DeleteCost:    1,
		ReplaceCost:   2,
	}

	dist, sim = strutil.Levenshtein("HELLO", "jello", opts)
	fmt.Printf("(HELLO, jello): distance: %d, similarity: %.2f\n", dist, sim)

	// Output:
	// (book, brick): distance: 3, similarity: 0.40
	// (HELLO, jello): distance: 2, similarity: 0.60
}

func ExampleJaro() {
	fmt.Printf("(sort, shirt): %.2f\n", strutil.Jaro("sort", "shirt"))

	// Output:
	// (sort, shirt): 0.62
}

func ExampleJaroWinkler() {
	fmt.Printf("(sort, shirt): %.2f\n", strutil.JaroWinkler("sort", "shirt"))

	// Output:
	// (sort, shirt): 0.66
}

func ExampleSmithWatermanGotoh() {
	// Default options.
	sim := strutil.SmithWatermanGotoh("a pink kitten", "a kitten", nil)
	fmt.Printf("(a pink kitten, a kitten): %.2f\n", sim)

	// Custom options.
	opts := &strutil.SWGOpts{
		CaseSensitive: false,
		GapPenalty:    -0.1,
		Substitution: strutil.SWGMatchMismatch{
			Match:    1,
			Mismatch: -0.5,
		},
	}

	sim = strutil.SmithWatermanGotoh("a pink kitten", "A KITTEN", opts)
	fmt.Printf("(a pink kitten, A KITTEN): %.2f\n", sim)

	// Output:
	// (a pink kitten, a kitten): 0.88
	// (a pink kitten, A KITTEN): 0.94
}

func ExampleCommonPrefix() {
	fmt.Println("(answer, anvil):", strutil.CommonPrefix("answer", "anvil"))
	// Output:
	// (answer, anvil): an
}

func ExampleUnique() {
	sample := []string{"a", "b", "a", "b", "b", "c"}
	fmt.Println("[a b a b b c]:", strutil.Unique(sample))
	// Output:
	// [a b a b b c]: [a b c]
}
