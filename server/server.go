package server

import (
	"net/http"

	"github.com/Sreethecool/filestore/server/controllers"
	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	return nil
}

func RunServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome to File Store Service")
	})
	e.POST("/upload", controllers.Upload)
	e.GET("/list", List)
	e.Logger.Fatal(e.Start(":8080"))
}
