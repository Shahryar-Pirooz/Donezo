package port

import (
	"context"
	"donezo/internal/user/domain"
)

type Service interface {
	Create(ctx context.Context, user domain.User) (domain.UserID, error)
	Update(ctx context.Context, id domain.UserID, user domain.User) error
	Delete(ctx context.Context, id domain.UserID) error
	GetInformation(ctx context.Context, id domain.UserID) (*domain.User, error)
}
