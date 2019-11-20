package metrics_test

import (
	"fmt"

	"github.com/adrg/strutil/metrics"
)

func ExampleLevenshtein() {
	// Default options.
	lev := metrics.NewLevenshtein()

	sim := lev.Compare("book", "brick")
	fmt.Printf("(book, brick) similarity: %.2f\n", sim)

	dist := lev.Distance("book", "brick")
	fmt.Printf("(book, brick) distance: %d\n", dist)

	// Custom options.
	lev.CaseSensitive = false
	lev.ReplaceCost = 2

	sim = lev.Compare("HELLO", "jello")
	fmt.Printf("(HELLO, jello) similarity: %.2f\n", sim)

	dist = lev.Distance("HELLO", "jello")
	fmt.Printf("(HELLO, jello) distance: %d\n", dist)

	// Output:
	// (book, brick) similarity: 0.40
	// (book, brick) distance: 3
	// (HELLO, jello) similarity: 0.60
	// (HELLO, jello) distance: 2
}

func ExampleJaro() {
	jaro := metrics.NewJaro()
	sim := jaro.Compare("sort", "shirt")
	fmt.Printf("(sort, shirt) similarity: %.2f\n", sim)

	// Output:
	// (sort, shirt) similarity: 0.78
}

func ExampleJaroWinkler() {
	jw := metrics.NewJaroWinkler()
	sim := jw.Compare("sort", "shirt")
	fmt.Printf("(sort, shirt) similarity: %.2f\n", sim)

	// Output:
	// (sort, shirt) similarity: 0.80
}

func ExampleSmithWatermanGotoh() {
	// Default options.
	swg := metrics.NewSmithWatermanGotoh()

	sim := swg.Compare("a pink kitten", "a kitten")
	fmt.Printf("(a pink kitten, a kitten) similarity: %.2f\n", sim)

	// Custom options.
	swg.CaseSensitive = false
	swg.GapPenalty = -0.1
	swg.Substitution = metrics.MatchMismatch{
		Match:    1,
		Mismatch: -0.5,
	}

	sim = swg.Compare("a pink kitten", "A KITTEN")
	fmt.Printf("(a pink kitten, A KITTEN) similarity: %.2f\n", sim)

	// Output:
	// (a pink kitten, a kitten) similarity: 0.88
	// (a pink kitten, A KITTEN) similarity: 0.94
}

func ExampleSorensenDice() {
	// Default options.
	sd := metrics.NewSorensenDice()
	sim := sd.Compare("night", "alight")
	fmt.Printf("(night, alright) similarity: %.2f\n", sim)

	// Custom options.
	sd.CaseSensitive = false
	sd.NgramSize = 3

	sim = sd.Compare("night", "alright")
	fmt.Printf("(night, alright) similarity: %.2f\n", sim)

	// Output:
	// (night, alright) similarity: 0.67
	// (night, alright) similarity: 0.75
}
