package tasks

import "net/http"

func NewMuxRouter(handler TaskHandler) http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("POST /task", handler.CreateTask)
	mux.HandleFunc("GET /task/{id}", handler.GetTaskByID)
	mux.HandleFunc("GET /tasks", handler.GetAllTasks)
	mux.HandleFunc("PATCH /task/{id}", handler.UpdateTask)
	mux.HandleFunc("DELETE /task/{id}", handler.DeleteTask)

	return mux
}
