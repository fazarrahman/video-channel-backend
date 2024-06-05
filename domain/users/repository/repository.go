package repository

import (
	"context"

	"github.com/fazarrahman/video-channel-backend/domain/users/entity"
	"github.com/labstack/echo/v4"
)

type UsersInterface interface {
	Create(ctx context.Context, user *entity.Users) (*entity.Users, *echo.HTTPError)
	GetByParam(ctx context.Context, param string, value interface{}) (*entity.Users, *echo.HTTPError)
}
