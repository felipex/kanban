package domain

type Board struct {
	ID         uint64       `json:"id"`
	Name       string       `json:"name"`
	Columns    []Column     `json:"columns"`
	Team       []TeamMember `json:"team"`
	Categories []Category   `json:"categories"`
}
