package post

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wolf1848/taxiportal/api/response"
	"github.com/wolf1848/taxiportal/service/jwt/entity"
)

func Handler() echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user")
		if user == nil {
			return response.InvalidAuthorize(c)
		}

		claims, ok := user.(*entity.AccessClaim)
		if !ok {
			return response.InvalidAuthorize(c)
		}

		return c.JSON(http.StatusOK, claims)
	}
}
