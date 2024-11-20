package domain

import (
	"time"
)

type Card struct {
	ID          uint64     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Assignee    TeamMember `json:"assignee"`
	DueDate     time.Time  `json:"dueDate"`
	Labels      []string   `json:"labels"`
	Starts      int64      `json:"starts"`
	Category    Category   `json:"category"`
}
