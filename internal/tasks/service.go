package tasks

import (
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
	repo TasksRepository
}

func NewService(repo TasksRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (t *Service) createTask(title, description string) (Task, error) {

	// todo add some length validations
	newTask := Task{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	t.repo.createTask(newTask)

	return newTask, nil
}

func (t *Service) getTask(taskID uuid.UUID) (Task, error) {
	task, err := t.repo.getTaskByID(taskID)

	if err != nil {
		return Task{}, err
	}

	return task, nil
}

func (t *Service) getAllTasks(isCompleted *bool) []Task {
	tasks := make([]Task, 0, t.repo.getLength())

	for _, task := range t.repo.getTasks() {
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

	task, err := t.repo.getTaskByID(taskID)
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

	t.repo.updateTask(task)
	return task, nil
}

func (t *Service) deleteTask(taskID uuid.UUID) (Task, error) {
	err := t.repo.deleteTask(taskID)
	if err != nil {
		return Task{}, err
	}
	return Task{}, nil
}
