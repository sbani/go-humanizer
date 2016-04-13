package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTruncate(t *testing.T) {
	text := "Lorem ipsum dolorem si amet, lorem ipsum. Dolorem sic et nunc."
	assert.Equal(t, "Lorem", Truncate(text, 2))
	assert.Equal(t, "Lorem ipsum", Truncate(text, 10))
	assert.Equal(t, "Lorem ipsum dolorem si amet, lorem", Truncate(text, 30))
	assert.Equal(t, "Lorem", Truncate(text, 0))
	assert.Equal(t, text, Truncate(text, -2))

	textShort := "Short text"
	assert.Equal(t, "Short", Truncate(textShort, 1))
	assert.Equal(t, "Short", Truncate(textShort, 2))
	assert.Equal(t, "Short", Truncate(textShort, 3))
	assert.Equal(t, "Short", Truncate(textShort, 4))
	assert.Equal(t, "Short", Truncate(textShort, 5))
	assert.Equal(t, "Short", Truncate(textShort, 6))
	assert.Equal(t, "Short text", Truncate(textShort, 7))
	assert.Equal(t, "Short text", Truncate(textShort, 8))
	assert.Equal(t, "Short text", Truncate(textShort, 9))
	assert.Equal(t, "Short text", Truncate(textShort, 10))
	assert.Equal(t, "Short text", Truncate(textShort, 100))
}
