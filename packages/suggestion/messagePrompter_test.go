package suggestion

import (
	"blinders/packages/common"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildPrompt(t *testing.T) {
	_, _, _, prompter := initPrompt(t)
	prompt, err := prompter.Build()
	assert.Nil(t, err)
	assert.NotEmpty(t, prompt)
}

func TestUpdatePrompt(t *testing.T) {
	sender, receiver, userContext, prompter := initPrompt(t)
	assert.Equal(t, userContext, prompter.UserData)

	newContext := newUserContext(
		&receiver,
		common.Language{},
		common.Language{},
	)
	assert.Nil(t, prompter.Update(newContext))
	assert.Equal(t, newContext, prompter.UserData)

	lenMessages := 5
	msgs := []common.Message{}
	for i := 0; i < lenMessages; i++ {
		switch i % 2 {
		case 0:
			msgs = append(msgs, *common.NewMessage(sender.ID, receiver.ID, fmt.Sprintf("test_%d", i)))
		default:
			msgs = append(msgs, *common.NewMessage(receiver.ID, sender.ID, fmt.Sprintf("test_%d", i)))
		}
	}
	assert.Nil(t, prompter.Update(msgs))
	assert.Equal(t, msgs, prompter.messages)

	assert.NotNil(t, prompter.Update(sender))
}

func initPrompt(t *testing.T) (
	sender common.User,
	receiver common.User,
	senderData common.UserData,
	prompter *MessageSuggestionPrompterere,
) {
	sender = newUser("sender", "sender@email")
	receiver = newUser("receiver", "receiver@email")
	senderData = newUserContext(
		&sender,
		common.Language{
			Lang:  common.LangVi,
			Level: common.Beginner,
		}, common.Language{
			Lang:  common.LangEn,
			Level: common.Beginner,
		})
	lenMessages := 10
	msgs := []common.Message{}
	for i := 0; i < lenMessages; i++ {
		switch i % 2 {
		case 0:
			msgs = append(msgs, *common.NewMessage(sender.ID, receiver.ID, fmt.Sprintf("msg_%d", i)))
		default:
			msgs = append(msgs, *common.NewMessage(receiver.ID, sender.ID, fmt.Sprintf("msg_%d", i)))
		}
	}

	prompter = NewMessageSuggestionPrompter()
	assert.NotNil(t, prompter)
	assert.Nil(t, prompter.Update(senderData, msgs))
	return
}

func newUserContext(user *common.User, native common.Language, language common.Language) common.UserData {
	return common.UserData{
		UserID:   user.ID,
		Native:   native,
		Learning: language,
	}
}

func newUser(id string, email string) common.User {
	return common.User{
		ID:    id,
		Email: email,
	}
}
