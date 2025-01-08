package storage

import (
	"context"
	projectDomain "donezo/internal/project/domain"
	projectPort "donezo/internal/project/port"
	"donezo/pkg/adapter/storage/types"
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type projectRepo struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) projectPort.Repo {
	return &projectRepo{
		db: db,
	}
}

func (r *projectRepo) Create(ctx context.Context, project projectDomain.Project) (projectDomain.ProjectID, error) {
	var err error
	record := new(types.Project)
	if err = project.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("failed to validate project: %w", err)
	}
	err = copier.Copy(record, &project)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to copy project data: %w", err)
	}
	result := r.db.Model(record).Create(record)
	if result.Error != nil {
		return uuid.Nil, fmt.Errorf("failed to create project in database: %w", result.Error)
	}
	projectID, err := uuid.Parse(record.ID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to parse project ID: %w", err)
	}
	return projectID, nil
}

func (r *projectRepo) Update(ctx context.Context, id projectDomain.ProjectID, project projectDomain.Project) error {
	var err error
	newRecord := new(types.Project)
	if err = project.Validate(); err != nil {
		return fmt.Errorf("failed to validate project update: %w", err)
	}
	result := r.db.First(newRecord, id.String())
	if result.Error != nil {
		return fmt.Errorf("failed to find project with ID %s: %w", id.String(), result.Error)
	}
	err = copier.Copy(newRecord, &project)
	if err != nil {
		return fmt.Errorf("failed to copy updated project data: %w", err)
	}
	result = r.db.Save(&newRecord)
	if result.Error != nil {
		return fmt.Errorf("failed to save updated project to database: %w", result.Error)
	}
	return nil
}

func (r *projectRepo) GetByID(ctx context.Context, id projectDomain.ProjectID) (*projectDomain.Project, error) {
	record := new(types.Project)
	domainProject := new(projectDomain.Project)
	result := r.db.First(record, id.String())
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find project with ID %s: %w", id.String(), result.Error)
	}
	if err := copier.Copy(domainProject, record); err != nil {
		return nil, fmt.Errorf("failed to copy project data: %w", err)
	}
	return domainProject, nil
}

func (r *projectRepo) List(ctx context.Context, page, limit uint) ([]projectDomain.Project, error) {
	records := new([]types.Project)
	domainProjects := new([]projectDomain.Project)
	offset := (page - 1) * limit
	result := r.db.Limit(int(limit)).Offset(int(offset)).Find(&records)
	if result.Error != nil {
		return []projectDomain.Project{}, fmt.Errorf("failed to list projects: %w", result.Error)
	}
	if err := copier.Copy(domainProjects, records); err != nil {
		return []projectDomain.Project{}, fmt.Errorf("failed to copy projects data: %w", err)
	}
	return *domainProjects, nil
}

func (r *projectRepo) Filter(ctx context.Context, page, limit uint, filter *projectDomain.ProjectFilter) ([]projectDomain.Project, error) {
	records := new([]types.Project)
	domainProjects := new([]projectDomain.Project)
	offset := (page - 1) * limit
	result := r.db.Limit(int(limit)).Offset(int(offset)).Where(filter).Find(&records)
	if result.Error != nil {
		return []projectDomain.Project{}, fmt.Errorf("failed to filter projects: %w", result.Error)
	}
	if err := copier.Copy(domainProjects, records); err != nil {
		return []projectDomain.Project{}, fmt.Errorf("failed to copy filtered projects data: %w", err)
	}
	return *domainProjects, nil
}

func (r *projectRepo) Delete(ctx context.Context, id projectDomain.ProjectID) error {
	record := new(types.Project)
	result := r.db.Delete(record, id.String())
	if result.Error != nil {
		return fmt.Errorf("failed to delete project with ID %s: %w", id.String(), result.Error)
	}
	return nil
}
