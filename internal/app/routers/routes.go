package routers

import (
	"github/felipex/kanban-server/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func categoria(c *gin.Context) {
	cat := domain.Category{1, "moveis cariutaba", "red"}

	c.JSON(http.StatusOK, cat)
}

func Rotas(r *gin.Engine) {
	r.GET("/categoria", categoria)
}
