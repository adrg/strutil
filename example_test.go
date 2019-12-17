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

func ExampleUniqueSlice() {
	sample := []string{"a", "b", "a", "b", "b", "c"}
	fmt.Println("[a b a b b c]:", strutil.UniqueSlice(sample))

	// Output:
	// [a b a b b c]: [a b c]
}

func ExampleSliceContains() {
	terms := []string{"a", "b", "c"}
	fmt.Println("([a b c], b):", strutil.SliceContains(terms, "b"))
	fmt.Println("([a b c], d):", strutil.SliceContains(terms, "d"))

	// Output:
	// ([a b c], b): true
	// ([a b c], d): false
}

func ExampleNgrams() {
	fmt.Println("abbcd n-grams (size 2):", strutil.Ngrams("abbcd", 2))
	fmt.Println("abbcd n-grams (size 3):", strutil.Ngrams("abbcd", 3))

	// Output:
	// abbcd n-grams (size 2): [ab bb bc cd]
	// abbcd n-grams (size 3): [abb bbc bcd]
}

func ExampleNgramMap() {
	// 2 character n-gram map.
	ngrams, total := strutil.NgramMap("abbcabb", 2)
	fmt.Printf("abbcabb n-gram map (size 2): %v (%d ngrams)\n", ngrams, total)

	// 3 character n-gram map.
	ngrams, total = strutil.NgramMap("abbcabb", 3)
	fmt.Printf("abbcabb n-gram map (size 3): %v (%d ngrams)\n", ngrams, total)

	// Output:
	// abbcabb n-gram map (size 2): map[ab:2 bb:2 bc:1 ca:1] (6 ngrams)
	// abbcabb n-gram map (size 3): map[abb:2 bbc:1 bca:1 cab:1] (5 ngrams)
}

func ExampleNgramIntersection() {
	ngrams, common, totalA, totalB := strutil.NgramIntersection("ababc", "ababd", 2)
	fmt.Printf("(ababc, ababd) n-gram intersection: %v (%d/%d n-grams)\n",
		ngrams, common, totalA+totalB)

	// Output:
	// (ababc, ababd) n-gram intersection: map[ab:2 ba:1] (3/8 n-grams)
}
