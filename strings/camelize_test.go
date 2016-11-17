package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCamelize(t *testing.T) {
	tests := map[string]string{
		"foo_bar":            "fooBar",
		"Foo_Bar":            "fooBar",
		"FOO_BAR":            "fooBar",
		"foo__bar":           "fooBar",
		"b":                  "b",
		"b_":                 "b",
		"one_two_three_four": "oneTwoThreeFour",
	}

	for input, expected := range tests {
		assert.Equal(t, expected, Camelize(input))
	}
}

func TestUnderscore(t *testing.T) {
	tests := map[string]string{
		"oneTwo":          "one_two",
		"oneTwoThreeFour": "one_two_three_four",
	}

	for input, expected := range tests {
		assert.Equal(t, expected, Underscore(input))
	}
}
