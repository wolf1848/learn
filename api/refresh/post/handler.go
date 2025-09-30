package post

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wolf1848/taxiportal/api/refresh/post/dto"
	"github.com/wolf1848/taxiportal/api/response"
	"github.com/wolf1848/taxiportal/service/authorize/entity"
)

type Service interface {
	RefreshAuthorize(string) (*entity.Output, error)
}

func Handler(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req *dto.Request

		if err := c.Bind(&req); err != nil {
			return response.InvalidJson(c)
		}

		output, err := service.RefreshAuthorize(req.RefreshToken)

		if err != nil {
			response.InternalServerError(c)
		}

		response := &dto.Response{
			ID:           output.ID,
			Name:         output.Name,
			Email:        output.Email,
			Token:        output.Token,
			RefreshToken: output.RefreshToken,
		}

		return c.JSON(http.StatusOK, response)
	}
}
