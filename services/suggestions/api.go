package suggestion

import (
	"blinders/packages/common"
	"blinders/packages/user"
	"blinders/utils"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/gofiber/fiber/v2"
)

type SuggestionPayload struct {
	Text   string `json:"text"`
	UserID string `json:"userID"`
}

func (s *Service) HandleTextSuggestion() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")
		if token == "" {
			return ctx.Status(400).JSON(fiber.Map{
				"error": "suggestion: token in Authorization header not found",
			})
		}
		usr, err := utils.VerifyFirestoreToken(token)
		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": fmt.Sprintf("suggestion: cannot verify user with given token (%s)", token),
			})
		}

		req := new(SuggestionPayload)
		if err := json.Unmarshal(ctx.Body(), req); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error":       err.Error(),
				"suggestions": []string{},
			})
		}
		userData, err := user.GetUserData(usr.ID)
		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error": fmt.Sprintf("suggestion: cannot get data of user, err: (%s)", err.Error()),
			})
		}

		suggestions, err := s.suggester.TextCompletion(ctx.Context(), userData, req.Text)
		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"error":       err.Error(),
				"suggestions": []string{},
			})
		}

		return ctx.Status(200).JSON(fiber.Map{
			"suggestions": suggestions,
		})
	}
}

type ChatSuggestionPayload struct {
	UserID   string          `json:"userID"`
	Messages []ClientMessage `json:"messages"`
}

type ClientMessage struct {
	Timestamp any    `json:"time"`
	ID        string `json:"id"`
	Content   string `json:"content"`
	FromID    string `json:"senderId"`
	ChatID    string `json:"roomId"`
}

func (m ClientMessage) ToCommonMessage() common.Message {
	var Timestamp int64
	switch timestamp := m.Timestamp.(type) {
	case int:
		Timestamp = int64(timestamp)
	case string:
		// expect date time as string type, "Tue Dec 05 2023 12:35:04 GMT+0700"
		layout := "Mon Jan 02 2006 15:04:05 GMT-0700"
		t, err := time.Parse(layout, timestamp)
		if err != nil {
			panic(fmt.Sprintf("clienmessage: given time (%s) cannot parse with layout (%s)", timestamp, layout))
		}
		Timestamp = t.Unix()
	default:
		panic(fmt.Sprintf("clienmessage: unknow timestamp type (%s)", reflect.TypeOf(m.Timestamp).String()))
	}

	return common.Message{
		FromID:    m.FromID,
		ToID:      m.ChatID,
		Content:   m.Content,
		Timestamp: Timestamp,
	}
}

func (s *Service) HandleChatSuggestion() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req := new(ChatSuggestionPayload)
		if err := json.Unmarshal(ctx.Body(), req); err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"suggestions": []string{},
			})
		}

		// should communicate with user service
		userData, err := user.GetUserData(req.UserID)
		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"suggestions": []string{},
			})
		}

		msgs := []common.Message{}
		for _, msg := range req.Messages {
			msgs = append(msgs, msg.ToCommonMessage())
		}

		suggestions, err := s.suggester.ChatCompletion(ctx.Context(), userData, msgs)
		if err != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"suggestions": []string{},
			})
		}

		return ctx.Status(200).JSON(fiber.Map{
			"suggestions": suggestions,
		})
	}
}
