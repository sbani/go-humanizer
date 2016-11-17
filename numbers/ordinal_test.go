package numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdinalizesNumbers(t *testing.T) {
	assert.Equal(t, "1st", Ordinalize(1))
	assert.Equal(t, "2nd", Ordinalize(2))
	assert.Equal(t, "23rd", Ordinalize(23))
	assert.Equal(t, "1002nd", Ordinalize(1002))
	assert.Equal(t, "-111th", Ordinalize(-111))
	assert.Equal(t, "5th", Ordinalize(5))
}

func TestReturnsOrdinalSuffix(t *testing.T) {
	assert.Equal(t, "st", Ordinal(1))
	assert.Equal(t, "nd", Ordinal(2))
	assert.Equal(t, "rd", Ordinal(23))
	assert.Equal(t, "nd", Ordinal(1002))
	assert.Equal(t, "th", Ordinal(-111))
	assert.Equal(t, "th", Ordinal(5))
}
