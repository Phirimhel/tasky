package main

import (
	"os"
	"sync"
	"tasks-app/internal/httpServer"
	"tasks-app/internal/tasks"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	repo := tasks.NewRepository(&sync.RWMutex{})
	service := tasks.NewService(repo)
	handler := tasks.NewHandler(service)
	routes := tasks.RegisterRoutes(handler)

	server := httpServer.NewHTTPServer(routes, port)
	server.StartServer()
}
