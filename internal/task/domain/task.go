package domain

import (
	"donezo/internal/project/domain"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Task related errors
var (
	ErrInvalidID       = errors.New("ID is not valid")
	ErrNilID           = errors.New("ID cannot be a Nil UUID")
	ErrEmptyTitle      = errors.New("title is required")
	ErrUnknownPriority = errors.New("priority cannot be unknown")
)

// TaskID represents the unique identifier for a task
type TaskID = uuid.UUID

// PriorityType represents the priority level of a task
type PriorityType uint8

// Priority levels
const (
	PriorityTypeUnknown PriorityType = iota
	PriorityTypeLow
	PriorityTypeMedium
	PriorityTypeHigh
)

// Task represents a work item that needs to be completed
type Task struct {
	ID          TaskID // Renamed from UUID for clarity
	Title       string
	Description string           // Renamed from Describe for clarity
	ProjectID   domain.ProjectID // Renamed from Parent for clarity
	Done        bool
	Priority    PriorityType
	CreatedAt   time.Time // Fixed typo in field name
	DeletedAt   time.Time // Fixed typo in field name
}

// TaskFilter contains criteria for filtering tasks
type TaskFilter struct {
	Title     string
	ProjectID domain.ProjectID // Renamed from Parent for consistency
	Done      bool
}

// Valid checks if the task has all required fields properly set
func (t *Task) Valid() error {
	// Check for nil UUID
	if t.ID == uuid.Nil {
		return ErrNilID
	}

	// Validate UUID format
	if err := uuid.Validate(t.ID.String()); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidID, err)
	}

	// Validate title
	if t.Title == "" {
		return ErrEmptyTitle
	}

	// Validate priority
	if t.Priority == PriorityTypeUnknown {
		return ErrUnknownPriority
	}

	return nil
}
