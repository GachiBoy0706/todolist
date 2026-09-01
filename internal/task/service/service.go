package service

import (
	"ToDoList/internal/task/repository"
	"time"
)

type TaskModel struct {
	ID             int
	Topic          string
	Description    string
	CreationTime   time.Time
	CompletionTime *time.Time
}

type ServiceInterface interface {
	CreateTask(m TaskModel) (TaskModel, error)
	GetTasks() ([]TaskModel, error)
	CompleteTask(id int) (TaskModel, error)
	DeleteTask(id int) error
}

type Service struct {
	rep repository.ResopitoryInterface
}

func NewService(rep repository.ResopitoryInterface) *Service {
	return &Service{rep: rep}
}

func CreateTask(m TaskModel) (TaskModel, error) {

}

func GetTasks() ([]TaskModel, error) {

}

func CompleteTask(id int) (TaskModel, error) {

}

func DeleteTask(id int) error {

}
