package main

import (
	"fmt"
	"os"

	"github.com/fazarrahman/video-channel-backend/config/postgre"
	userRepoPostgre "github.com/fazarrahman/video-channel-backend/domain/users/repository/postgre"
	"github.com/fazarrahman/video-channel-backend/handler/rest"
	"github.com/fazarrahman/video-channel-backend/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	godotenv.Load()
	db := postgre.Connect()
	userRepo := userRepoPostgre.New(db)
	svc := service.New(userRepo)
	rest.New(svc).HandlerRegister(e)
	fmt.Println("App run at port " + os.Getenv("APP_PORT"))
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
