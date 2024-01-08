package suggestion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEmbedString(t *testing.T) {
	assert.NotNil(t, randSource)

	embedStr, err := randomEmbed()
	assert.Nil(t, err)
	assert.NotEmpty(t, embedStr)
}

func TestGetEmbedWithIndex(t *testing.T) {
	assert.NotNil(t, randSource)

	for i, embed := range messageSuggestionEmbed {
		e, err := randomEmbed(i)
		assert.Nil(t, err)
		assert.Equal(t, embed, e)
	}

	embed, err := randomEmbed(len(messageSuggestionEmbed) + 5)
	assert.NotNil(t, err)
	assert.Empty(t, embed)
}
