package repository

import (
	"context"

	"github.com/fazarrahman/video-channel-backend/domain/films/entity"
	"github.com/labstack/echo/v4"
)

type FilmsInterface interface {
	Create(ctx context.Context, film *entity.Films) (*entity.Films, *echo.HTTPError)
	GetByParam(ctx context.Context, param string, value interface{}) (*entity.Films, *echo.HTTPError)
	GetAll(ctx context.Context) ([]*entity.Films, *echo.HTTPError)
	Update(ctx context.Context, film entity.Films) *echo.HTTPError
	DeleteById(ctx context.Context, id int64) *echo.HTTPError
}
