package auth

import (
	"blinders/packages/common"
	"context"
	"time"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/iterator"
)

type FirebaseAuthManager struct {
	// authClient *auth.Client
	client *firebase.App
}

func (m FirebaseAuthManager) Generate(user *common.User) (string, error) {
	return "", nil
}

type firestoreUser struct {
	Email string `firestore:"email"`
	UID   string `firestore:"firebaseUid"`
}

func (m FirebaseAuthManager) Verify(token string) (*common.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	authClient, err := m.client.Auth(ctx)
	if err != nil {
		return nil, err
	}

	authToken, err := authClient.VerifyIDTokenAndCheckRevoked(ctx, token)
	if err != nil {
		return nil, err
	}
	firestore, err := m.client.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	usersRef := firestore.Collection("Users").Where("firebaseUid", "==", authToken.UID).Documents(ctx)
	defer usersRef.Stop()

	userDoc, err := usersRef.Next()
	if err != nil {
		if err == iterator.Done {
			return nil, nil
		}
		return nil, err
	}
	firestoreUser := new(firestoreUser)
	if err := userDoc.DataTo(firestoreUser); err != nil {
		return nil, err
	}
	return &common.User{
		ID:    userDoc.Ref.ID,
		Email: firestoreUser.Email,
		UID:   firestoreUser.UID,
	}, nil
}

func NewFirebaseAuthManager(firebaseApp *firebase.App) (AuthManager, error) {
	manager := &FirebaseAuthManager{
		client: firebaseApp,
	}

	return manager, nil
}
