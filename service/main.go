package service

import (
	"github.com/wolf1848/taxiportal/model"
)

type Logger interface {
	Debug(msg string, fields ...map[string]any)
	Info(msg string, fields ...map[string]any)
	Warn(msg string, fields ...map[string]any)
	Error(msg string, fields ...map[string]any)
}

type Repositories interface {
	UserInsert(input *model.User) (*int, error)
	UserFindByEmail(email string) (*model.User, error)
	UserGetById(id int) (*model.User, error)
}

type Services struct {
	*model.AppApiConfig
	Logger
	Repositories
}

func NewServices(cfg *model.AppApiConfig, repositories Repositories, log Logger) *Services {
	return &Services{cfg, log, repositories}
}
