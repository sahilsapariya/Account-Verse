package types

import (
	"github.com/golang-jwt/jwt/v5"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

//nolint:typecheck // jwt.RegisteredClaims is falsely marked as undefined by linter
type AccessClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

//nolint:typecheck // jwt.RegisteredClaims is falsely marked as undefined by linter
type RefreshClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
