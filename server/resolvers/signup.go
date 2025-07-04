package resolvers

import (
	"context"
	"server/database"
	"server/graph/model"
	logs "server/logs"

	"github.com/google/uuid"
)

func SignupResolver(ctx context.Context, params model.SignUpInput) (*model.AuthResponse, error) {
	logger := logs.InitLog("debug")

	userID := uuid.New().String()

	user := &model.User{
		ID:    userID,
		Name:  params.Name,
		Email: params.Email,
	}

	user, err := database.Provider.AddUser(ctx, user)
	if err != nil {
		logger.Debug("Failed to create user: ", err)
		return nil, err
	}

	// Return a proper AuthResponse
	res := &model.AuthResponse{
		Message: "User created successfully",
	}

	return res, nil
}
