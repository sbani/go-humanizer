package numbers

import (
	"errors"
	"regexp"
	"sort"
)

// MinValue is the min number which should be converted
const MinValue = 1

// MaxValue is the maxnumber which should be converted
const MaxValue = 3999

// RomanStringMatcher is a regex to check if string is a roman number
const RomanStringMatcher = "^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$"

// StringMap is the map of roman (string) to arab number with string as key
var StringMap = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

// StringMapReversed is the map of roman (string) to arab number with integer as key
var StringMapReversed = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	0:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

// Get reversed sorted slice from StringMapReversed map
func getReverseIntSlice() []int {
	intSlice := make([]int, 0, len(StringMapReversed))
	for _, num := range StringMap {
		intSlice = append(intSlice, num)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))

	return intSlice
}

// ToRoman converts a number to a roman string. Number has to be between 1 and 3999.
// represented by MinValue and MaxValue
func ToRoman(number int) (string, error) {
	if number < MinValue || number > MaxValue {
		return "", errors.New("Number not convertable")
	}
	romanString := ""
	intSlice := getReverseIntSlice()
	for number > 0 {
		for _, num := range intSlice {
			if number >= num {
				romanString += StringMapReversed[num]
				number -= num
				break
			}
		}
	}
	return romanString, nil
}

// FromRoman converts a roman string to an integer. Returns an error if string is not a roman number.
func FromRoman(input string) (int, error) {
	i := len(input)

	if i == 0 || !regexp.MustCompile(RomanStringMatcher).MatchString(input) {
		return 0, errors.New("String not convertable")
	}

	total := 0
	for i > 0 {
		i--
		digit := StringMap[string(input[i])]
		if i > 0 {
			previousDigit := StringMap[string(input[(i-1)])]
			if previousDigit < digit {
				digit -= previousDigit
				i--
			}
		}
		total += digit
	}
	return total, nil
}
