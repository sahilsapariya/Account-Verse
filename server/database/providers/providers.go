package providers

import (
	"context"
	"server/graph/model"
)

type Provider interface {
	AddUser(ctx context.Context, user *model.User) (*model.User, error)
}
