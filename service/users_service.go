package service

import (
	"net/http"

	"github.com/fazarrahman/video-channel-backend/domain/users/entity"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Create(ctx echo.Context, users entity.Users) (*entity.Users, *echo.HTTPError) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(users.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Error when encrypting password "+err.Error())
	}
	users.PasswordHash = pwdHash
	res, errl := s.UsersRepository.Create(ctx.Request().Context(), &users)
	if errl != nil {
		return nil, errl
	}
	res.Password = ""
	res.PasswordHash = nil
	return res, nil
}
