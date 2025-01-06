package port

import (
	"context"
	"donezo/internal/project/domain"
)

type Repo interface{
	Create(ctx context.Context, record domain.Project)(domain.ProjectID , error)
	Update(ctx context.Context, UUID domain.ProjectID , newRecord domain.Project) error
	GetAllProjects(ctx context.Context , pageIndex , pageSize uint)([]domain.Project , error)
	FilterProject(ctx context.Context , pageIndex , pageSize uint , filter domain.ProjectFilter)([]domain.Project , error)
	Delete(ctx context.Context , UUID domain.ProjectID) error
}
