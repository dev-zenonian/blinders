package message

import (
	"blinders/packages/common"
	"context"
)

type Manager interface {
	/*
		Read latest messages from chat room rid
	*/
	GetMessagesOfRoom(ctx context.Context, rid string, offset int, limit int) ([]common.Message, error)

	GetRoom(ctx context.Context, rid string) (common.ChatRoom, error)
}
