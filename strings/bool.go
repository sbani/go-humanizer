package strings

import (
	"fmt"
)

// TrueValues is a list of words which leads to true
var TrueValues = []string{
	"yes",
	"y",
	"true",
	"on",
	"1",
}

// FalseValues is a list of words which leads to false
var FalseValues = []string{
	"no",
	"n",
	"false",
	"off",
	"0",
}

// ToBool converts a string value to boolean
func ToBool(input string) (bool, error) {
	if isInSlice(input, TrueValues) {
		return true, nil
	} else if isInSlice(input, FalseValues) {
		return false, nil
	}

	return false, fmt.Errorf("Can't convert value \"%s\" to boolean", input)
}

// Checks if value is in slice
func isInSlice(value string, slice []string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
