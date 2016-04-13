package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOxford(t *testing.T) {
	assert.Equal(t, "", Oxford([]string{}, -1))

	var testCollection = []string{
		"Albert",
		"Norbert",
		"Michael",
		"Kevin",
	}

	assert.Equal(t, "Albert", Oxford(testCollection[:1], -1))
	assert.Equal(t, "Albert and Norbert", Oxford(testCollection[:2], -1))
	assert.Equal(t, "Albert, Norbert, Michael and Kevin", Oxford(testCollection, -1))
	assert.Equal(t, "Albert, Norbert and 2 more", Oxford(testCollection, 2))
}
