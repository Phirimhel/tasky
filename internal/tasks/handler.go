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

func NewHandler(service TaskService) *Handler {
	return &Handler{Service: service}
}

/*
pattern - /task
method - POST
info - JSON request createTaskRequestDTO
*/
func (t *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {

	var req createTaskRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := req.requestValidationDTO()
	if err != nil {
		utils.ResponseWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	Task, err := t.Service.createTask(req.Title, req.Description)
	if err != nil {
		utils.ResponseWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := Task.ToResponse()
	err = utils.ResponseWithJSON(w, res, http.StatusCreated)
	if err != nil {
		utils.ResponseWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
pattern - /task/{id}
method - GET
info - pattern
*/
func (t *Handler) GetTaskByID(w http.ResponseWriter, r *http.Request) {

	stringTaskID := r.PathValue("id")
	taskID, err := uuid.Parse(stringTaskID)
	if err != nil {
		utils.ResponseWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	Task, err := t.Service.getTask(taskID)
	if err != nil {
		utils.ResponseWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := Task.ToResponse()
	err = utils.ResponseWithJSON(w, res, http.StatusOK)
	if err != nil {
		utils.ResponseWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
pattern - /tasks
method - GET
info - pattern + query params: [completed] for filter
*/
func (t *Handler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	//nil, true, false for the filter of tasks
	var isCompleted *bool
	if val := queries.Get("completed"); val != "" {
		b, err := strconv.ParseBool(val)
		if err != nil {
			utils.ResponseWithError(w, err.Error(), http.StatusBadRequest)
			return
		}
		isCompleted = &b
	}

	tasks := t.Service.getAllTasks(isCompleted)
	res := TasksToResponse(tasks)
	err := utils.ResponseWithJSON(w, res, http.StatusOK)
	if err != nil {
		utils.ResponseWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
pattern - /task/{id}
method - UPDATE
info - JSON request updateTaskRequestDTO
*/
func (t *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")
	taksID, err := uuid.Parse(idString)
	if err != nil {
		utils.ResponseWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req updateTaskRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	Task, err := t.Service.updateTask(
		taksID,
		req.Title,
		req.Description,
		req.Completed,
	)
	if err != nil {
		utils.ResponseWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := Task.ToResponse()
	err = utils.ResponseWithJSON(w, res, http.StatusOK)
	if err != nil {
		utils.ResponseWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
pattern - /task/{id}
method - DELETE
info - _
*/
func (t *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {

	stringTaskID := r.PathValue("id")
	taskID, err := uuid.Parse(stringTaskID)
	if err != nil {
		utils.ResponseWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = t.Service.deleteTask(taskID)
	if err != nil {
		utils.ResponseWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = utils.ResponseWithJSON(w, taskResponseDTO{}, http.StatusNoContent)
	if err != nil {
		utils.ResponseWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
