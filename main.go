package main

import (
	"os"
	"tasks-app/internal/httpServer"
	"tasks-app/internal/tasks"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	repo := tasks.NewRepository()
	service := tasks.NewService(repo)
	handler := tasks.NewHandler(service)
	routes := tasks.RegisterRoutes(handler)

	server := httpServer.NewHTTPServer(routes, port)
	server.StartServer()
}
