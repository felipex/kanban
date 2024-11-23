package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.Default()

	Rotas(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world!"})
	})

	router.Run(":8080")
}
