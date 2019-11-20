strutil
=======
[![Build Status](https://travis-ci.org/adrg/strutil.svg?branch=master)](https://travis-ci.org/adrg/strutil)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/adrg/strutil)
[![License: MIT](https://img.shields.io/badge/license-MIT-red.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/adrg/strutil)](https://goreportcard.com/report/github.com/adrg/strutil)

strutil provides string metrics for calculating string similarity as well as
other string utility functions.  
Full documentation can be found at: https://godoc.org/github.com/adrg/strutil

## Installation

```
go get github.com/adrg/strutil
```

## Usage

#### Levenshtein

```go
distance, similarity := strutil.Levenshtein("graph", "giraffe", nil)
fmt.Printf("%d, %.2f\n", distance, similarity)

// Output: 4, 0.43
```

```go
// Customomize operation costs.
opts := &strutil.LevOpts{
    CaseSensitive: false,
    InsertCost:    1,
    DeleteCost:    1,
    ReplaceCost:   2,
}

distance, similarity := strutil.Levenshtein("make", "Cake", opts)
fmt.Printf("%d, %.2f\n", distance, similarity)

// Output: 2, 0.50
```
More information and additional examples can be found on
[GoDoc](https://godoc.org/github.com/adrg/strutil#Levenshtein).

#### Jaro

```go
similarity := strutil.Jaro("think", "tank")
fmt.Printf("%.2f\n", similarity)

// Output: 0.78
```

More information and additional examples can be found on
[GoDoc](https://godoc.org/github.com/adrg/strutil#Jaro).

#### Jaro-Winkler

```go
similarity := strutil.JaroWinkler("think", "tank")
fmt.Printf("%.2f\n", similarity)

// Output: 0.80
```

More information and additional examples can be found on
[GoDoc](https://godoc.org/github.com/adrg/strutil#JaroWinkler).

#### Smith-Waterman-Gotoh

```go
similarity := strutil.SmithWatermanGotoh("times roman", "times new roman", nil)
fmt.Printf("%.2f\n", similarity)

// Output: 0.82
```

```go
// Customize gap penalty and substitution function.
opts := &strutil.SWGOpts{
    CaseSensitive: false,
    GapPenalty:    -0.1,
    Substitution: strutil.SWGMatchMismatch{
        Match:    1,
        Mismatch: -0.5,
    },
}

similarity := strutil.SmithWatermanGotoh("times roman", "times new roman", opts)
fmt.Printf("%.2f\n", similarity)

// Output: 0.96
```

More information and additional examples can be found on
on [GoDoc](https://godoc.org/github.com/adrg/strutil#SmithWatermanGotoh).

#### Unique

```go
sample := []string{"cats", "dogs", "tigers", "dogs", "dogs", "cats", "mice"}
fmt.Println(strutil.Unique(sample))

// Output: [cats dogs tigers mice]
```

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
