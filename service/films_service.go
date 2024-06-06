package service

import (
	"net/http"
	"strings"

	"github.com/fazarrahman/video-channel-backend/domain/films/entity"
	"github.com/labstack/echo/v4"
)

func validation(films *entity.Films) *echo.HTTPError {
	if strings.TrimSpace(films.Title) == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Title is required")
	} else if strings.TrimSpace(films.Description) == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Description is required")
	} else if strings.TrimSpace(films.ImageThumbnail) == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Image thumbnail is required")
	}

	return nil
}

func (s *Service) CreateFilm(ctx echo.Context, films *entity.Films) (*entity.Films, *echo.HTTPError) {
	if err := validation(films); err != nil {
		return nil, err
	}
	res, err := s.FilmsRepository.Create(ctx.Request().Context(), films)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) GetAllFilm(ctx echo.Context) ([]*entity.Films, *echo.HTTPError) {
	res, err := s.FilmsRepository.GetAll(ctx.Request().Context())
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) GetFilmById(ctx echo.Context, id int64) (*entity.Films, *echo.HTTPError) {
	res, err := s.FilmsRepository.GetByParam(ctx.Request().Context(), "id", id)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Data not found")
	}
	return res, nil
}

func (s *Service) UpdateFilm(ctx echo.Context, films entity.Films) *echo.HTTPError {
	if films.Id < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}
	if err := validation(&films); err != nil {
		return err
	}

	existing, err := s.FilmsRepository.GetByParam(ctx.Request().Context(), "id", films.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error when getting film data : "+err.Error())
	}
	if existing == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Data not found")
	}

	err = s.FilmsRepository.Update(ctx.Request().Context(), films)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteFilmById(ctx echo.Context, id int64) *echo.HTTPError {
	existing, err := s.FilmsRepository.GetByParam(ctx.Request().Context(), "id", id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error when getting film data : "+err.Error())
	}
	if existing == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Data not found")
	}
	err = s.FilmsRepository.DeleteById(ctx.Request().Context(), id)
	if err != nil {
		return err
	}
	return nil
}
