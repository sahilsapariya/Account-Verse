package resolvers

import (
	"context"
	"errors"
	"server/database"
	"server/graph/model"
	logs "server/logs"
	"server/utils"

	"github.com/google/uuid"
)

func SignupResolver(ctx context.Context, params model.SignUpInput) (*model.AuthResponse, error) {
	logger := logs.InitLog("debug")

	if params.Email == "" {
		return nil, errors.New("email is required")
	}

	// check if user exist or not
	result, err := database.Provider.GetUserByEmail(ctx, params.Email)
	if err != nil {
		return nil, err
	}
	if result != nil {
		return &model.AuthResponse{
			Message: "User already exist",
		}, nil
	}

	if params.Password == "" {
		return nil, errors.New("password is required")
	}
	if params.ConfirmPassword == "" {
		return nil, errors.New("confirm password is required")
	}
	if params.Password != params.ConfirmPassword {
		return nil, errors.New("password and confirm password do not match")
	}
	if params.SignupMethod == nil {
		defaultMethod := model.SignupMethodBasicAuth
		params.SignupMethod = &defaultMethod
	}
	if params.Roles == nil {
		defaultRole := []string{"user"}
		params.Roles = defaultRole
	}

	userID := uuid.New().String()

	user := &model.User{
		ID:           userID,
		Username:     params.Username,
		Roles:        params.Roles,
		SignupMethod: *params.SignupMethod,
		Email:        params.Email,
		Password:     params.Password,
		GivenName:    params.GivenName,
		MiddleName:   params.MiddleName,
		FamilyName:   params.FamilyName,
		Gender:       params.Gender,
		DateOfBirth:  params.DateOfBirth,
	}

	_, err = database.Provider.AddUser(ctx, user)
	if err != nil {
		logger.Debug("Failed to create user: ", err)
		return nil, err
	}

	tokenDetails, err := utils.GenerateTokenPair(userID)
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
	}

	return res, nil
}
