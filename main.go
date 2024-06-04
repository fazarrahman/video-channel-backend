package main

import (
	"github.com/fazarrahman/video-channel-backend/config/postgre"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	_ = postgre.Connect()
}
