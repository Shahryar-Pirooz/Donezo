package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrNameRequired     = errors.New("name is required")
	ErrEmailRequired    = errors.New("email is required")
	ErrInvalidEmail     = errors.New("invalid email format")
	ErrPasswordTooShort = errors.New("password must be at least 6 characters")
)

type UserID = uuid.UUID

type User struct {
	ID        UserID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	DeleteAt  time.Time
}

type UserFilter struct {
	Name     *string
	Email    *string
	FromDate *string
	ToDate   *string
	Page     *int
	Limit    *int
}

func (u *User) Validate() error {
	if strings.TrimSpace(u.Name) == "" {
		return ErrNameRequired
	}

	if strings.TrimSpace(u.Email) == "" {
		return ErrEmailRequired
	}

	if !strings.Contains(u.Email, "@") {
		return ErrInvalidEmail
	}

	if len(u.Password) < 6 {
		return ErrPasswordTooShort
	}

	return nil
}
