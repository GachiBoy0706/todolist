package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DomainTask struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type ResopitoryInterface interface {
	CreateTask(ctx context.Context, task DomainTask) (DomainTask, error)
	// GetTasks(ctx context.Context) ([]DomainTask, error)
	// CompleteTask(ctx context.Context, id int) (DomainTask, error)
	// DeleteTask(ctx context.Context, id int) error
}

type Repo struct {
	pool *pgxpool.Pool
}

func NewRepo(pool *pgxpool.Pool) *Repo {
	return &Repo{pool: pool}
}

func (r *Repo) CreateTask(ctx context.Context, task DomainTask) (DomainTask, error) {
	query := `INSERT INTO tasks (title, description) 
	VALUES ($1, $2) 
	RETURNING id, title, description, created_at, completed_at`

	row := r.pool.QueryRow(ctx, query, task.Title, task.Description)
	var taskReturned DomainTask
	err := row.Scan(&taskReturned.ID, &taskReturned.Title, &taskReturned.Description, &taskReturned.CreatedAt, &taskReturned.CompletedAt)
	if err != nil {
		return DomainTask{}, err
	}
	return taskReturned, nil
}
