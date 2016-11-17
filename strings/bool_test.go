package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBool(t *testing.T) {
	tests := map[string]bool{
		"yes":   true,
		"no":    false,
		"y":     true,
		"n":     false,
		"true":  true,
		"false": false,
		"on":    true,
		"off":   false,
		"1":     true,
		"0":     false,
	}

	for input, expected := range tests {
		out, err := ToBool(input)
		assert.Equal(t, expected, out)
		assert.Nil(t, err)
	}
}

func TestToBoolFail(t *testing.T) {
	out, err := ToBool("ILikeApple")
	assert.NotNil(t, err)
	assert.False(t, out)
}
