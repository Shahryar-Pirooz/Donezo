package project

import (
	"context"
	projectDomain "donezo/internal/project/domain"
	"donezo/internal/project/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateProject(ctx context.Context, record projectDomain.Project) (projectDomain.ProjectID, error) {
	panic("")
}
func (s *service) UpdateProject(ctx context.Context, uuid projectDomain.ProjectID, newRecord projectDomain.Project) error {
	panic("")
}

func (s *service) GetProject(ctx context.Context, pageIndex uint, pageSize uint, filters ...projectDomain.ProjectFilter) ([]projectDomain.Project, error) {
	panic("")
}

func (s *service) DeleteProject(ctx context.Context, uuid projectDomain.ProjectID) error {
	panic("")
}
