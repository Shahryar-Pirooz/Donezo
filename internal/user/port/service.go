package port

import (
	"context"
	"donezo/internal/user/domain"
)

type Service interface {
	Create(ctx context.Context, user domain.User) (domain.UserID, error)
	Update(ctx context.Context, user domain.User) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (domain.User, error)
	GetAll(ctx context.Context) ([]domain.User, error)
}
