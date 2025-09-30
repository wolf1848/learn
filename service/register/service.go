package register

import (
	"errors"

	"github.com/wolf1848/taxiportal/model"
	serviceErrors "github.com/wolf1848/taxiportal/service/errors"
	"github.com/wolf1848/taxiportal/service/register/entity"
	"github.com/wolf1848/taxiportal/service/tools"
)

type Repository interface {
	InsertUser(*model.User) error
}

type Service struct {
	config     *model.AppApiConfig
	logger     tools.Logger
	repository Repository
}

func NewService(cfg *model.AppApiConfig, repository Repository, log tools.Logger) *Service {
	return &Service{
		config:     cfg,
		logger:     log,
		repository: repository,
	}
}

func (s *Service) Register(input *entity.Input) (*entity.Output, error) {
	user := &model.User{
		Name:  input.Name,
		Email: input.Email,
	}

	var err error

	err = user.SetPwd(input.Pwd)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, entity.ErrHashPwd
	}

	err = s.repository.InsertUser(user)
	if err != nil {
		if errors.Is(err, entity.ErrUniqueEmail) {
			return nil, err
		}

		s.logger.Error(err.Error())

		return nil, serviceErrors.ErrService
	}

	return &entity.Output{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
