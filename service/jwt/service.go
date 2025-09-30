package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wolf1848/taxiportal/model"
	"github.com/wolf1848/taxiportal/service/jwt/entity"
)

type Service struct {
	config *model.AppApiConfig
}

func NewService(cfg *model.AppApiConfig) *Service {
	return &Service{
		config: cfg,
	}
}

func (s *Service) GetAccessToken(userId int) (string, error) {
	claims := &entity.AccessClaim{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.config.Jwt.Time * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.config.Jwt.Secret))
}

func (s *Service) GetRefreshToken(userId int) (string, error) {
	claims := &entity.RefreshClaim{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.config.Jwt.Long * time.Minute)),
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(s.config.Jwt.Refresh))
}

func (s *Service) ValidateRefreshToken(value string) (*entity.RefreshClaim, error) {
	claim := &entity.RefreshClaim{}

	token, err := jwt.ParseWithClaims(
		value,
		claim,
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, entity.ErrInvalidToken
			}
			return []byte(s.config.Jwt.Refresh), nil
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

func (s *Service) ValidateAccessToken(value string) (*entity.AccessClaim, error) {
	claim := &entity.AccessClaim{}

	token, err := jwt.ParseWithClaims(
		value,
		claim,
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, entity.ErrInvalidToken
			}
			return []byte(s.config.Jwt.Secret), nil
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
