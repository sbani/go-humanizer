# Humanizer
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.md) [![Build Status](https://travis-ci.org/sbani/go-humanizer.svg?branch=master)](https://travis-ci.org/sbani/go-humanizer) [![GoDoc](https://godoc.org/github.com/sbani/go-humanizer?status.svg)](https://godoc.org/github.com/sbani/go-humanizer) [![codecov.io](https://codecov.io/github/sbani/go-humanizer/coverage.svg?branch=master)](https://codecov.io/github/sbani/go-humanizer?branch=master)


Humanize values to make them easier to read.

## Installation
```bash
go get github.com/sbani/go-humanizer
```

## Usage
### Strings
#### Humanize:
```go
import "github.com/sbani/go-humanizer/strings"

Humanize("news_count", true) // "News count"
Humanize("User", false) // "user"
Humanize("news_id", true) // "News"
```
#### Truncate:
Truncate string but never cut within a word.
```go
import "github.com/sbani/go-humanizer/strings"

textShort := "Short text"
Truncate(textShort, 1) // Short
Truncate(textShort, 6) // Short
Truncate(textShort, 7) // Short text
```
### Numbers
#### Ordinalize:
```go
import "github.com/sbani/go-humanizer/numbers"

Ordinalize(0) // "0th"
Ordinalize(1) // "1st"
Ordinalize(2) // "2nd"
Ordinalize(23) // "23rd"
Ordinalize(1002) // "1002nd"
Ordinalize(-111) // "-111th"
```
#### Ordinal:
```go
import "github.com/sbani/go-humanizer/numbers"

Ordinal(0) // "th"
Ordinal(1) // "st"
Ordinal(2) // "nd"
Ordinal(23) // "rd"
Ordinal(1002) // "nd"
Ordinal(-111) // "th"
```
#### Roman:
```go
import "github.com/sbani/go-humanizer/numbers"

s, err := ToRoman(1) // "I"
s, err := ToRomanToRoman(5) // "V"
s, err := ToRomanToRoman(1300) // "MCCC"

i, err := ToRomanFromRoman("MMMCMXCIX") // 3999
i, err := ToRomanFromRoman("V") // 5
i, err := ToRomanFromRoman("CXXV") // 125
```
#### Binary Suffix
```go
import "github.com/sbani/go-humanizer/numbers"

s := BinarySuffix(0) // "0 bytes"
s := BinarySuffix(1536) // "1.5 kB"
s := BinarySuffix(1048576 * 5) // "5.00 MB"
s := BinarySuffix(1073741824 * 2) // "2.00 GB"
```
### Collection
#### Oxford
```go
import "github.com/sbani/go-humanizer/collection"

Oxford([]string{"Albert"}, -1) // "Albert"
Oxford([]string{"Albert", "Norbert"}, -1) // "Albert and Norbert"
Oxford([]string{"Albert", "Norbert", "Michael", "Kevin"}, -1) // "Albert, Norbert, Michael and Kevin"
Oxford([]string{"Albert", "Norbert", "Michael", "Kevin"}, 2)) // Albert, Norbert and 2 more
```

## License
MIT License. See LICENSE file for more informations.

## Credits
A special WOW goes to [PHP Humanizer](https://github.com/coduo/php-humanizer). (This lib is just a port)

## Contributions
Contributions are very welcome! Feel free to contact me, send a PR or open an issue.

## Roadmap
Things that are missing:
- [x] Strings: Humanize
- [x] Strings: Truncate
- [x] Numbers: Roman
- [x] Numbers: Ordinal
- [x] Collection: Oxford
- [x] Numbers: Binary Suffix
- [ ] Numbers: Metric Suffix
- [ ] Date time: Difference
- [ ] Date time: Precise difference
- [ ] Translations
