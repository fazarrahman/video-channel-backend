package rest

import (
	"net/http"

	"github.com/fazarrahman/video-channel-backend/domain/users/entity"
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
}

func (r *Rest) CreateUser(c echo.Context) error {
	user := entity.Users{}
	errl := c.Bind(&user)
	if errl != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Error binding : " + errl.Error()})
	}
	res, err := r.service.Create(c, user)
	if err != nil {
		return c.JSON(err.Code, echo.Map{"message": err.Message})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Success", "Users": res})
}
