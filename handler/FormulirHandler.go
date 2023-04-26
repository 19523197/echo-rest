package handler

import (
	"belajar-echo/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

// type ProductHandlerCont interface {
// 	GetAll(echo.Context) error
// }

type ProductHandler struct {
	Repo *repository.ProductRepo
}

func NewProductHandler(repo *repository.ProductRepo) *ProductHandler {
	return &ProductHandler{Repo: repo}
}

func (p *ProductHandler) GetAll(c echo.Context) error {
	products, err := p.Repo.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, products)
}
