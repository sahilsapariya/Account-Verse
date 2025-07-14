package utils

import (
	"encoding/json"
	"errors"
	"os"
	"server/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessSecret  = []byte(os.Getenv("ACCESS_SECRET"))  // set in env
	refreshSecret = []byte(os.Getenv("REFRESH_SECRET")) // set in env
)

func GenerateTokenPair(userID string) (*types.TokenDetails, error) {
	atExpires := time.Now().Add(15 * time.Minute).Unix()
	rtExpires := time.Now().Add(7 * 24 * time.Hour).Unix()

	accessToken, err := createToken(userID, atExpires, accessSecret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := createToken(userID, rtExpires, refreshSecret)
	if err != nil {
		return nil, err
	}

	return &types.TokenDetails{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AtExpires:    atExpires,
		RtExpires:    rtExpires,
	}, nil
}

func createToken(userID string, exp int64, secret []byte) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Unix(exp, 0)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ValidateAccessToken(tokenStr string) (*types.AccessClaims, error) {
	return validateToken[types.AccessClaims](tokenStr, accessSecret)
}

func ValidateRefreshToken(tokenStr string) (*types.RefreshClaims, error) {
	return validateToken[types.RefreshClaims](tokenStr, refreshSecret)
}

func validateToken[T any](tokenStr string, secret []byte) (*T, error) {
	token, err := jwt.ParseWithClaims(tokenStr, new(jwt.MapClaims), func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return nil, errors.New("token claims type mismatch")
	}

	// Unmarshal jwt.MapClaims into the expected type T
	var t T
	bytes, err := json.Marshal(claims)
	if err != nil {
		return nil, errors.New("failed to marshal claims")
	}
	if err := json.Unmarshal(bytes, &t); err != nil {
		return nil, errors.New("failed to unmarshal claims into target type")
	}
	return &t, nil
}
