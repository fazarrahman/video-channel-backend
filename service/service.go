package service

import (
	"github.com/fazarrahman/video-channel-backend/domain/users/entity"
	usersRepository "github.com/fazarrahman/video-channel-backend/domain/users/repository"
	"github.com/fazarrahman/video-channel-backend/model"
	"github.com/labstack/echo/v4"
)

type Service struct {
	UsersRepository usersRepository.UsersInterface
}

func New(userRepo usersRepository.UsersInterface) *Service {
	return &Service{UsersRepository: userRepo}
}

type ServiceInterface interface {
	Create(ctx echo.Context, users entity.Users) (*entity.Users, *echo.HTTPError)
	SignIn(ctx echo.Context, signinInput model.Signin) (*string, *echo.HTTPError)
}
