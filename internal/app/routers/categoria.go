package routers

type CreateCategoryRequest struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}
