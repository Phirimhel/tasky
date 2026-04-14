package main

import (
	"net/http"
	"os"
	"tasks-app/internal/tasks"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := ":" + os.Getenv("PORT")

	repo := tasks.NewTasksRepository()
	service := tasks.NewTasksService(repo)
	handler := tasks.NewTaskHandler(service)
	router := tasks.NewMuxRouter(handler)

	http.ListenAndServe(port, router)
}
