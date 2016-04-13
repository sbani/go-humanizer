package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHumanize(t *testing.T) {
	assert.Equal(t, "News count", Humanize("news_count", true))
	assert.Equal(t, "user", Humanize("User", false))
	assert.Equal(t, "News", Humanize("news_id", true))
}
