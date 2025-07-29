# Servidor Web Básico en Go

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hola, mundo")
	})
	r.Run(":8080")
}
```