package authorize

import (
	"errors"

	"github.com/wolf1848/taxiportal/model"
	"github.com/wolf1848/taxiportal/service/authorize/entity"
	serviceErrors "github.com/wolf1848/taxiportal/service/errors"
	jwtEntity "github.com/wolf1848/taxiportal/service/jwt/entity"
	"github.com/wolf1848/taxiportal/service/tools"
)

type Repository interface {
	UserFindByEmail(string) (*model.User, error)
	UserGetById(int) (*model.User, error)
}

type JwtService interface {
	GetAccessToken(int) (string, error)
	ValidateAccessToken(string) (*jwtEntity.AccessClaim, error)
	GetRefreshToken(int) (string, error)
	ValidateRefreshToken(string) (*jwtEntity.RefreshClaim, error)
}

type Service struct {
	logger     tools.Logger
	repository Repository
	jwt        JwtService
}

func NewService(repository Repository, log tools.Logger, jwt JwtService) *Service {
	return &Service{
		logger:     log,
		repository: repository,
		jwt:        jwt,
	}
}

func (s *Service) Authorize(input *entity.Input) (*entity.Output, error) {

	user, err := s.repository.UserFindByEmail(input.Email)
	if err != nil {
		if errors.Is(err, serviceErrors.ErrRepositoryNoRows) {
			return nil, serviceErrors.ErrAuthorized
		}
		s.logger.Error(err.Error())
		return nil, serviceErrors.ErrService
	}

	if ok := user.CheckPasswordHash(input.Pwd); !ok {
		return nil, serviceErrors.ErrAuthorized
	}

	token, err := s.jwt.GetAccessToken(user.ID)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, serviceErrors.ErrService
	}

	refreshToken, err := s.jwt.GetRefreshToken(user.ID)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, serviceErrors.ErrService
	}

	return &entity.Output{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) RefreshAuthorize(value string) (*entity.Output, error) {

	claim, err := s.jwt.ValidateRefreshToken(value)
	if err != nil {
		return nil, err
	}

	user, err := s.repository.UserGetById(claim.UserID)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, serviceErrors.ErrService
	}

	token, err := s.jwt.GetAccessToken(user.ID)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, serviceErrors.ErrService
	}

	refreshToken, err := s.jwt.GetRefreshToken(user.ID)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, serviceErrors.ErrService
	}

	return &entity.Output{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
