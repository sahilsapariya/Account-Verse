package resolvers

import (
	"context"
	"errors"
	"server/graph/model"
	"server/utils"
)

func ValidateJwtTokenResolver(ctx context.Context, params model.ValidateJWTTokenInput) (*model.ValidateJWTTokenResponse, error) {
	if params.TokenType == "" {
		return nil, errors.New("token type is required")
	}
	if params.Token == "" {
		return nil, errors.New("token is required")
	}

	var (
		claimsMap map[string]interface{}
		err       error
	)

	switch params.TokenType {
		case model.TokenTypeAccessToken:
			claims, e := utils.ValidateAccessToken(params.Token)
			if e != nil {
				err = e
				break
			}
			claimsMap, err = utils.StructToMap(claims)

		case model.TokenTypeRefreshToken:
			claims, e := utils.ValidateRefreshToken(params.Token)
			if e != nil {
				err = e
				break
			}
			claimsMap, err = utils.StructToMap(claims)

		case model.TokenTypeIDToken:
			return nil, errors.New("ID token validation is not supported yet")

		default:
			return nil, errors.New("invalid token type")
	}

	if err != nil {
		return nil, err
	}

	return &model.ValidateJWTTokenResponse{
		IsValid:  true,
		Claims: claimsMap,
	}, nil
}
