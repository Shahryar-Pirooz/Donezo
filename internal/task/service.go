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

func (s *service) Create(ctx context.Context, record taskDomain.Task) (taskDomain.TaskID, error) {
	panic("")
}
func (s *service) Update(ctx context.Context, id taskDomain.TaskID, newRecord taskDomain.Task) error {
	panic("")
}
func (s *service) MarkDone(ctx context.Context, id taskDomain.TaskID) error {
	panic("")
}
func (s *service) GetTitle(ctx context.Context, id taskDomain.TaskID) (string, error) {
	panic("")
}
func (s *service) ListByPriority(ctx context.Context, priority taskDomain.PriorityType) ([]taskDomain.Task, error) {
	panic("")
}
func (s *service) GetProject(ctx context.Context, id taskDomain.TaskID) (projectDomain.Project, error) {
	panic("")
}
func (s *service) List(ctx context.Context, pageIndex, pageSize uint, filter *taskDomain.TaskFilter) ([]taskDomain.Task, error) {
	panic("")
}
func (s *service) Delete(ctx context.Context, id taskDomain.Task) error {
	panic("")
}
