package tasks

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tasks-app/internal/utils"

	"github.com/google/uuid"
)

type TaskHandler interface {
	// C
	CreateTask(w http.ResponseWriter, r *http.Request)

	// R
	GetTaskByID(w http.ResponseWriter, r *http.Request)
	GetAllTasks(w http.ResponseWriter, r *http.Request)

	// U
	UpdateTask(w http.ResponseWriter, r *http.Request)

	// D
	DeleteTask(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	Service TaskService
}

func NewTaskHandler(service TaskService) *Handler {
	return &Handler{Service: service}
}

func (t *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {

	var req createTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	Task, err := t.Service.createTask(req.Title, req.Description)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := Task.ToResponse()
	utils.ResponseWithJSON(w, http.StatusCreated, res)
}

func (t *Handler) GetTaskByID(w http.ResponseWriter, r *http.Request) {

	stringTaskID := r.PathValue("id")
	taskID, err := uuid.Parse(stringTaskID)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	Task, err := t.Service.getTask(taskID)
	if err != nil {
		utils.ResponseWithError(w, http.StatusNotFound, err.Error())
		return
	}

	res := Task.ToResponse()
	utils.ResponseWithJSON(w, http.StatusOK, res)
}

func (t *Handler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	//nil, true, false for the filter of tasks
	var isCompleted *bool
	if val := queries.Get("completed"); val != "" {
		b, err := strconv.ParseBool(val)
		if err != nil {
			utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		isCompleted = &b
	}

	tasks := t.Service.getAllTasks(isCompleted)
	res := TasksToResponse(tasks)
	utils.ResponseWithJSON(w, http.StatusOK, res)
}

func (t *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")
	taksID, err := uuid.Parse(idString)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var req updateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	Task, err := t.Service.updateTask(
		taksID,
		req.Title,
		req.Description,
		req.Completed,
	)
	if err != nil {
		utils.ResponseWithError(w, http.StatusNotFound, err.Error())
		return
	}

	res := Task.ToResponse()
	utils.ResponseWithJSON(w, http.StatusOK, res)
}

func (t *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {

	stringTaskID := r.PathValue("id")
	taskID, err := uuid.Parse(stringTaskID)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = t.Service.deleteTask(taskID)
	if err != nil {
		utils.ResponseWithError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.ResponseWithJSON(w, http.StatusNoContent, taskResponse{})
}
