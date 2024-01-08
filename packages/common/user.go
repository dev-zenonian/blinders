package common

import (
	"fmt"
	"strings"
)

type User struct {
	ID    string `json:"userID"`
	Email string `json:"userEmail"`
	UID   string `json:"firebaseUID"`
}

func (u User) String() string {
	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("UserID: %s", u.ID))
	str.WriteString(fmt.Sprintf("Email: %s", u.Email))
	return str.String()
}

type UserData struct {
	UserID   string   `json:"userID"`
	Native   Language `json:"nativeLanguage"`
	Learning Language `json:"learningLanguage"`
}

func (d UserData) String() string {
	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("Native language: %s\n", d.Native))
	str.WriteString(fmt.Sprintf("Learning language: %s\n", d.Native))
	return str.String()
}
