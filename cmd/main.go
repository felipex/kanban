// Pacote peincipal da aplicacao
package main

import (
	"github/felipex/kanban-server/internal/app/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routers.Rotas(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world!"})
	})

	router.Run(":8080")
}
