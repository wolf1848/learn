package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wolf1848/taxiportal/api/dto"
	"github.com/wolf1848/taxiportal/service/entity"
)

type RefreshService interface {
	RefreshToken(token string) (*entity.AuthorizeOutput, error)
}

func refresh(service RefreshService) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req *dto.RefreshRequest

		if err := c.Bind(&req); err != nil {
			return invalidJsonResponse(c)
		}

		output, err := service.RefreshToken(req.RefreshToken)

		if err != nil {
			if errors.Is(err, entity.ErrInvalidToken) {
				return invalidAuthorizeResponse(c)
			}
			return internalServerResponse(c)
		}

		response := &dto.AuthorizeResponse{
			ID:           output.ID,
			Name:         output.Name,
			Email:        output.Email,
			Token:        output.Token,
			RefreshToken: output.RefreshToken,
		}

		return c.JSON(http.StatusOK, response)
	}
}
