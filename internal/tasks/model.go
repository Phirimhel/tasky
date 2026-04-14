package tasks

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t Task) ToResponse() taskResponse {
	return taskResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func TasksToResponse(tasks []Task) TasksResponse {
	data := make([]taskResponse, len(tasks))

	for i := range tasks {
		data[i] = tasks[i].ToResponse()
	}

	return TasksResponse{
		Data: data,
	}
}
