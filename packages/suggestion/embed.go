package suggestion

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	randSource             *rand.Rand = nil
	messageSuggestionEmbed            = []string{
		`Sender information:
    Language: %s
    Level: %s
Recent messages: [
%s
]
Prompt:
    You are the sender, and you have to reply the latest message.
    Ensure that the message is contextually relevant, considerate, and aligned with the sender's language proficiency.
    Aim for a response that flows seamlessly within the ongoing conversation.
    Just return the text.`,
	}
)

func randomEmbed(i ...int) (string, error) {
	if len(i) == 1 {
		if 0 > i[0] || i[0] >= len(messageSuggestionEmbed) {
			return "", fmt.Errorf("embed: index out of length, got (%d)", i[0])
		}
		return messageSuggestionEmbed[i[0]], nil
	}

	if randSource == nil {
		randSource = rand.New(rand.NewSource(time.Now().Unix()))
	}
	return messageSuggestionEmbed[randSource.Intn(len(messageSuggestionEmbed))], nil
}
