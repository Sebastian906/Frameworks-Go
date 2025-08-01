package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		// return c.String(http.StatusOK, "Hola mundo")
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hola desde Echo!",
		})
	})

	e.GET("/saludo", func(c echo.Context) error {
		return c.String(http.StatusOK, "Saludos desde Echo!")
	})

	e.Start(":8080")
}