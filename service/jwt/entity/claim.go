package entity

import (
	"github.com/golang-jwt/jwt/v5"
)

type AccessClaim struct {
	UserID int
	jwt.RegisteredClaims
}

type RefreshClaim struct {
	UserID int
	jwt.RegisteredClaims
}
