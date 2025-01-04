package port

import (
	"context"
	projectDomain "donezo/internal/project/domain"

)

type Service interface {
	CreateProject(ctx context.Context , record projectDomain.Project) (projectDomain.ProjectID , error)
	UpdateProject(ctx context.Context , uuid projectDomain.ProjectID , newRecord projectDomain.Project) error
	GetProject(ctx context.Context , pageIndex uint , pageSize uint , filters ...projectDomain.ProjectFilter) ([]projectDomain.Project , error)
	DeleteProject(ctx context.Context , uuid projectDomain.ProjectID) error
}
