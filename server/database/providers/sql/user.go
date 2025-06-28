package sql

import (
	"context"
	"server/graph/model"

	"github.com/google/uuid"
)

func (p *provider) AddUser(ctx context.Context, user *model.User) (*model.User, error) {
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	res := p.db.Create(&user)

	if res.Error != nil {
		return user, res.Error
	}

	return user, nil
}
