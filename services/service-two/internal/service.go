package internal

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// app abstracts entire application with all the dependencies
type app struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func NewApp(db *sqlx.DB, logger *zap.Logger) *app {
	return &app{db: db, logger: logger}
}

// Execute executes the application
func (a *app) Execute(ctx context.Context) error {
	a.logger.Info("executing application")
	_, e := a.db.QueryContext(ctx, "SELECT 1")

	return e
}
