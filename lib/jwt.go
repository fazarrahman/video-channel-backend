package lib

import (
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID int64 `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int64) (*string, *echo.HTTPError) {
	expHour, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_HOUR"))
	if err != nil {
		expHour = 6
	}
	expirationTime := time.Now().Add(time.Duration(expHour) * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError,
			"Error when generating access token : ", err.Error())
	}
	return &accessToken, nil
}
