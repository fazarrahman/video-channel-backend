package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fazarrahman/video-channel-backend/config/postgre"
	filmRepoPostgre "github.com/fazarrahman/video-channel-backend/domain/films/entity/repository/postgre"
	userRepoPostgre "github.com/fazarrahman/video-channel-backend/domain/users/repository/postgre"
	"github.com/fazarrahman/video-channel-backend/handler/rest"
	"github.com/fazarrahman/video-channel-backend/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodOptions, http.MethodPost, http.MethodDelete, http.MethodPut},
	}))

	db := postgre.Connect()
	userRepo := userRepoPostgre.New(db)
	filmRepo := filmRepoPostgre.New(db)
	svc := service.New(userRepo, filmRepo)
	rest.New(svc).HandlerRegister(e)
	fmt.Println("App run at port " + os.Getenv("APP_PORT"))
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
