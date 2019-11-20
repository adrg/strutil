strutil
=======
[![Build Status](https://travis-ci.org/adrg/strutil.svg?branch=master)](https://travis-ci.org/adrg/strutil)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/adrg/strutil)
[![License: MIT](https://img.shields.io/badge/license-MIT-red.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/adrg/strutil)](https://goreportcard.com/report/github.com/adrg/strutil)

strutil provides string metrics for calculating string similarity as well as
other string utility functions.  
Full documentation can be found at: https://godoc.org/github.com/adrg/strutil.

## Installation

```
go get github.com/adrg/strutil
```

## String metrics

- [Levenshtein](#levenshtein)
- [Jaro](#jaro)
- [Jaro-Winkler](#jaro-winkler)
- [Smith-Waterman-Gotoh](#smith-waterman-gotoh)
- [Sorensen-Dice](#sorensen-dice)

The package defines the `StringMetric` interface, which is implemented by all
the string metrics. The interface is used with the `Similarity` function, which
calculates the similarity between the specified strings, using the provided
string metric.

```go
type StringMetric interface {
	Compare(a, b string) float64
}

func Similarity(a, b string, metric StringMetric) float64 {
}
```

All defined string metrics can be found in the
[metrics](https://godoc.org/github.com/adrg/strutil/metrics) package.

#### Levenshtein

Calculate similarity using default options.
```go
similarity := strutil.Similarity("graph", "giraffe", metrics.NewLevenshtein())
fmt.Printf("%.2f\n", similarity) // Output: 0.43
```

Configure edit operation costs.
```go
lev := metrics.NewLevenshtein()
lev.CaseSensitive = false
lev.InsertCost = 1
lev.ReplaceCost = 2
lev.DeleteCost = 1

similarity := strutil.Similarity("make", "Cake", lev)
fmt.Printf("%.2f\n", similarity) // Output: 0.50
```

Calculate distance.
```go
lev := metrics.NewLevenshtein()
fmt.Printf("%d\n", lev.Distance("graph", "giraffe")) // Output: 4
```

More information and additional examples can be found on
[GoDoc](https://godoc.org/github.com/adrg/strutil/metrics#Levenshtein).

#### Jaro

```go
similarity := strutil.Similarity("think", "tank", metrics.NewJaro())
fmt.Printf("%.2f\n", similarity) // Output: 0.78
```

More information and additional examples can be found on
[GoDoc](https://godoc.org/github.com/adrg/strutil/metrics#Jaro).

#### Jaro-Winkler

```go
similarity := strutil.Similarity("think", "tank", metrics.NewJaroWinkler())
fmt.Printf("%.2f\n", similarity) // Output: 0.80
```

More information and additional examples can be found on
[GoDoc](https://godoc.org/github.com/adrg/strutil/metrics#JaroWinkler).

#### Smith-Waterman-Gotoh

Calculate similarity using default options.
```go
swg := metrics.NewSmithWatermanGotoh()
similarity := strutil.Similarity("times roman", "times new roman", swg)
fmt.Printf("%.2f\n", similarity) // Output: 0.82
```

Customize gap penalty and substitution function.
```go
swg := metrics.NewSmithWatermanGotoh()
swg.CaseSensitive = false
swg.GapPenalty = -0.1
swg.Substitution = metrics.MatchMismatch {
    Match:    1,
    Mismatch: -0.5,
}

similarity := strutil.Similarity("Times Roman", "times new roman", swg)
fmt.Printf("%.2f\n", similarity) // Output: 0.96
```

More information and additional examples can be found on
on [GoDoc](https://godoc.org/github.com/adrg/strutil/metrics#SmithWatermanGotoh).

#### Sorensen-Dice

Calculate similarity using default options.
```go
sd := metrics.NewSorensenDice()
similarity := strutil.Similarity("time to make haste", "no time to waste", sd)
fmt.Printf("%.2f\n", similarity) // Output: 0.62
```

Customize n-gram size.
```go
sd := metrics.NewSorensenDice()
sd.CaseSensitive = false
sd.NgramSize = 3

similarity := strutil.Similarity("Time to make haste", "no time to waste", sd)
fmt.Printf("%.2f\n", similarity) // Output: 0.80
```

More information and additional examples can be found on
on [GoDoc](https://godoc.org/github.com/adrg/strutil/metrics#SorensenDice).

## References

For more information see:
- [Levenshtein distance](https://en.wikipedia.org/wiki/Levenshtein_distance)
- [Jaro-Winkler distance](https://en.wikipedia.org/wiki/Jaro-Winkler_distance)
- [Smith-Waterman algorithm](https://en.wikipedia.org/wiki/Smith-Waterman_algorithm)
- [Sorensen-Dice coefficient](https://en.wikipedia.org/wiki/Sorensen–Dice_coefficient)

## Contributing

Contributions in the form of pull requests, issues or just general feedback,
are always welcome.
See [CONTRIBUTING.MD](https://github.com/adrg/strutil/blob/master/CONTRIBUTING.md).

## License
Copyright (c) 2019 Adrian-George Bostan.

This project is licensed under the [MIT license](https://opensource.org/licenses/MIT).
See [LICENSE](https://github.com/adrg/strutil/blob/master/LICENSE) for more details.
