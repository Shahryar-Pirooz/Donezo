package user

import (
	"context"
	userDomain "donezo/internal/user/domain"
	userPort "donezo/internal/user/port"
)

type service struct {
	repo userPort.Repo
}

func NewService(repo userPort.Repo) userPort.Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, user userDomain.User) (userDomain.UserID, error) {
	panic("v any")
}
func (s *service) Update(ctx context.Context, user userDomain.User) error {
	panic("v any")
}
func (s *service) Delete(ctx context.Context, id string) error {
	panic("v any")
}
func (s *service) GetByID(ctx context.Context, id string) (userDomain.User, error) {
	panic("v any")
}
func (s *service) GetAll(ctx context.Context) ([]userDomain.User, error) {
	panic("v any")
}
