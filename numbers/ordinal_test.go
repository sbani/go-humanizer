package numbers

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestOrdinalizesNumbers(t *testing.T) {
    assert.Equal(t, "1st", Odinalize(1))
    assert.Equal(t, "2nd", Odinalize(2))
    assert.Equal(t, "23rd", Odinalize(23))
    assert.Equal(t, "1002nd", Odinalize(1002))
    assert.Equal(t, "-111th", Odinalize(-111))
}

func TestReturnsOrdinalSuffix(t *testing.T) {
    assert.Equal(t, "st", Ordinal(1))
    assert.Equal(t, "nd", Ordinal(2))
    assert.Equal(t, "rd", Ordinal(23))
    assert.Equal(t, "nd", Ordinal(1002))
    assert.Equal(t, "th", Ordinal(-111))
}
