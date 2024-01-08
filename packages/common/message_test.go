package common

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	var (
		fromID  = "sender"
		toID    = "receiver"
		content = "Hello, World!"
	)

	msg := NewMessage(fromID, toID, content)
	assert.NotNil(t, msg)
	assert.Equal(t, fromID, msg.FromID)
	assert.Equal(t, toID, msg.ToID)
	assert.Equal(t, content, msg.Content)
	fmt.Printf("msg: %v\n", msg)
}
