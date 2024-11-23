package server

import (
	"github/felipex/kanban-server/internal/domain"
	"github/felipex/kanban-server/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateCategoryRequest struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

// @title API kanban
// @version  0.1
// @description Nova Categoria
// @host localhost:8080
// @basePath /
func newCategory(ctx *gin.Context) {
	request := CreateCategoryRequest{}
	ctx.BindJSON(&request)

	category := domain.Category{
		Name:  request.Name,
		Color: request.Color,
	}
	repository.NewCategory(category)
	ctx.JSON(http.StatusCreated, category)
}

func updateCategory(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	request := CreateCategoryRequest{}
	ctx.BindJSON(&request)

	category := domain.Category{
		ID:    id,
		Name:  request.Name,
		Color: request.Color,
	}
	repository.UpdateCategory(category)
	ctx.JSON(http.StatusOK, category)
}

func deleteCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := repository.DeleteCategory(id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, nil)
}

func getCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	category, err := repository.GetCategory(id)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, category)
}

func getCategories(c *gin.Context) {
	categories, err := repository.GetAllCategory()

	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, categories)
}
