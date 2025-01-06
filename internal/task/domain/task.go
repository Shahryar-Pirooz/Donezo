package domain

import (
	"donezo/internal/project/domain"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ERROR_ID_VALID         = errors.New("ID is not valid")
	ERROR_ID_NIL           = errors.New("ID cannot be a Nil UUID")
	ERROR_TITLE_NIL        = errors.New("title is required")
	ERROR_PRIORITY_UNKNOWN = errors.New("priority cannot be unknown")
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
	UUID     TaskID
	Title    string
	Describe string
	Parent   domain.ProjectID
	Done     bool
	Priority PriorityType
	CreateAt time.Time
	DeleteAt time.Time
}

type TaskFilter struct {
	Title  string
	Parent domain.ProjectID
	Done   bool
}

func (t *Task) Valid() error {
	if err := uuid.Validate(t.UUID.String()); err != nil {
		return ERROR_ID_NIL
	}
	if err := uuid.Validate(t.UUID.String()); err != nil {
		return fmt.Errorf("%w : %w", ERROR_ID_VALID, err)
	}
	if t.Title == "" {
		return ERROR_TITLE_NIL
	}
	if t.Priority == PriorityTypeUnknown {
		return ERROR_PRIORITY_UNKNOWN
	}
	return nil
}
