package post

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wolf1848/taxiportal/api/authorize/post/dto"
	"github.com/wolf1848/taxiportal/api/response"
	"github.com/wolf1848/taxiportal/service/authorize/entity"
	serviceErrors "github.com/wolf1848/taxiportal/service/errors"
)

type Service interface {
	Authorize(*entity.Input) (*entity.Output, error)
}

func Handler(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req *dto.Request

		if err := c.Bind(&req); err != nil {
			return response.InvalidJson(c)
		}

		if errs := validate(req); len(errs) > 0 {
			return response.InvalidDataRequest(c, errs)
		}

		output, err := service.Authorize(&entity.Input{
			Email: req.Email,
			Pwd:   req.Pwd,
		})

		if err != nil {
			if errors.Is(err, serviceErrors.ErrAuthorized) {
				return response.InvalidAuthorize(c)
			}
			return response.InternalServerError(c)
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
