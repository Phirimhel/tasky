package tasks

import (
	"time"

	"github.com/google/uuid"
)

type createTaskRequestDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type updateTaskRequestDTO struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}

type taskResponseDTO struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TasksResponse struct {
	Data []taskResponseDTO `json:"data"`
}
