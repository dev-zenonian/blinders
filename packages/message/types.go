package message

import (
	"blinders/packages/common"
	"time"
)

type FirestoreMessage struct {
	FromID    string    `firestore:"senderId"`
	ToID      string    `firestore:"roomId"`
	Timestamp time.Time `firestore:"time"`
	Content   string    `firestore:"content"`
}

func (f FirestoreMessage) ToCommonMessage() common.Message {
	return common.Message{
		FromID:    f.FromID,
		ToID:      f.ToID,
		Timestamp: f.Timestamp.Unix(),
		Content:   f.Content,
	}
}

type FirestoreChatRoom struct {
	Type    string   `firestore:"type"`
	UserIDs []string `firestore:"members"`
}

func (r FirestoreChatRoom) ToCommonRoom() common.ChatRoom {
	return common.ChatRoom{
		Type:    common.ChatType(r.Type),
		UserIDs: r.UserIDs,
	}
}
