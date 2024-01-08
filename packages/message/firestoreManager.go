package message

import (
	"blinders/packages/common"
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type FirestoreManager struct {
	client *firestore.Client
}

func (m FirestoreManager) GetMessagesOfRoom(ctx context.Context, rid string, offset int, limit int) ([]common.Message, error) {
	var (
		msgsPath = fmt.Sprintf("ChatRooms/%s/Messages", rid)
		msgsRef  = m.client.Collection(msgsPath)
		msgs     = []common.Message{}
	)
	query := msgsRef.Offset(offset).Limit(limit).Documents(ctx)
	defer query.Stop()

	docs, err := query.GetAll()
	if err != nil {
		return msgs, err
	}

	for _, doc := range docs {
		firestoreMessage := new(FirestoreMessage)
		if err := doc.DataTo(firestoreMessage); err != nil {
			return nil, err
		}
		msgs = append(msgs, firestoreMessage.ToCommonMessage())
	}
	return msgs, nil
}

func (m FirestoreManager) GetRoom(ctx context.Context, rid string) (common.ChatRoom, error) {
	fireStoreRoom, err := m.getFirestoreRoom(ctx, rid)
	if err != nil {
		return common.ChatRoom{}, err
	}
	return fireStoreRoom.ToCommonRoom(), err
}

func (m FirestoreManager) Ping(ctx context.Context) error {
	iter := m.client.Collections(ctx)
	_, err := iter.Next()
	if err != nil {
		if err == iterator.Done {
			return nil
		}
		return err
	}
	return nil
}

func (m *FirestoreManager) getFirestoreRoom(ctx context.Context, rid string) (*FirestoreChatRoom, error) {
	roomPath := fmt.Sprintf("ChatRooms/%s", rid)
	docRef, err := m.client.Doc(roomPath).Get(ctx)
	if err != nil {
		return nil, err
	}

	firestoreChatRoom := new(FirestoreChatRoom)
	if err := docRef.DataTo(firestoreChatRoom); err != nil {
		return nil, err
	}

	return firestoreChatRoom, nil
}

func (m *FirestoreManager) GetUsersIdOfRoom(ctx context.Context, rid string) ([]string, error) {
	room, err := m.getFirestoreRoom(ctx, rid)
	if err != nil {
		return nil, err
	}
	return room.UserIDs, nil
}

func NewFirestoreManager(client *firestore.Client) *FirestoreManager {
	return &FirestoreManager{
		client: client,
	}
}
