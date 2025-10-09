package resolvers

import (
	"context"
	"errors"
	"server/database"
	"server/graph/model"
	"server/utils"
)

func SigninResolver(ctx context.Context, params model.SignInInput) (*model.AuthResponse, error) {
	if params.Email == "" {
		return nil, errors.New("email is required")
	}
	if params.Password == "" {
		return nil, errors.New("password is required")
	}

	user, err := database.Provider.GetUserByEmail(ctx, params.Email)

	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user is not registered")
	}

	if user.Password != params.Password {
		return nil, errors.New("invalid credentials")
	}
	
	tokenDetails, err := utils.GenerateTokenPair(user.ID)
	if err != nil {
		return nil, err
	}

	// Return a proper AuthResponse
	res := &model.AuthResponse{
		Message:      "User created successfully",
		AccessToken:  &tokenDetails.AccessToken,
		RefreshToken: &tokenDetails.RefreshToken,
		AtExpiresIn:  &tokenDetails.AtExpires,
		RtExpiresIn:  &tokenDetails.RtExpires,
		User:         user,
	}

	return res, nil
}
