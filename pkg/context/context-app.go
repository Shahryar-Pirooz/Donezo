package context

import (
	"context"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type appContext struct {
	context.Context
	db     *gorm.DB
	logger zerolog.Logger
}
type AppContextOpt func(*appContext) *appContext

var defaultLogger zerolog.Logger

func WithDB(db *gorm.DB) AppContextOpt {
	return func(ac *appContext) *appContext {
		ac.db = db
		return ac
	}
}

func WithLogger(logger zerolog.Logger) AppContextOpt {
	return func(ac *appContext) *appContext {
		ac.logger = logger
		return ac
	}
}

func NewAppContext(parrent context.Context, opts ...AppContextOpt) context.Context {
	ctx := &appContext{
		Context: parrent,
	}
	for _, opt := range opts {
		ctx = opt(ctx)
	}
	return ctx
}

func SetDB(ctx context.Context, db *gorm.DB) {
	ctxApp, ok := ctx.(*appContext)
	if !ok {
		return
	}
	ctxApp.db = db
}
func GetDB(ctx context.Context) *gorm.DB {
	ctxApp, ok := ctx.(*appContext)
	if !ok {
		return nil
	}
	return ctxApp.db
}
func SetLogger(ctx context.Context, logger zerolog.Logger) {
	ctxApp, ok := ctx.(*appContext)
	if !ok {
		return
	}
	ctxApp.logger = logger
}

func GetLogger(ctx context.Context) zerolog.Logger {
	ctxApp, ok := ctx.(*appContext)
	if !ok {
		return defaultLogger
	}
	return ctxApp.logger
}
