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
More information and additional examples can be found on
[GoDoc](https://godoc.org/github.com/adrg/strutil#Levenshtein).

#### Jaro

```go
similarity = strutil.Jaro("think", "tank")
fmt.Printf("%.2f\n", similarity)

// Output: 0.62
```

More information and additional examples can be found on
[GoDoc](https://godoc.org/github.com/adrg/strutil#Jaro).

#### Jaro-Winkler

```go
similarity = strutil.JaroWinkler("think", "tank")
fmt.Printf("%.2f\n", similarity)

// Output: 0.66
```

More information and additional examples can be found on
[GoDoc](https://godoc.org/github.com/adrg/strutil#JaroWinkler).

#### Smith-Waterman-Gotoh

```go
similarity = strutil.SmithWatermanGotoh("times roman", "times new roman", nil)
fmt.Printf("%.2f\n", similarity)

// Output: 0.82
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
- [Jaro-Winkler distance](https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance)
- [Smith-Waterman algorithm](https://en.wikipedia.org/wiki/Smith%E2%80%93Waterman_algorithm)

## Contributing

Contributions in the form of pull requests, issues or just general feedback,
are always welcome.
See [CONTRIBUTING.MD](https://github.com/adrg/strutil/blob/master/CONTRIBUTING.md).

## License
Copyright (c) 2019 Adrian-George Bostan.

This project is licensed under the [MIT license](https://opensource.org/licenses/MIT).
See [LICENSE](https://github.com/adrg/strutil/blob/master/LICENSE) for more details.
