package port

import (
	"context"
	projectDomain "donezo/internal/project/domain"
)

type Service interface {
	Create(ctx context.Context, project projectDomain.Project) (projectDomain.ProjectID, error)
	Update(ctx context.Context, id projectDomain.ProjectID, project projectDomain.Project) error
	List(ctx context.Context, page, size uint, filter *projectDomain.ProjectFilter) ([]projectDomain.Project, error)
	Delete(ctx context.Context, id projectDomain.ProjectID) error
}
