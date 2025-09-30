package service

import (
	"github.com/wolf1848/taxiportal/model"
	repositoryAuthorize "github.com/wolf1848/taxiportal/repository/authorize"
	repositoryRegister "github.com/wolf1848/taxiportal/repository/register"
	serviceAuthorize "github.com/wolf1848/taxiportal/service/authorize"
	"github.com/wolf1848/taxiportal/service/jwt"
	serviceRegister "github.com/wolf1848/taxiportal/service/register"
	"github.com/wolf1848/taxiportal/service/tools"
)

type Repositories interface {
	Register() *repositoryRegister.Repository
	Authorize() *repositoryAuthorize.Repository
}

type Services struct {
	register  *serviceRegister.Service
	authorize *serviceAuthorize.Service
	jwt       *jwt.Service
}

func NewServices(cfg *model.AppApiConfig, repositories Repositories, log tools.Logger) *Services {
	jwtService := jwt.NewService(cfg)
	return &Services{
		register:  serviceRegister.NewService(cfg, repositories.Register(), log),
		authorize: serviceAuthorize.NewService(cfg, repositories.Authorize(), log, jwtService),
		jwt:       jwtService,
	}
}

func (s *Services) RegisterService() *serviceRegister.Service {
	return s.register
}

func (s *Services) AuthorizeService() *serviceAuthorize.Service {
	return s.authorize
}

func (s *Services) JwtService() *jwt.Service {
	return s.jwt
}
