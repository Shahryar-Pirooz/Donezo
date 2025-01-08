package task

import (
	"context"
	projectDomain "donezo/internal/project/domain"
	taskDomain "donezo/internal/task/domain"
	taskPort "donezo/internal/task/port"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrInvalidPage     = errors.New("page number must be greater than 0")
	ErrInvalidPageSize = errors.New("page size must be between 1 and 10")
	ErrInvalidTaskID   = errors.New("invalid task ID format")
	ErrTaskNotFound    = errors.New("task not found")
)

type service struct {
	repo taskPort.Repo
}

func NewService(repo taskPort.Repo) taskPort.Service {
	if repo == nil {
		panic("task repository cannot be nil")
	}
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, task taskDomain.Task) (taskDomain.TaskID, error) {
	if ctx == nil {
		return uuid.Nil, errors.New("context cannot be nil")
	}
	if err := task.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("invalid task: %w", err)
	}
	taskID, err := s.repo.Create(ctx, task)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create task: %w", err)
	}
	return taskID, nil
}

func (s *service) Update(ctx context.Context, id taskDomain.TaskID, task taskDomain.Task) error {
	if ctx == nil {
		return errors.New("context cannot be nil")
	}
	if err := uuid.Validate(id.String()); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidTaskID, err)
	}
	if err := task.Validate(); err != nil {
		return fmt.Errorf("invalid task: %w", err)
	}
	if err := s.repo.Update(ctx, id, task); err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	return nil
}

func (s *service) MarkDone(ctx context.Context, id taskDomain.TaskID) error {
	if ctx == nil {
		return errors.New("context cannot be nil")
	}
	if err := uuid.Validate(id.String()); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidTaskID, err)
	}
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}
	if task == nil {
		return ErrTaskNotFound
	}
	task.Done = true
	if err := s.repo.Update(ctx, id, *task); err != nil {
		return fmt.Errorf("failed to mark task as done: %w", err)
	}
	return nil
}

func (s *service) ListByPriority(ctx context.Context, page, size uint, priority taskDomain.PriorityType) ([]taskDomain.Task, error) {
	if ctx == nil {
		return nil, errors.New("context cannot be nil")
	}
	if err := validatePagination(page, size); err != nil {
		return nil, err
	}

	filter := &taskDomain.TaskFilter{Priority: priority}
	tasks, err := s.repo.Filter(ctx, page, size, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to list tasks by priority: %w", err)
	}
	return tasks, nil
}

func (s *service) GetDone(ctx context.Context, page, size uint) ([]taskDomain.Task, error) {
	if ctx == nil {
		return nil, errors.New("context cannot be nil")
	}
	if err := validatePagination(page, size); err != nil {
		return nil, err
	}

	filter := &taskDomain.TaskFilter{Done: true}
	tasks, err := s.repo.Filter(ctx, page, size, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to get done tasks: %w", err)
	}
	return tasks, nil
}

func (s *service) GetProject(ctx context.Context, id taskDomain.TaskID) (projectDomain.ProjectID, error) {
	if ctx == nil {
		return uuid.Nil, errors.New("context cannot be nil")
	}
	if err := uuid.Validate(id.String()); err != nil {
		return uuid.Nil, fmt.Errorf("%w: %v", ErrInvalidTaskID, err)
	}
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to get task: %w", err)
	}
	if task == nil {
		return uuid.Nil, ErrTaskNotFound
	}
	return task.ProjectID, nil
}

func (s *service) List(ctx context.Context, page, size uint, filter *taskDomain.TaskFilter) ([]taskDomain.Task, error) {
	if ctx == nil {
		return nil, errors.New("context cannot be nil")
	}
	if err := validatePagination(page, size); err != nil {
		return nil, err
	}

	var tasks []taskDomain.Task
	var err error

	if filter == nil {
		tasks, err = s.repo.List(ctx, page, size)
	} else {
		tasks, err = s.repo.Filter(ctx, page, size, filter)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to list tasks: %w", err)
	}
	return tasks, nil
}

func (s *service) Delete(ctx context.Context, id taskDomain.TaskID) error {
	if ctx == nil {
		return errors.New("context cannot be nil")
	}
	if err := uuid.Validate(id.String()); err != nil {
		return fmt.Errorf("%w: %v", ErrInvalidTaskID, err)
	}
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}

func validatePagination(page, size uint) error {
	if page < 1 {
		return ErrInvalidPage
	}
	if size < 1 || size > 10 {
		return ErrInvalidPageSize
	}
	return nil
}
