package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wolf1848/taxiportal/service/entity"
)

func userList() echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user")
		if user == nil {
			return invalidAuthorizeResponse(c)
		}

		claims, ok := user.(*entity.JwtAccessClaim)
		if !ok {
			return invalidAuthorizeResponse(c)
		}

		return c.JSON(http.StatusOK, claims)
	}
}
