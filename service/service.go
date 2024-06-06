package service

import (
	filmEntity "github.com/fazarrahman/video-channel-backend/domain/films/entity"
	filmsRepository "github.com/fazarrahman/video-channel-backend/domain/films/entity/repository"
	"github.com/fazarrahman/video-channel-backend/domain/users/entity"
	usersRepository "github.com/fazarrahman/video-channel-backend/domain/users/repository"
	"github.com/fazarrahman/video-channel-backend/model"
	"github.com/labstack/echo/v4"
)

type Service struct {
	UsersRepository usersRepository.UsersInterface
	FilmsRepository filmsRepository.FilmsInterface
}

func New(userRepo usersRepository.UsersInterface, filmRepo filmsRepository.FilmsInterface) *Service {
	return &Service{UsersRepository: userRepo, FilmsRepository: filmRepo}
}

type ServiceInterface interface {
	CreateUser(ctx echo.Context, users entity.Users) (*entity.Users, *echo.HTTPError)
	SignIn(ctx echo.Context, signinInput model.Signin) (*string, *echo.HTTPError)
	CreateFilm(ctx echo.Context, films *filmEntity.Films) (*filmEntity.Films, *echo.HTTPError)
	GetAllFilm(ctx echo.Context) ([]*filmEntity.Films, *echo.HTTPError)
	GetFilmById(ctx echo.Context, id int64) (*filmEntity.Films, *echo.HTTPError)
	UpdateFilm(ctx echo.Context, films filmEntity.Films) *echo.HTTPError
	DeleteFilmById(ctx echo.Context, id int64) *echo.HTTPError
}
