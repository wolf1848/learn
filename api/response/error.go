package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type invalidData struct {
	Message string `json:"message"`
}

var invalidAuthorize = &invalidData{
	Message: "Некорректные учетные данные",
}

func InvalidAuthorize(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, invalidAuthorize)
}

var invalidJson = &invalidData{
	Message: "Некорректный JSON",
}

func InvalidJson(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, invalidJson)
}

var internalServer = &invalidData{
	Message: "Сервер временно недоступен, попробуйте позже",
}

func InternalServerError(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, internalServer)
}

func InvalidDataRequest(c echo.Context, response any) error {
	return c.JSON(http.StatusUnprocessableEntity, response)
}
