package project

import (
	"context"
	projectDomain "donezo/internal/project/domain"
	"donezo/internal/project/port"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, record projectDomain.Project) (projectDomain.ProjectID, error) {
	if err := record.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("invalid project: %w", err)
	}
	newID, err := s.repo.Create(ctx, record)
	if err != nil {
		return newID, err
	}
	return newID, nil
}

func (s *service) Update(ctx context.Context, UUID projectDomain.ProjectID, newRecord projectDomain.Project) error {
	if err := uuid.Validate(UUID.String()); err != nil {
		return fmt.Errorf("invalid UUID: %w", err)
	}

	if err := newRecord.Validate(); err != nil {
		return fmt.Errorf("invalid project: %w", err)
	}
	if err := s.repo.Update(ctx, UUID, newRecord); err != nil {
		return fmt.Errorf("something went wrong in update: %w", err)
	}
	return nil
}

func (s *service) List(ctx context.Context, page uint, size uint, filter *projectDomain.ProjectFilter) ([]projectDomain.Project, error) {
	if page < 1 {
		return []projectDomain.Project{}, errors.New("page number must be greater than 0")
	}
	if size < 1 {
		return []projectDomain.Project{}, errors.New("page size must be greater than 0")
	}
	if size > 10 {
		return []projectDomain.Project{}, errors.New("page size cannot exceed 10")
	}
	if filter != nil {
		projects, err := s.repo.Filter(ctx, page, size, filter)
		if err != nil {
			return nil, fmt.Errorf("error filtering projects: %w", err)
		}
		return projects, nil
	}

	projects, err := s.repo.List(ctx, page, size)
	if err != nil {
		return nil, fmt.Errorf("error getting projects: %w", err)
	}
	return projects, nil
}

func (s *service) Delete(ctx context.Context, UUID projectDomain.ProjectID) error {
	if err := uuid.Validate(UUID.String()); err != nil {
		return fmt.Errorf("invalid UUID: %w", err)
	}
	if err := s.repo.Delete(ctx, UUID); err != nil {
		return fmt.Errorf("something went wrong in delete: %w", err)
	}
	return nil
}
