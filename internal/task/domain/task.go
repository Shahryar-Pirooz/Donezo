package domain

import (
	"donezo/internal/project/domain"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidID       = errors.New("ID is not valid")
	ErrNilID           = errors.New("ID cannot be a Nil UUID")
	ErrEmptyTitle      = errors.New("title is required")
	ErrUnknownPriority = errors.New("priority cannot be unknown")
)

type TaskID = uuid.UUID

type PriorityType uint8

const (
	PriorityTypeUnknown PriorityType = iota
	PriorityTypeLow
	PriorityTypeMedium
	PriorityTypeHigh
)

type Task struct {
	ID          TaskID
	Title       string
	Description string
	ProjectID   domain.ProjectID
	Done        bool
	Priority    PriorityType
	CreatedAt   time.Time
	DeletedAt   time.Time
}

type TaskFilter struct {
	Title     string
	ProjectID domain.ProjectID
	Priority  PriorityType
	Done      bool
}

func (t *Task) Validate() error {

	if t.ID == uuid.Nil {
		return ErrNilID
	}

	if err := uuid.Validate(t.ID.String()); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidID, err)
	}

	if t.Title == "" {
		return ErrEmptyTitle
	}

	if t.Priority == PriorityTypeUnknown {
		return ErrUnknownPriority
	}

	return nil
}
