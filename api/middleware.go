package api

import (
	"github.com/labstack/echo/v4"
	"github.com/wolf1848/taxiportal/service/entity"
)

type JwtMiddlewareService interface {
	JwtValidateAccessToken(value string) (*entity.JwtAccessClaim, error)
}

func JWTAuthMiddleware(service JwtMiddlewareService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return invalidAuthorizeResponse(c)
			}

			claim, err := service.JwtValidateAccessToken(token)
			if err != nil {
				return invalidAuthorizeResponse(c)
			}

			c.Set("user", claim)

			return next(c)
		}
	}
}
