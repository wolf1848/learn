package api

import (
	"github.com/labstack/echo/v4"
)

type Services interface {
	RegisterService
	AuthorizeService
	RefreshService
	JwtMiddlewareService
}

func New(services Services) *echo.Echo {
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	v1 := server.Group("/api/v1")

	v1.POST("/register", register(services))
	v1.POST("/authorize", authorize(services))
	v1.POST("/refresh", refresh(services))

	users := v1.Group("/users")
	users.Use(JWTAuthMiddleware(services))
	users.POST("", userList())
	// users.GET("/:id", getUserHandler)
	//users.PUT("/:id", updateUserHandler)

	return server
}
