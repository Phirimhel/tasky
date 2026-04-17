package tasks

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type createTaskRequestDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (c *createTaskRequestDTO) requestValidationDTO() error {
	if c.Title == "" {
		return errors.New("title is empty")
	}
	if c.Description == "" {
		return errors.New("descritpion is empty")
	}

	return nil
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

type tasksResponseDTO struct {
	Data []taskResponseDTO `json:"data"`
}
