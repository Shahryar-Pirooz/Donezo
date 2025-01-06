package port

import (
	"context"
	projectDomain "donezo/internal/project/domain"
	taskDomain "donezo/internal/task/domain"
)

type Service interface {
	Create(ctx context.Context, task taskDomain.Task) (taskDomain.TaskID, error)
	Update(ctx context.Context, id taskDomain.TaskID, task taskDomain.Task) error
	MarkDone(ctx context.Context, id taskDomain.TaskID) error
	GetTitle(ctx context.Context, id taskDomain.TaskID) (string, error)
	ListByPriority(ctx context.Context, priority taskDomain.PriorityType) ([]taskDomain.Task, error)
	GetProject(ctx context.Context, id taskDomain.TaskID) (projectDomain.Project, error)
	List(ctx context.Context, page, size uint, filter *taskDomain.TaskFilter) ([]taskDomain.Task, error)
	Delete(ctx context.Context, task taskDomain.Task) error
}
