package task

import (
	"context"
	projectDomain "donezo/internal/project/domain"
	taskDomain "donezo/internal/task/domain"
	taskPort "donezo/internal/task/port"
)

type service struct {
	repo taskPort.Repo
}

func NewService(repo taskPort.Repo) taskPort.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateNewTask(ctx context.Context, record taskDomain.Task) (taskDomain.TaskID, error) {
	panic("")
}
func (s *service) UpdateTask(ctx context.Context, UUID taskDomain.TaskID, newRecord taskDomain.Task) error {
	panic("")
}
func (s *service) DoneTask(ctx context.Context, UUID taskDomain.TaskID) error {
	panic("")
}
func (s *service) GetTitleByID(ctx context.Context, UUID taskDomain.TaskID) (string, error) {
	panic("")
}
func (s *service) GetTasksByPriority(ctx context.Context, priority taskDomain.PriorityType) ([]taskDomain.Task, error) {
	panic("")
}
func (s *service) GetParent(ctx context.Context, UUID taskDomain.TaskID) (projectDomain.Project, error) {
	panic("")
}
func (s *service) GetTask(ctx context.Context, pageIndex, pageSize uint, filter *taskDomain.TaskFilter) ([]taskDomain.Task, error) {
	panic("")
}
func (s *service) DeleteTask(ctx context.Context, UUID taskDomain.Task) error {
	panic("")
}
