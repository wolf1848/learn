package service

import (
	"errors"

	"github.com/wolf1848/taxiportal/model"
	"github.com/wolf1848/taxiportal/service/entity"
)

func (service *Services) Register(input *entity.RegisterInput) (*entity.RegisterOutput, error) {
	user := &model.User{
		Name:  input.Name,
		Email: input.Email,
	}

	err := user.SetPwd(input.Pwd)
	if err != nil {
		return nil, entity.NewErrValidRegister("pwd", model.ErrHashPwd)
	}

	id, err := service.UserInsert(user)
	if err != nil {
		if errors.Is(err, model.ErrUniqueEmail) {
			return nil, entity.NewErrValidRegister("email", err)
		}

		service.Error(err.Error())

		return nil, entity.ErrService
	}

	return &entity.RegisterOutput{
		ID:    *id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (service *Services) Authorize(input *entity.AuthorizeInput) (*entity.AuthorizeOutput, error) {

	user, err := service.UserFindByEmail(input.Email)
	if err != nil {
		if errors.Is(err, entity.ErrNoRows) {
			return nil, entity.ErrAuthorized
		}
		service.Error(err.Error())
		return nil, entity.ErrService
	}

	if ok := user.CheckPasswordHash(input.Pwd); !ok {
		return nil, entity.ErrAuthorized
	}

	token, err := service.JwtGenerateAccessToken(user.ID)
	if err != nil {
		service.Error(err.Error())
		return nil, entity.ErrService
	}

	refreshToken, err := service.JwtGenerateRefreshToken(user.ID)
	if err != nil {
		service.Error(err.Error())
		return nil, entity.ErrService
	}

	return &entity.AuthorizeOutput{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (service *Services) RefreshToken(value string) (*entity.AuthorizeOutput, error) {

	claim, err := service.JwtValidateRefreshToken(value)
	if err != nil {
		return nil, err
	}

	user, err := service.UserGetById(claim.UserID)
	if err != nil {
		return nil, entity.ErrService
	}

	token, err := service.JwtGenerateAccessToken(user.ID)
	if err != nil {
		service.Error(err.Error())
		return nil, entity.ErrService
	}

	refreshToken, err := service.JwtGenerateRefreshToken(user.ID)
	if err != nil {
		service.Error(err.Error())
		return nil, entity.ErrService
	}

	return &entity.AuthorizeOutput{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
