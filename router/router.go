package router

import (
	"belajar-echo/handler"
	"belajar-echo/repository"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo, db *sql.DB) {

	ProductRepo := repository.ProductRepo{Sql: db}
	ProductHandler := handler.NewProductHandler(&ProductRepo)

	v1 := e.Group("/v1")
	v1.GET("/ping", ping)
	v1.POST("/login", nil).Name = "Signin"

	product := v1.Group("/product")
	product.Use()
	product.GET("/", ProductHandler.GetAll)
}

func ping(c echo.Context) error {
	return c.String(200, "ping")
}
