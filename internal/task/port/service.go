package port

import (
	"context"
	projectDomain "donezo/internal/project/domain"
	taskDomain "donezo/internal/task/domain"
)

type Service interface {
	CreateNewTask(ctx context.Context , record taskDomain.Task) (taskDomain.TaskID , error)
	UpdateTask(ctx context.Context , UUID taskDomain.TaskID , newRecord taskDomain.Task) error
	DoneTask(ctx context.Context , UUID taskDomain.TaskID) error
	GetTitleByID(ctx context.Context ,UUID taskDomain.TaskID) (string , error)
	GetTasksByPriority (ctx context.Context , priority taskDomain.PriorityType) ([]taskDomain.Task , error)
	GetParent(ctx context.Context , UUID taskDomain.TaskID) (projectDomain.Project , error)
	GetTask(ctx context.Context, pageIndex , pageSize uint , filter *taskDomain.TaskFilter) ([]taskDomain.Task , error)
	DeleteTask(ctx context.Context , UUID taskDomain.Task) error
}
