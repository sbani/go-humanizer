# Humanizer
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.md) ![Test Workflow](https://github.com/sbani/go-humanizer/actions/workflows/test.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/sbani/go-humanizer)](https://goreportcard.com/report/github.com/sbani/go-humanizer) [![codecov](https://codecov.io/gh/sbani/go-humanizer/branch/master/graph/badge.svg)](https://codecov.io/gh/sbani/go-humanizer) [![GoDoc](https://godoc.org/github.com/sbani/go-humanizer?status.svg)](https://godoc.org/github.com/sbani/go-humanizer)

Humanize values to make them easier to read.

## Installation
```bash
go get github.com/sbani/go-humanizer
```

## Usage

### Strings

#### Humanize
```go
import "github.com/sbani/go-humanizer/strings"

Humanize("news_count", true) // "News count"
Humanize("User", false) // "user"
Humanize("news_id", true) // "News"
```

#### Truncate
Truncate string but never cut within a word.
```go
import "github.com/sbani/go-humanizer/strings"

textShort := "Short text"
Truncate(textShort, 1) // Short
Truncate(textShort, 6) // Short
Truncate(textShort, 7) // Short text
```

#### Bool
```go
import "github.com/sbani/go-humanizer/strings"

ToBool(textShort, "no") // false
ToBool(textShort, "false") // false
ToBool(textShort, "yes") // true
```
### Time

#### Difference
```go
import "github.com/sbani/go-humanizer/time"

baseTime := time.Date(2016, 11, 3, 13, 0, 0, 0, time.UTC)
Difference(baseTime, baseTime) // "just now"
Difference(baseTime, baseTime.Add(5*time.Second)) // "5 seconds from now"
Difference(baseTime, baseTime.Add(-61*time.Second)) // "1 minute ago"
Difference(baseTime, baseTime.Add(-(15*time.Minute+3*time.Nanosecond))) // "15 minutes ago"
Difference(baseTime, baseTime.Add(2*time.Hour+3*time.Minute)) // "2 hours from now"
Difference(baseTime, baseTime.Add(-(25*time.Hour))) // "1 day ago"
Difference(baseTime, baseTime.Add(14*24*time.Hour)) // "2 weeks from now"
Difference(baseTime, baseTime.Add(-31*24*time.Hour)) // "1 month ago"
Difference(baseTime, baseTime.Add(366*24*time.Hour)) // "1 year from now"
```

### Units

#### Binary Suffix
```go
import "github.com/sbani/go-humanizer/units"

s := BinarySuffix(0) // "0 bytes"
s := BinarySuffix(1536) // "1.5 kB"
s := BinarySuffix(1048576 * 5) // "5.00 MB"
s := BinarySuffix(1073741824 * 2) // "2.00 GB"
```
### Numbers

#### Ordinalize
```go
import "github.com/sbani/go-humanizer/numbers"

Ordinalize(0) // "0th"
Ordinalize(1) // "1st"
Ordinalize(2) // "2nd"
Ordinalize(23) // "23rd"
Ordinalize(1002) // "1002nd"
Ordinalize(-111) // "-111th"
```

#### Ordinal
```go
import "github.com/sbani/go-humanizer/numbers"

Ordinal(0) // "th"
Ordinal(1) // "st"
Ordinal(2) // "nd"
Ordinal(23) // "rd"
Ordinal(1002) // "nd"
Ordinal(-111) // "th"
```
#### Roman
```go
import "github.com/sbani/go-humanizer/numbers"

s, err := ToRoman(1) // "I"
s, err := ToRomanToRoman(5) // "V"
s, err := ToRomanToRoman(1300) // "MCCC"

i, err := FromRoman("MMMCMXCIX") // 3999
i, err := FromRoman("V") // 5
i, err := FromRoman("CXXV") // 125
```

### Collection

#### Oxford
```go
import "github.com/sbani/go-humanizer/collection"

Oxford([]string{"Albert"}, -1) // "Albert"
Oxford([]string{"Albert", "Norbert"}, -1) // "Albert and Norbert"
Oxford([]string{"Albert", "Norbert", "Michael", "Kevin"}, -1) // "Albert, Norbert, Michael, and Kevin"
Oxford([]string{"Albert", "Norbert", "Michael", "Kevin"}, 2)) // Albert, Norbert, and 2 more
```

## License
MIT License. See LICENSE file for more informations.

## Credits
Special thank goes to [PHP Humanizer](https://github.com/coduo/php-humanizer).

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
- [x] Numbers: Metric Suffix
- [x] Date time: Difference
- [ ] Date time: Precise difference
- [ ] Translations
