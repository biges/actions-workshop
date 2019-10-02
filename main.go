package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.BodyLimit("2M"))
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, make(map[string]interface{}))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
