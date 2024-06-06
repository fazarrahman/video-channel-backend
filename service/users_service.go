package service

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/fazarrahman/video-channel-backend/domain/users/entity"
	jwtLib "github.com/fazarrahman/video-channel-backend/lib"
	"github.com/fazarrahman/video-channel-backend/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateUser(ctx echo.Context, users entity.Users) (*entity.Users, *echo.HTTPError) {
	if strings.TrimSpace(users.Username) == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Username is required")
	} else if strings.TrimSpace(users.Password) == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Password is required")
	} else if strings.TrimSpace(users.Email) == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Email is required")
	} else if !isValidEmail(users.Email) {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Incorrect email format")
	}

	userByUsername, err := s.UsersRepository.GetByParam(ctx.Request().Context(), "username", users.Username)
	if err != nil {
		return nil, err
	}
	if userByUsername != nil {
		return nil, echo.NewHTTPError(http.StatusConflict, "Username is already used")
	}

	userByEmail, err := s.UsersRepository.GetByParam(ctx.Request().Context(), "email", users.Email)
	if err != nil {
		return nil, err
	}
	if userByEmail != nil {
		return nil, echo.NewHTTPError(http.StatusConflict, "Email is already used")
	}

	pwdHash, errl := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	if errl != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Error when encrypting password "+errl.Error())
	}
	users.PasswordHash = pwdHash
	res, err := s.UsersRepository.Create(ctx.Request().Context(), &users)
	if err != nil {
		return nil, err
	}
	res.Password = ""
	res.PasswordHash = nil
	return res, nil
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)

	return re.MatchString(email)
}

func (s *Service) SignIn(ctx echo.Context, signinInput model.Signin) (*string, *echo.HTTPError) {
	if strings.TrimSpace(signinInput.Email) == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Email is required")
	}
	if strings.TrimSpace(signinInput.Password) == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Password is required")
	}

	user, err := s.CheckEmailAndPassword(ctx, signinInput)
	if user == nil && err != nil {
		return nil, err
	}

	token, erro := jwtLib.GenerateToken(user.Id)

	if erro != nil {
		return nil, erro
	}

	return token, nil
}

func (s *Service) CheckEmailAndPassword(ctx echo.Context, r model.Signin) (*entity.Users, *echo.HTTPError) {
	userEntity, err := s.UsersRepository.GetByParam(ctx.Request().Context(), "email", r.Email)
	if userEntity == nil && err == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid email address")
	} else if err != nil {
		return nil, err
	}

	erro := bcrypt.CompareHashAndPassword(userEntity.PasswordHash, []byte(r.Password))
	if erro != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid password : "+erro.Error())
	}
	return userEntity, nil
}
