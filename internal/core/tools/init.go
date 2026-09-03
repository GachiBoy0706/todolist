package tools

import (
	"ToDoList/internal/task/repository"
	"ToDoList/internal/task/service"
	"ToDoList/internal/task/transport"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Init(router *mux.Router, pool *pgxpool.Pool) {
	repo := repository.NewRepo(pool)
	svc := service.NewService(repo)
	handler := transport.NewHandler(svc)

	router.Path("/tasks/").Methods("POST").HandlerFunc(handler.CreateTask)
}
