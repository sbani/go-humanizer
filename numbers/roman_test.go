package numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var workingTests = map[string]int{
	"I":          1,
	"XC":         90,
	"DCCCXCIX":   899,
	"IX":         9,
	"CXXV":       125,
	"MMMMCMXCIX": 4999,
}

func TestToRoman(t *testing.T) {
	for expected, input := range workingTests {
		output, err := ToRoman(input)

		if err == nil {
			assert.Equal(t, expected, output)
		} else {
			assert.Fail(t, "An error was given unexpectatly")
		}
	}
}

func TestToRomanError(t *testing.T) {
	tests := []int{
		-1,
		0,
	}

	for _, input := range tests {
		_, err := ToRoman(input)
		assert.EqualError(t, err, "Number not convertable")
	}
}

func TestFromRoman(t *testing.T) {
	for input, expected := range workingTests {
		output, err := FromRoman(input)

		if err == nil {
			assert.Equal(t, expected, output)
		} else {
			assert.Fail(t, "An error was given unexpectatly")
		}
	}
}

func TestFromRomanError(t *testing.T) {
	tests := []string{
		"foobar",
		"",
	}

	for _, input := range tests {
		_, err := FromRoman(input)
		assert.EqualError(t, err, "String not convertable")
	}
}
