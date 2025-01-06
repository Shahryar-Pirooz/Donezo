package project

import (
	"context"
	projectDomain "donezo/internal/project/domain"
	"donezo/internal/project/port"
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

func (s *service) CreateProject(ctx context.Context, record projectDomain.Project) (projectDomain.ProjectID, error) {
	if err:= record.IsValid() ; err!=nil{
		return uuid.Nil , fmt.Errorf("%w" , err)
	}
	projectID , err := s.repo.Create(ctx,record)
	return projectID , err
}

func (s *service) UpdateProject(ctx context.Context, UUID projectDomain.ProjectID, newRecord projectDomain.Project) error {
	if err:= uuid.Validate(UUID.String()); err!=nil{
		return fmt.Errorf("%w" , err)
	}
	if err := s.repo.Update(ctx , UUID , newRecord); err!=nil {
		return fmt.Errorf("%w" , err)
	}
	return nil
}

func (s *service) GetProject(ctx context.Context, pageIndex uint, pageSize uint, filter *projectDomain.ProjectFilter) ([]projectDomain.Project, error) {
	var projects []projectDomain.Project
	var err error
	if filter != nil{
		projects , err = s.repo.FilterProject(ctx , pageIndex , pageSize , *filter)
	}else {
		projects , err = s.repo.GetAllProjects(ctx,pageIndex , pageSize)
	}
	if err!=nil{ 
		return projects , fmt.Errorf("%w" , err)
	}
	return projects , err 
}

func (s *service) DeleteProject(ctx context.Context, UUID projectDomain.ProjectID) error {
 if err := uuid.Validate(UUID.String()); err!=nil{
	 return fmt.Errorf("%w" , err)
 }
 if err := s.repo.Delete(ctx , UUID); err!=nil{
 
	 return fmt.Errorf("%w" , err)
 }
 return nil
}
