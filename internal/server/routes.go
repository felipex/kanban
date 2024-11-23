package server

import (
	"github.com/gin-gonic/gin"
)

func Rotas(r *gin.Engine) {
	r.POST("/category", newCategory)
	r.PUT("/category/:id", updateCategory)
	r.DELETE("/category/:id", deleteCategory)
	r.GET("/category/:id", getCategory)
	r.GET("/category", getCategories)
}
