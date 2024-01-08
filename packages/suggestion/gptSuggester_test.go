package suggestion

import (
	"blinders/packages/common"
	"blinders/packages/message"
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/assert"
)

var authToken = os.Getenv("OPENAI_API_KEY")

func TestInit(t *testing.T) {
	suggester := initSuggester(t)
	fmt.Println(suggester.GPTSuggesterOptions)
}

func TestTextCompletion(t *testing.T) {
	suggester := initSuggester(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	prompt := "Just reply 'hello, world!'"
	suggetions, err := suggester.TextCompletion(ctx, common.UserData{}, prompt)
	assert.Nil(t, err)
	assert.Equal(t, suggester.nText, len(suggetions))

	fmt.Println(suggetions)
}

// This test will use the OpenAI API, it may be charged, consider uncomment for testing
func TestSuggest(t *testing.T) {
	suggester := initSuggester(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	sender := newUser("user1", "sample1")
	receiver := newUser("user2", "sample2")
	userContext := newUserContext(
		&sender,
		common.Language{
			Lang:  common.LangVi,
			Level: common.Advanced,
		},
		common.Language{
			Lang:  common.LangEn,
			Level: common.Beginner,
		},
	)
	msgs := []common.Message{
		*common.NewMessage(sender.ID, receiver.ID, "Hello, how are you?"),
		*common.NewMessage(receiver.ID, sender.ID, "Fine, how about you?"),
		*common.NewMessage(sender.ID, receiver.ID, "Too. Did you come to the class yesterday?"),
		*common.NewMessage(receiver.ID, sender.ID, "Yes, yesterday the teacher gave the students some homework."),
	}

	suggestions, err := suggester.ChatCompletion(ctx, userContext, msgs)
	assert.Nil(t, err)
	assert.NotNil(t, suggestions)
	assert.Equal(t, suggester.nChat, len(suggestions))

	for _, suggestion := range suggestions {
		fmt.Printf("suggestion: %v\n", suggestion)
	}
}

func TestIntegrateWithMessagePackage(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var (
		user        = newUser("RyuIyfR24uo9l8DCTGjS", "minhdat15012002@gmail.com")
		userContext = newUserContext(
			&user,
			common.Language{
				Lang:  common.LangVi,
				Level: common.Advanced,
			}, common.Language{
				Lang:  common.LangEn,
				Level: common.Beginner,
			})
		roomID = "Hp8ugceFOrycOGPxC7C9"
	)
	fireStoreManager := initFirestoreManager(t, ctx)
	room, err := fireStoreManager.GetRoom(ctx, roomID)
	assert.Nil(t, err)
	assert.NotEmpty(t, room.Type)
	assert.NotEmpty(t, room.UserIDs)

	assert.Contains(t, room.UserIDs, user.ID)
	msgs, err := fireStoreManager.GetMessagesOfRoom(ctx, roomID, 0, 3)
	assert.Nil(t, err)
	assert.NotNil(t, msgs)

	suggester := initSuggester(t)
	suggestions, err := suggester.ChatCompletion(ctx, userContext, msgs)
	assert.Nil(t, err)
	assert.NotNil(t, suggestions)

	for _, suggestion := range suggestions {
		fmt.Printf("suggestion: %v\n", suggestion)
	}
}

func initSuggester(t *testing.T) *GPTSuggester {
	client := openai.NewClient(authToken)
	opts := []Option{
		WithNChat(1),
	}
	suggester, err := NewGPTSuggester(client, opts...)
	assert.Nil(t, err)
	assert.NotNil(t, suggester)
	assert.Equal(t, 1, suggester.nChat)
	return suggester
}

func initFirestoreManager(t *testing.T, ctx context.Context) *message.FirestoreManager {
	app, err := firebase.NewApp(ctx, nil)
	assert.Nil(t, err)
	assert.NotNil(t, app)

	client, err := app.Firestore(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, client)

	m := message.NewFirestoreManager(client)
	assert.NotNil(t, m)

	assert.Nil(t, m.Ping(ctx))
	return m
}
