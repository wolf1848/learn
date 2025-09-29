package entity

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtAccessClaim struct {
	UserID int
	jwt.RegisteredClaims
}

type JwtRefreshClaim struct {
	UserID int
	jwt.RegisteredClaims
}
