package user

import (
	"context"
	userDomain "donezo/internal/user/domain"
	userPort "donezo/internal/user/port"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type service struct {
	repo userPort.Repo
}

var (
	ErrInvalidUserID = errors.New("invalid user ID")
	ErrUserNotFound  = errors.New("user not found")
)

func NewService(repo userPort.Repo) userPort.Service {
	return &service{
		repo: repo,
	}
}
func (s *service) Create(ctx context.Context, user userDomain.User) (userDomain.UserID, error) {
	if ctx == nil {
		return uuid.Nil, errors.New("context cannot be nil")
	}
	if err := user.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("invalid user: %w", err)
	}
	userID, err := s.repo.Create(ctx, user)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create user: %w", err)
	}
	return userID, nil
}

func (s *service) Update(ctx context.Context, id userDomain.UserID, user userDomain.User) error {
	if ctx == nil {
		return errors.New("context cannot be nil")
	}
	if err := uuid.Validate(user.ID.String()); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidUserID, err)
	}
	if err := user.Validate(); err != nil {
		return fmt.Errorf("invalid user: %w", err)
	}
	if err := s.repo.Update(ctx, id, user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (s *service) Delete(ctx context.Context, id userDomain.UserID) error {
	if ctx == nil {
		return errors.New("context cannot be nil")
	}
	if err := uuid.Validate(id.String()); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidUserID, err)
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (s *service) GetInformation(ctx context.Context, id userDomain.UserID) (*userDomain.User, error) {
	if ctx == nil {
		return nil, errors.New("context cannot be nil")
	}
	if err := uuid.Validate(id.String()); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidUserID, err)
	}
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user information: %w", err)
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
