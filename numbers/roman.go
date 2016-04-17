package numbers

import (
	"errors"
	"regexp"
	"sort"
)

// MinValue is the min number which should be converted
const MinValue = 1

// RomanStringMatcher is a regex to check if string is a roman number
const RomanStringMatcher = "^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$"

// ToRoman converts a number to a roman string. Number has to be higher than MinValue
func ToRoman(number int) (string, error) {
	if number < MinValue {
		return "", errors.New("Number not convertable")
	}

	table := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	var sortedArabs []int
	for arab := range table {
		sortedArabs = append(sortedArabs, arab)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sortedArabs)))

	romanString := ""
	for number > 0 {
		for _, arab := range sortedArabs {
			if number >= arab {
				number -= arab
				romanString += table[arab]
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

	var table = map[string]int{
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

	total := 0
	for i > 0 {
		i--
		digit := table[string(input[i])]
		if i > 0 {
			previousDigit := table[string(input[(i-1)])]
			if previousDigit < digit {
				digit -= previousDigit
				i--
			}
		}
		total += digit
	}
	return total, nil
}
