package domain

import (
	"time"

	"github.com/google/uuid"
)

type TodoItem struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	FileID      string    `json:"fileId,omitempty"`
}

func NewTodoItem(description string, dueDate time.Time, fileID string) *TodoItem {
	return &TodoItem{
		ID:          uuid.New(),
		Description: description,
		DueDate:     dueDate,
		FileID:      fileID,
	}
}
