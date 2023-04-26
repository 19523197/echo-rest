package handler

import (
	"belajar-echo/auth"
	"belajar-echo/model"
	"belajar-echo/repository"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Username string
	Password string
}

type LoginHandler struct {
	Repo *repository.UserRepository
}

func (h *LoginHandler) SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.QueryParam("Username")
		password := c.QueryParam("Password")

		if res, err := h.Repo.LoginUser(username, password); res == false {
			return echo.ErrUnauthorized
		} else if err != nil {
			return echo.ErrInternalServerError
		}

		user := model.User{username, password}
		err := auth.GenerateTokenAndCookies(&user, c)
		if err != nil {
			return echo.ErrUnauthorized
		}

		return c.Redirect(http.StatusMovedPermanently, "/product")
	}
}
