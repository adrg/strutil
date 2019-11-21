package strutil_test

import (
	"fmt"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

func ExampleSimilarity() {
	sim := strutil.Similarity("riddle", "needle", metrics.NewJaroWinkler())
	fmt.Printf("(riddle, needle) similarity: %.2f\n", sim)

	// Output:
	// (riddle, needle) similarity: 0.56
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
