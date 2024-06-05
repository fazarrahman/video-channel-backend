package postgre

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/fazarrahman/video-channel-backend/domain/users/entity"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Postgre struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *Postgre {
	return &Postgre{Db: db}
}

func (p *Postgre) Create(ctx context.Context, user *entity.Users) (*entity.Users, *echo.HTTPError) {
	lastInsertId := 0
	row := p.Db.QueryRowContext(ctx,
		`INSERT INTO users (username, password_hash, email) 
		VALUES ($1,$2,$3) RETURNING id`, user.Username, user.PasswordHash, user.Email).
		Scan(&lastInsertId)

	if row != nil && row.Error() != "" {
		return nil, echo.NewHTTPError(http.StatusInternalServerError,
			"Error when inserting user data : "+row.Error())
	}
	user.Id = int64(lastInsertId)
	return user, nil
}

func (p *Postgre) GetByParam(ctx context.Context, param string, value interface{}) (*entity.Users, *echo.HTTPError) {
	var user entity.Users
	err := p.Db.GetContext(ctx, &user, "SELECT id, username, password_hash, email FROM users where "+param+" = $1", value)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError,
			"Error when getting user data : "+err.Error())
	}
	return &user, nil
}
