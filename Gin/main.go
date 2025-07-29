package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hola, mundo")
	})

	r.GET("/saludo/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.String(http.StatusOK, "Hola, %s!", nombre)
	})

	r.Run(":8080")
}
