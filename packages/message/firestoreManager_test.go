package message

import (
	"context"
	"fmt"
	"testing"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/stretchr/testify/assert"
)

func TestFirestoreAdapter(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	manager := initManager(t, ctx)

	var (
		roomID = "Hp8ugceFOrycOGPxC7C9"
		limit  = 5
		offset = 0
	)

	uids, err := manager.GetUsersIdOfRoom(ctx, roomID)
	assert.Nil(t, err)
	fmt.Println(uids)

	msgs, err := manager.GetMessagesOfRoom(ctx, roomID, offset, limit)
	assert.Nil(t, err)
	assert.NotNil(t, msgs)

	for _, msg := range msgs {
		fmt.Printf("msg: %v\n", msg)
	}
}

func initManager(t *testing.T, ctx context.Context) *FirestoreManager {
	app, err := firebase.NewApp(ctx, nil)
	assert.Nil(t, err)
	assert.NotNil(t, app)

	client, err := app.Firestore(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, client)

	m := NewFirestoreManager(client)
	assert.NotNil(t, m)

	assert.Nil(t, m.Ping(ctx))
	return m
}
