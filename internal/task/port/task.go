package port

import (
	"context"
	"donezo/internal/task/domain"
)

type Repo interface{
	Create(ctx context.Context, record domain.Task)(domain.TaskID , error)
	Update(ctx context.Context, UUID domain.TaskID , newRecord domain.Task) error
	GetAllTasks(ctx context.Context , pageIndex , pageSize uint)([]domain.Task , error)
	FilterTask(ctx context.Context , pageIndex , pageSize uint , filter domain.TaskFilter)([]domain.Task , error)
	Delete(ctx context.Context , UUID domain.TaskID) error
}		
