package server

import (
	"net/http"

	"github.com/Sreethecool/filestore/server/controllers"
	"github.com/Sreethecool/filestore/server/utils"
	"github.com/labstack/echo/v4"
)

func initServer() {
	utils.CreateDirIfNotExists("upload/")
}

func RunServer() {
	initServer()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Welcome to File Store Service")
	})
	e.POST("/upload", controllers.Upload)
	e.GET("/list", controllers.List)
	e.POST("/delete", controllers.Delete)
	e.POST("/run", controllers.Execute)
	e.Logger.Fatal(e.Start(":8080"))
}
