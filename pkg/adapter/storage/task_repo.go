package storage

import (
	"context"
	taskDomain "donezo/internal/task/domain"
	taskPort "donezo/internal/task/port"
	"donezo/pkg/adapter/storage/types"
	"fmt"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
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
	var err error
	record := new(types.Task)
	if err = task.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("failed to validate task: %w", err)
	}
	err = copier.Copy(record, &task)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to copy task data: %w", err)
	}
	result := r.db.Model(record).Create(record)
	if result.Error != nil {
		return uuid.Nil, fmt.Errorf("failed to create task in database: %w", result.Error)
	}
	taskID, err := uuid.Parse(record.ID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to parse task ID: %w", err)
	}
	return taskID, nil
}

func (r *taskRepo) Update(ctx context.Context, id taskDomain.TaskID, task taskDomain.Task) error {
	var err error
	newRecord := new(types.Task)
	if err = task.Validate(); err != nil {
		return fmt.Errorf("failed to validate task update: %w", err)
	}
	result := r.db.First(newRecord, id.String())
	if result.Error != nil {
		return fmt.Errorf("failed to find task with ID %s: %w", id.String(), result.Error)
	}
	err = copier.Copy(newRecord, &task)
	if err != nil {
		return fmt.Errorf("failed to copy updated task data: %w", err)
	}
	result = r.db.Save(&newRecord)
	if result.Error != nil {
		return fmt.Errorf("failed to save updated task to database: %w", result.Error)
	}
	return nil
}

func (r *taskRepo) GetByID(ctx context.Context, id taskDomain.TaskID) (*taskDomain.Task, error) {
	record := new(types.Task)
	domainTask := new(taskDomain.Task)
	result := r.db.First(record, id.String())
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find task with ID %s: %w", id.String(), result.Error)
	}
	if err := copier.Copy(domainTask, record); err != nil {
		return nil, fmt.Errorf("failed to copy task data: %w", err)
	}
	return domainTask, nil
}

func (r *taskRepo) List(ctx context.Context, page, limit uint) ([]taskDomain.Task, error) {
	records := new([]types.Task)
	domainTasks := new([]taskDomain.Task)
	offset := (page - 1) * limit
	result := r.db.Limit(int(limit)).Offset(int(offset)).Find(&records)
	if result.Error != nil {
		return []taskDomain.Task{}, fmt.Errorf("failed to list tasks: %w", result.Error)
	}
	if err := copier.Copy(domainTasks, records); err != nil {
		return []taskDomain.Task{}, fmt.Errorf("failed to copy tasks data: %w", err)
	}
	return *domainTasks, nil
}

func (r *taskRepo) Filter(ctx context.Context, page, limit uint, filter *taskDomain.TaskFilter) ([]taskDomain.Task, error) {
	records := new([]types.Task)
	domainTasks := new([]taskDomain.Task)
	offset := (page - 1) * limit
	result := r.db.Limit(int(limit)).Offset(int(offset)).Where(filter).Find(&records)
	if result.Error != nil {
		return []taskDomain.Task{}, fmt.Errorf("failed to filter tasks: %w", result.Error)
	}
	if err := copier.Copy(domainTasks, records); err != nil {
		return []taskDomain.Task{}, fmt.Errorf("failed to copy filtered tasks data: %w", err)
	}
	return *domainTasks, nil
}

func (r *taskRepo) Delete(ctx context.Context, id taskDomain.TaskID) error {
	record := new(types.Task)
	result := r.db.Delete(record, id.String())
	if result.Error != nil {
		return fmt.Errorf("failed to delete task with ID %s: %w", id.String(), result.Error)
	}
	return nil
}
