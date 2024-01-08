package suggestion

import (
	"blinders/packages/common"
	"errors"
	"fmt"
	"strings"
)

type MessageSuggestionPrompter struct {
	embed    string // embed string that could be use to embed user's context and messages to make complete prompt
	UserData common.UserData
	messages []common.Message
}

func (p MessageSuggestionPrompter) Build() (string, error) {
	msgs := []string{}
	for _, msg := range p.messages {
		switch msg.FromID {
		case p.UserData.UserID:
			msgs = append(msgs, fmt.Sprintf("\tsender: %s", msg.Content))
		default:
			msgs = append(msgs, fmt.Sprintf("\treceiver: %s", msg.Content))
		}
	}
	return fmt.Sprintf(p.embed, p.UserData.Learning.Lang, p.UserData.Learning.Level, strings.Join(msgs, "\n")), nil
}

func (p *MessageSuggestionPrompter) Update(objs ...any) error {
	for _, obj := range objs {
		switch doc := obj.(type) {
		case common.UserData:
			p.UserData = doc
		case []common.Message:
			p.messages = doc
		default:
			return errors.New("messageSuggestionPrompter: expected(common.UserContext, []common.Message) got unknown")
		}
	}
	return nil
}

func NewMessageSuggestionPrompter() *MessageSuggestionPrompter {
	embed, _ := randomEmbed()

	return &MessageSuggestionPrompter{
		embed: embed,
	}
}
