package tasks

import "github.com/google/uuid"

func NewTasksRepository() map[uuid.UUID]Task {
	return make(map[uuid.UUID]Task)
}
