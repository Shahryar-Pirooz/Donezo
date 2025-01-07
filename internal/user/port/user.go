package port

import (
	"context"
	"donezo/internal/user/domain"
)

type Repo interface {
	Create(ctx context.Context, user domain.User) (domain.UserID, error)
	Update(ctx context.Context, id domain.UserID, user domain.User) error
	GetByID(ctx context.Context, id domain.UserID) (*domain.User, error)
	List(ctx context.Context, page, limit uint) ([]domain.User, error)
	Filter(ctx context.Context, page, limit uint, filter *domain.UserFilter) ([]domain.User, error)
	Delete(ctx context.Context, id domain.UserID) error
}
