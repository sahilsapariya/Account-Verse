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

func (p *provider) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user *model.User
	result := p.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
