package app

import (
	"context"
	"donezo/config"
	porjectPort "donezo/internal/project/port"
	taskPort "donezo/internal/task/port"
	userPort "donezo/internal/user/port"

	"gorm.io/gorm"
)

type App interface {
	ProjectService(ctx context.Context) porjectPort.Service
	TaskService(ctx context.Context) taskPort.Service
	UserService(ctx context.Context) userPort.Service
	DB() *gorm.DB
	Config() config.Config
}
