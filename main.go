package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		items := []gin.H{
			{"id": 1, "titulo": "Titulo 1"},
			{"id": 2, "titulo": "Titulo 2"},
			{"id": 3, "titulo": "Titulo 3"},
		} // Sua lista de itens aqui
		c.JSON(http.StatusOK, gin.H{
			"items": items,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}