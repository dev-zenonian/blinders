package suggestion

import (
	"blinders/packages/common"
	"context"
)

type Suggester interface {
	ChatCompletion(context.Context, common.UserData, []common.Message, ...Prompter) ([]string, error)
	TextCompletion(context.Context, common.UserData, string) ([]string, error)
}
