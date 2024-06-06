package rest

import (
	"net/http"
	"strconv"

	filmEntity "github.com/fazarrahman/video-channel-backend/domain/films/entity"
	"github.com/fazarrahman/video-channel-backend/domain/users/entity"
	"github.com/fazarrahman/video-channel-backend/model"
	"github.com/fazarrahman/video-channel-backend/service"
	"github.com/labstack/echo/v4"
)

type Rest struct {
	service service.ServiceInterface
}

func New(service service.ServiceInterface) *Rest {
	return &Rest{service: service}
}

func (r *Rest) HandlerRegister(e *echo.Echo) {
	user := e.Group("/api/users")
	user.POST("/register", r.CreateUser)
	user.POST("/signin", r.SignIn)
	film := e.Group("/api/films")
	film.POST("", r.CreateFilm)
	film.GET("", r.GetAllFilms)
	film.GET("/:id", r.GetFilmById)
	film.PUT("/:id", r.Update)
	film.DELETE("/:id", r.DeleteById)
}

func (r *Rest) CreateUser(c echo.Context) error {
	user := entity.Users{}
	errl := c.Bind(&user)
	if errl != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Error binding : " + errl.Error()})
	}
	res, err := r.service.CreateUser(c, user)
	if err != nil {
		return c.JSON(err.Code, echo.Map{"message": err.Message})
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "Success", "Users": res})
}

func (r *Rest) SignIn(c echo.Context) error {
	signIn := model.Signin{}
	errl := c.Bind(&signIn)
	if errl != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Error binding : " + errl.Error()})
	}
	accessToken, err := r.service.SignIn(c, signIn)
	if err != nil {
		return c.JSON(err.Code, echo.Map{"message": err.Message})
	}
	return c.JSON(http.StatusOK, echo.Map{"accessToken": accessToken})
}

func (r *Rest) CreateFilm(c echo.Context) error {
	film := filmEntity.Films{}
	errl := c.Bind(&film)
	if errl != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Error binding : " + errl.Error()})
	}
	res, err := r.service.CreateFilm(c, &film)
	if err != nil {
		return c.JSON(err.Code, echo.Map{"message": err.Message})
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "Success", "Films": res})
}

func (r *Rest) GetAllFilms(c echo.Context) error {
	res, err := r.service.GetAllFilm(c)
	if err != nil {
		return c.JSON(err.Code, echo.Map{"message": err.Message})
	}
	if res == nil {
		return c.JSON(http.StatusOK, echo.Map{"message": "Success", "Films": []filmEntity.Films{}})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "Films": res})
}

func (r *Rest) GetFilmById(c echo.Context) error {
	idParam := c.Param("id")
	id, errl := strconv.Atoi(idParam)
	if errl != nil {
		return c.JSON(http.StatusBadRequest, "Invalid id")
	}
	res, err := r.service.GetFilmById(c, int64(id))
	if err != nil {
		return c.JSON(err.Code, echo.Map{"message": err.Message})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "Films": res})
}

func (r *Rest) Update(c echo.Context) error {
	film := filmEntity.Films{}
	idParam := c.Param("id")
	id, errl := strconv.Atoi(idParam)
	if errl != nil {
		return c.JSON(http.StatusBadRequest, "Invalid id")
	}
	errl = c.Bind(&film)
	if errl != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Error binding : " + errl.Error()})
	}
	film.Id = int64(id)
	err := r.service.UpdateFilm(c, film)
	if err != nil {
		return c.JSON(err.Code, echo.Map{"message": err.Message})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "Films": film})
}

func (r *Rest) DeleteById(c echo.Context) error {
	idParam := c.Param("id")
	id, errl := strconv.Atoi(idParam)
	if errl != nil {
		return c.JSON(http.StatusBadRequest, "Invalid id")
	}
	err := r.service.DeleteFilmById(c, int64(id))
	if err != nil {
		return c.JSON(err.Code, echo.Map{"message": err.Message})
	}
	return c.JSON(http.StatusNoContent, echo.Map{"message": "Success"})
}
