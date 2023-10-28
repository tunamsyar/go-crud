package main

import (
	"crud/router"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", RootPoint)

	router.SetupRouter(e)

	e.Start(":8080")
}

func RootPoint(c echo.Context) error {
	return c.JSON(http.StatusOK, "test")
}
