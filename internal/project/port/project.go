package port

import (
	"context"
	"donezo/internal/project/domain"
)

type Repo interface {
	Create(ctx context.Context, project domain.Project) (domain.ProjectID, error)
	Update(ctx context.Context, id domain.ProjectID, project domain.Project) error
	GetByID(ctx context.Context, id domain.ProjectID) (*domain.Project, error)
	List(ctx context.Context, page, limit uint) ([]domain.Project, error)
	Filter(ctx context.Context, page, limit uint, filter *domain.ProjectFilter) ([]domain.Project, error)
	Delete(ctx context.Context, id domain.ProjectID) error
}
