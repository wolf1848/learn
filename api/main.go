package api

import (
	"github.com/labstack/echo/v4"
	authorizePost "github.com/wolf1848/taxiportal/api/authorize/post"
	refreshPost "github.com/wolf1848/taxiportal/api/refresh/post"
	registerPost "github.com/wolf1848/taxiportal/api/register/post"
	usersPost "github.com/wolf1848/taxiportal/api/users/post"
	serviceAuthorize "github.com/wolf1848/taxiportal/service/authorize"
	"github.com/wolf1848/taxiportal/service/jwt"
	serviceRegister "github.com/wolf1848/taxiportal/service/register"
)

type Services interface {
	RegisterService() *serviceRegister.Service
	AuthorizeService() *serviceAuthorize.Service
	JwtService() *jwt.Service
}

func New(services Services) *echo.Echo {
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	v1 := server.Group("/api/v1")

	v1.POST("/register", registerPost.Handler(services.RegisterService()))
	v1.POST("/authorize", authorizePost.Handler(services.AuthorizeService()))
	v1.POST("/refresh", refreshPost.Handler(services.AuthorizeService()))

	users := v1.Group("/users")
	users.Use(JWTAuthMiddleware(services.JwtService()))
	users.POST("", usersPost.Handler())
	// users.GET("/:id", getUserHandler)
	//users.PUT("/:id", updateUserHandler)

	return server
}
