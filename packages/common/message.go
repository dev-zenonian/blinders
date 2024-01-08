package common

import (
	"fmt"
	"strings"
	"time"
)

type Message struct {
	FromID    string `json:"fromID"`
	ToID      string `json:"toID"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestampt"` // Unix timestamp
}

func (m Message) String() string {
	builder := new(strings.Builder)
	builder.WriteString("[")
	fmt.Fprintf(builder, "From: %s, ", m.FromID)
	fmt.Fprintf(builder, "To: %s, ", m.ToID)
	t := time.Unix(m.Timestamp, 0)
	fmt.Fprintf(builder, "At: %s, ", t.Format(DefaultTimeFormat))
	fmt.Fprintf(builder, "Content: %s", m.Content)
	builder.WriteString("]")
	return builder.String()
}

func NewMessage(fromID string, toID string, content string) *Message {
	return &Message{
		FromID:    fromID,
		ToID:      toID,
		Content:   content,
		Timestamp: time.Now().Unix(),
	}
}
