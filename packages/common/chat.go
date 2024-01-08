package common

const (
	Individual ChatType = "individual"
	Group      ChatType = "group"
)

func (t ChatType) String() string {
	return string(t)
}

type (
	ChatType string
	ChatRoom struct {
		Type    ChatType `json:"type"`
		UserIDs []string `json:"uids"`
	}
)
