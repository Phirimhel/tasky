package tasks

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TaskService interface {
	createTask(title, description string) (Task, error)
	getTask(taskID uuid.UUID) (Task, error)
	getAllTasks(completed *bool) []Task
	updateTask(taskID uuid.UUID, title, description *string, completed *bool) (Task, error)
	deleteTask(taskID uuid.UUID) (Task, error)
}

type Service struct {
	repo map[uuid.UUID]Task
}

func NewTasksService(repo map[uuid.UUID]Task) *Service {
	return &Service{
		repo: repo,
	}
}

func (t *Service) createTask(title, description string) (Task, error) {

	if title == "" || description == "" {
		return Task{}, errors.New("title or description have not be empty")
	}

	taskID := uuid.New()

	newTask := Task{
		ID:          taskID,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.repo[taskID] = newTask

	return newTask, nil
}

func (t *Service) getTask(taskID uuid.UUID) (Task, error) {
	task, ok := t.repo[taskID]

	if !ok {
		return Task{}, errors.New("resourse is not found")
	}

	return task, nil
}

func (t *Service) getAllTasks(isCompleted *bool) []Task {
	tasks := make([]Task, 0, len(t.repo))

	for _, task := range t.repo {
		// todo others filtres by queries

		if isCompleted != nil && task.Completed != *isCompleted {
			continue
		}

		tasks = append(tasks, task)
	}
	return tasks
}

func (t *Service) updateTask(
	taskID uuid.UUID,
	title, description *string,
	completed *bool,
) (Task, error) {

	task, err := t.getTask(taskID)
	if err != nil {
		return Task{}, err
	}

	if title != nil {
		task.Title = *title
	}

	if description != nil {
		task.Description = *description
	}

	if completed != nil {
		task.Completed = *completed
	}

	task.UpdatedAt = time.Now()

	t.repo[taskID] = task
	return task, nil
}

func (t *Service) deleteTask(taskID uuid.UUID) (Task, error) {
	_, err := t.getTask(taskID)
	if err != nil {
		return Task{}, err
	}

	delete(t.repo, taskID)

	return Task{}, nil
}
