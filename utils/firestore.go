package utils

import (
	"blinders/packages/auth"
	"blinders/packages/common"
	"context"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var (
	authenticater   auth.AuthManager = nil
	credentialsPath                  = os.Getenv("GOOLE_APPLICATION_CREDENTIALS")
)

func init() {
	if authenticater == nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(credentialsPath))
		if err != nil {
			panic(err)
		}
		authenticater, err = auth.NewFirebaseAuthManager(app)
		if err != nil {
			panic(err)
		}
	}
}

func VerifyFirestoreToken(token string) (*common.User, error) {
	return authenticater.Verify(token)
}
