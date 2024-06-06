package postgre

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/fazarrahman/video-channel-backend/domain/films/entity"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Postgre struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *Postgre {
	return &Postgre{Db: db}
}

func (p *Postgre) Create(ctx context.Context, film *entity.Films) (*entity.Films, *echo.HTTPError) {
	lastInsertId := 0
	row := p.Db.QueryRowContext(ctx,
		`INSERT INTO Films (title, description, image_thumbnail) 
		VALUES ($1,$2,$3) RETURNING id`, film.Title, film.Description, film.ImageThumbnail).
		Scan(&lastInsertId)

	if row != nil && row.Error() != "" {
		return nil, echo.NewHTTPError(http.StatusInternalServerError,
			"Error when inserting film data : "+row.Error())
	}
	film.Id = int64(lastInsertId)
	return film, nil
}

func (p *Postgre) GetByParam(ctx context.Context, param string, value interface{}) (*entity.Films, *echo.HTTPError) {
	var film entity.Films
	err := p.Db.GetContext(ctx, &film, "SELECT id, title, description, image_thumbnail FROM Films where "+param+" = $1", value)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError,
			"Error when getting film data : "+err.Error())
	}
	return &film, nil
}

func (p *Postgre) GetAll(ctx context.Context) ([]*entity.Films, *echo.HTTPError) {
	var films []*entity.Films
	err := p.Db.SelectContext(ctx, &films, "SELECT id, title, description, image_thumbnail FROM Films")
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError,
			"Error when getting all film data : "+err.Error())
	}
	return films, nil
}

func (p *Postgre) Update(ctx context.Context, film entity.Films) *echo.HTTPError {
	_, err := p.Db.NamedExecContext(ctx, `UPDATE Films 
	SET title=:title, description=:description, image_thumbnail=:image_thumbnail
	WHERE id=:id`, film)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error when updating the film data"+err.Error())
	}
	return nil
}

func (p *Postgre) DeleteById(ctx context.Context, id int64) *echo.HTTPError {
	_, err := p.Db.ExecContext(ctx, "DELETE FROM Films WHERE id=$1", id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error when deleting film data"+err.Error())
	}
	return nil
}
