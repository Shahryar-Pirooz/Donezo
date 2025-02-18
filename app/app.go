package app

import (
	"context"
	"donezo/config"
	"donezo/internal/project"
	porjectPort "donezo/internal/project/port"
	projectPort "donezo/internal/project/port"
	"donezo/internal/task"
	taskPort "donezo/internal/task/port"
	"donezo/internal/user"
	userPort "donezo/internal/user/port"
	"donezo/pkg/adapter/storage"
	appContext "donezo/pkg/context"
	"donezo/pkg/postgres"

	"gorm.io/gorm"
)

type app struct {
	db             *gorm.DB
	cnf            config.Config
	projectService projectPort.Service
	taskService    taskPort.Service
	userService    userPort.Service
}

func (a *app) ProjectService(ctx context.Context) porjectPort.Service {
	db := appContext.GetDB(ctx)
	if db == nil {
		if a.projectService == nil {
			a.projectService = a.projectServiceWithDB(a.db)
		}
		return a.projectService
	}
	return a.projectServiceWithDB(db)
}
func (a *app) projectServiceWithDB(db *gorm.DB) porjectPort.Service {
	return project.NewService(storage.NewProjectRepo(db))
}
func (a *app) TaskService(ctx context.Context) taskPort.Service {
	db := appContext.GetDB(ctx)
	if db == nil {
		if a.taskService == nil {
			a.taskService = a.taskServiceWithDB(a.db)
		}
		return a.taskService
	}
	return a.taskServiceWithDB(db)
}
func (a *app) taskServiceWithDB(db *gorm.DB) taskPort.Service {
	return task.NewService(storage.NewTaskRepo(db))
}
func (a *app) UserService(ctx context.Context) userPort.Service {
	db := appContext.GetDB(ctx)
	if db == nil {
		if a.userService == nil {
			a.userService = a.userServiceWithDB(a.db)
		}
		return a.userService
	}
	return a.userServiceWithDB(db)
}
func (a *app) userServiceWithDB(db *gorm.DB) userPort.Service {
	return user.NewService(storage.NewUserRepo(db))
}
func (a *app) DB() *gorm.DB {
	return a.db
}
func (a *app) setDB() error {
	db, err := postgres.NewConnection(postgres.DBConnectOption{
		User:     a.cnf.DB.User,
		Password: a.cnf.DB.Password,
		Host:     a.cnf.DB.Host,
		Port:     a.cnf.DB.Port,
		DBName:   a.cnf.DB.DBName,
		SSLMode:  a.cnf.DB.SSLMode,
	})
	if err != nil {
		return err
	}
	a.db = db
	return nil
}
func (a *app) Config() config.Config {
	return a.cnf
}

func NewApp(cnf config.Config) (App, error) {
	a := &app{
		cnf: cnf,
	}
	if err := a.setDB(); err != nil {
		return nil, err
	}
	return a, nil
}

func NewMustApp(cnf config.Config) App {
	app, err := NewApp(cnf)
	if err != nil {
		panic(err)
	}
	return app
}
