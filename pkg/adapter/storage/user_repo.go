package storage

import (
	"context"
	userDomain "donezo/internal/user/domain"
	userPort "donezo/internal/user/port"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) userPort.Repo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, user userDomain.User) (userDomain.UserID, error) {
	panic("v any")
}
func (r *userRepo) Update(ctx context.Context, id userDomain.UserID, user userDomain.User) error {
	panic("v any")
}
func (r *userRepo) GetByID(ctx context.Context, id userDomain.UserID) (*userDomain.User, error) {
	panic("v any")
}
func (r *userRepo) List(ctx context.Context, page, limit uint) ([]userDomain.User, error) {
	panic("v any")
}
func (r *userRepo) Filter(ctx context.Context, page, limit uint, filter *userDomain.UserFilter) ([]userDomain.User, error) {
	panic("v any")
}
func (r *userRepo) Delete(ctx context.Context, id userDomain.UserID) error {
	panic("v any")
}
