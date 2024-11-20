package domain

type Category struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"` // e.g., "red", "green", "blue"
}
