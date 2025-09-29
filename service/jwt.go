package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wolf1848/taxiportal/service/entity"
)

func (service *Services) JwtGenerateAccessToken(userId int) (string, error) {
	claims := &entity.JwtAccessClaim{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(service.Jwt.Time * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(service.Jwt.Secret))
}

func (service *Services) JwtGenerateRefreshToken(userId int) (string, error) {
	claims := &entity.JwtRefreshClaim{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(service.Jwt.Long * time.Minute)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(service.Jwt.Refresh))
}

func (service *Services) JwtValidateRefreshToken(value string) (*entity.JwtRefreshClaim, error) {
	claim := &entity.JwtRefreshClaim{}

	token, err := jwt.ParseWithClaims(
		value,
		claim,
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, entity.ErrInvalidToken
			}
			return []byte(service.Jwt.Refresh), nil
		},
	)

	if err != nil {
		return nil, entity.ErrInvalidToken
	}

	if !token.Valid {
		return nil, entity.ErrInvalidToken
	}

	return claim, nil

}

func (service *Services) JwtValidateAccessToken(value string) (*entity.JwtAccessClaim, error) {
	claim := &entity.JwtAccessClaim{}

	token, err := jwt.ParseWithClaims(
		value,
		claim,
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, entity.ErrInvalidToken
			}
			return []byte(service.Jwt.Secret), nil
		},
	)

	if err != nil {
		return nil, entity.ErrInvalidToken
	}

	if !token.Valid {
		return nil, entity.ErrInvalidToken
	}

	return claim, nil

}
