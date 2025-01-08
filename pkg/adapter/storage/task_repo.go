package storage

import (
	"context"
	taskDomain "donezo/internal/task/domain"
	taskPort "donezo/internal/task/port"

	"gorm.io/gorm"
)

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) taskPort.Repo {
	return &taskRepo{
		db: db,
	}
}
func (r *taskRepo) Create(ctx context.Context, task taskDomain.Task) (taskDomain.TaskID, error) {
	panic("v any")
}
func (r *taskRepo) Update(ctx context.Context, id taskDomain.TaskID, task taskDomain.Task) error {
	panic("v any")
}
func (r *taskRepo) GetByID(ctx context.Context, id taskDomain.TaskID) (*taskDomain.Task, error) {
	panic("v any")
}
func (r *taskRepo) Filter(ctx context.Context, page, size uint, filter taskDomain.TaskFilter) ([]taskDomain.Task, error) {
	panic("v any")
}
func (r *taskRepo) List(ctx context.Context, page, size uint) ([]taskDomain.Task, error) {
	panic("v any")
}
func (r *taskRepo) Delete(ctx context.Context, id taskDomain.TaskID) error {
	panic("v any")
}
