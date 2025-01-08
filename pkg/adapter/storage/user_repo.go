package storage

import (
	"context"
	userDomain "donezo/internal/user/domain"
	userPort "donezo/internal/user/port"
	"donezo/pkg/adapter/storage/types"
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
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
	var err error
	record := new(types.User)
	if err = user.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("failed to validate user: %w", err)
	}
	err = copier.Copy(record, &user)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to copy user data: %w", err)
	}
	result := r.db.Model(record).Create(record)
	if result.Error != nil {
		return uuid.Nil, fmt.Errorf("failed to create user in database: %w", result.Error)
	}
	userID, err := uuid.Parse(record.ID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to parse user ID: %w", err)
	}
	return userID, nil
}

func (r *userRepo) Update(ctx context.Context, id userDomain.UserID, user userDomain.User) error {
	var err error
	newRecord := new(types.User)
	if err = user.Validate(); err != nil {
		return fmt.Errorf("failed to validate user update: %w", err)
	}
	result := r.db.First(newRecord, id.String())
	if result.Error != nil {
		return fmt.Errorf("failed to find user with ID %s: %w", id.String(), result.Error)
	}
	err = copier.Copy(newRecord, &user)
	if err != nil {
		return fmt.Errorf("failed to copy updated user data: %w", err)
	}
	result = r.db.Save(&newRecord)
	if result.Error != nil {
		return fmt.Errorf("failed to save updated user to database: %w", result.Error)
	}
	return nil
}

func (r *userRepo) GetByID(ctx context.Context, id userDomain.UserID) (*userDomain.User, error) {
	record := new(types.User)
	domainUser := new(userDomain.User)
	result := r.db.First(record, id.String())
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find user with ID %s: %w", id.String(), result.Error)
	}
	if err := copier.Copy(domainUser, record); err != nil {
		return nil, fmt.Errorf("failed to copy user data: %w", err)
	}
	return domainUser, nil
}

func (r *userRepo) List(ctx context.Context, page, limit uint) ([]userDomain.User, error) {
	records := new([]types.User)
	domainUsers := new([]userDomain.User)
	offset := (page - 1) * limit
	result := r.db.Limit(int(limit)).Offset(int(offset)).Find(&records)
	if result.Error != nil {
		return []userDomain.User{}, fmt.Errorf("failed to list users: %w", result.Error)
	}
	if err := copier.Copy(domainUsers, records); err != nil {
		return []userDomain.User{}, fmt.Errorf("failed to copy users data: %w", err)
	}
	return *domainUsers, nil
}

func (r *userRepo) Filter(ctx context.Context, page, limit uint, filter *userDomain.UserFilter) ([]userDomain.User, error) {
	records := new([]types.User)
	domainUsers := new([]userDomain.User)
	offset := (page - 1) * limit
	result := r.db.Limit(int(limit)).Offset(int(offset)).Where(filter).Find(&records)
	if result.Error != nil {
		return []userDomain.User{}, fmt.Errorf("failed to filter users: %w", result.Error)
	}
	if err := copier.Copy(domainUsers, records); err != nil {
		return []userDomain.User{}, fmt.Errorf("failed to copy filtered users data: %w", err)
	}
	return *domainUsers, nil
}

func (r *userRepo) Delete(ctx context.Context, id userDomain.UserID) error {
	record := new(types.User)
	result := r.db.Delete(record, id.String())
	if result.Error != nil {
		return fmt.Errorf("failed to delete user with ID %s: %w", id.String(), result.Error)
	}
	return nil
}
