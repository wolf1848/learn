package api

import (
	"github.com/labstack/echo/v4"
	"github.com/wolf1848/taxiportal/api/response"
	"github.com/wolf1848/taxiportal/service/jwt/entity"
)

type JwtMiddlewareService interface {
	ValidateAccessToken(string) (*entity.AccessClaim, error)
}

func JWTAuthMiddleware(service JwtMiddlewareService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return response.InvalidAuthorize(c)
			}

			claim, err := service.ValidateAccessToken(token)
			if err != nil {
				return response.InvalidAuthorize(c)
			}

			c.Set("user", claim)

			return next(c)
		}
	}
}
