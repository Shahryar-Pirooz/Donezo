package port

import (
	"context"
	"donezo/internal/task/domain"
)

type Repo interface {
	Create(ctx context.Context, task domain.Task) (domain.TaskID, error)
	Update(ctx context.Context, id domain.TaskID, task domain.Task) error
	List(ctx context.Context, page, limit uint) ([]domain.Task, error)
	Filter(ctx context.Context, page, limit uint, filter domain.TaskFilter) ([]domain.Task, error)
	Delete(ctx context.Context, id domain.TaskID) error
}
