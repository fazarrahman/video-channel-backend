package repository

import (
	"context"

	"github.com/fazarrahman/video-channel-backend/domain/users/entity"
	"github.com/labstack/echo/v4"
)

type UsersInterface interface {
	Create(ctx context.Context, user *entity.Users) (*entity.Users, *echo.HTTPError)
	GetById(ctx context.Context, id int64) (*entity.Users, *echo.HTTPError)
}
