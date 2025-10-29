package models

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending"
	StatusInProgress TaskStatus = "in_progress"
	StatusCompleted  TaskStatus = "completed"
)

type Tasks struct {
	Title  string     `json:"title"`
	Status TaskStatus `json:"status"`
}
