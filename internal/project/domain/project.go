package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrIDNil   = errors.New("ID cannot be a Nil UUID")
	ErrIDValid = errors.New("ID is not valid")
	ErrParent  = errors.New("Cannot be self Parent")
)

type ProjectID = uuid.UUID

type Project struct {
	ID        ProjectID
	Name      string
	Parent    ProjectID
	CreatedAt time.Time
	DeletedAt time.Time
}

type ProjectFilter struct {
	Name string
}

func (p *Project) IsValid() error {
	if p.ID == uuid.Nil {
		return ErrIDNil
	}
	if err := uuid.Validate(p.ID.String()); err != nil {
		return fmt.Errorf("%w: %w", ErrIDValid, err)
	}
	if p.ID == p.Parent {
		return ErrParent
	}
	return nil
}
