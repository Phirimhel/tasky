package tasks

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type TasksRepository interface {
	createTask(task Task) error
	updateTask(updatedTask Task) (Task, error)
	getTaskByID(taskID uuid.UUID) (Task, error)
	getTasks() map[uuid.UUID]Task
	deleteTask(taskID uuid.UUID) error
	getLength() int
}

type tasksRepo struct {
	Repo map[uuid.UUID]Task
	mu   *sync.RWMutex
}

func NewRepository(mutex *sync.RWMutex) *tasksRepo {
	return &tasksRepo{
		Repo: make(map[uuid.UUID]Task),
		mu:   mutex,
	}

}

func (t *tasksRepo) createTask(task Task) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.Repo[task.ID] = task
	return nil
}

func (t *tasksRepo) updateTask(updatedTask Task) (Task, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	taskID := updatedTask.ID
	_, ok := t.Repo[taskID]
	if !ok {
		return Task{}, errors.New("task is not found")
	}

	t.Repo[taskID] = updatedTask
	return updatedTask, nil

}

func (t *tasksRepo) getTaskByID(taskID uuid.UUID) (Task, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	task, ok := t.Repo[taskID]
	if !ok {
		return Task{}, errors.New("task is not found")
	}

	return task, nil
}

func (t *tasksRepo) getTasks() map[uuid.UUID]Task {
	t.mu.RLock()
	defer t.mu.RUnlock()

	copyMap := make(map[uuid.UUID]Task, len(t.Repo))
	for _, v := range t.Repo {
		copyMap[v.ID] = v
	}
	return copyMap

}

func (t *tasksRepo) deleteTask(taskID uuid.UUID) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	_, ok := t.Repo[taskID]
	if !ok {
		return errors.New("task is not found")
	}

	delete(t.Repo, taskID)
	return nil
}

func (t *tasksRepo) getLength() int {
	t.mu.RLock()
	defer t.mu.RUnlock()

	length := len(t.Repo)
	return length
}
