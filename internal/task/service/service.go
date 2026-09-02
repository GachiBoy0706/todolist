package service

import (
	"ToDoList/internal/task/repository"
	"context"
	"time"
)

type TaskModel struct {
	ID             int
	Title          string
	Description    string
	CreationTime   time.Time
	CompletionTime *time.Time
}

type ServiceInterface interface {
	CreateTask(ctx context.Context, m TaskModel) (TaskModel, error)
	GetTasks(ctx context.Context) ([]TaskModel, error)
	CompleteTask(ctx context.Context, id int) (TaskModel, error)
	DeleteTask(ctx context.Context, id int) error
}

func ModeltoDomain(model TaskModel) repository.DomainTask {
	return repository.DomainTask{
		ID:             model.ID,
		Title:          model.Title,
		Description:    model.Description,
		CreationTime:   model.CreationTime,
		CompletionTime: model.CompletionTime,
	}
}

func DomaintoModel(domain repository.DomainTask) TaskModel {
	return TaskModel{
		ID:             domain.ID,
		Title:          domain.Title,
		Description:    domain.Description,
		CreationTime:   domain.CreationTime,
		CompletionTime: domain.CompletionTime,
	}
}

type Service struct {
	rep repository.ResopitoryInterface
}

func NewService(rep repository.ResopitoryInterface) *Service {
	return &Service{rep: rep}
}

func (s *Service) CreateTask(ctx context.Context, m TaskModel) (TaskModel, error) {
	task, err := s.rep.CreateTask(ctx, ModeltoDomain(m))
	return DomaintoModel(task), err
}

// func GetTasks() ([]TaskModel, error) {

// }

// func CompleteTask(id int) (TaskModel, error) {

// }

// func DeleteTask(id int) error {

// }
