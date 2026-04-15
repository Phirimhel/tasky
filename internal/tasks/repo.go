package tasks

import (
	"errors"

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
}

func NewRepository() *tasksRepo {
	return &tasksRepo{
		Repo: make(map[uuid.UUID]Task),
	}

}

func (t *tasksRepo) createTask(task Task) error {
	t.Repo[task.ID] = task

	return nil
}

func (t *tasksRepo) updateTask(updatedTask Task) (Task, error) {

	id := updatedTask.ID

	_, err := t.getTaskByID(id)
	if err != nil {
		return Task{}, err
	}

	t.Repo[id] = updatedTask

	return t.Repo[id], nil

}

func (t *tasksRepo) getTaskByID(taskID uuid.UUID) (Task, error) {
	task, ok := t.Repo[taskID]

	if !ok {
		return Task{}, errors.New("task is not found")
	}

	return task, nil
}

func (t *tasksRepo) getTasks() map[uuid.UUID]Task {
	copy := t.Repo
	return copy
}

func (t *tasksRepo) deleteTask(taskID uuid.UUID) error {
	_, err := t.getTaskByID(taskID)
	if err != nil {
		return err
	}
	delete(t.Repo, taskID)
	return nil
}

func (t *tasksRepo) getLength() int {
	return len(t.Repo)
}
