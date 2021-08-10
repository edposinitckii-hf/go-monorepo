package service

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// EchoMessage is an echoed message
type EchoMessage struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// echo abstracts entire application with all the dependencies
type echo struct {
	db     *sqlx.DB
	logger *zap.Logger
}

// NewEcho creates new echo service
func NewEcho(db *sqlx.DB, logger *zap.Logger) *echo {
	return &echo{db: db, logger: logger}
}

// Echo executes the application
func (a *echo) Echo(ctx context.Context, message string) (*EchoMessage, error) {
	a.logger.Info("echoing the message", zap.String("message", message))

	if _, err := a.db.QueryContext(ctx, "SELECT 1"); err != nil {
		return nil, err
	}

	return &EchoMessage{
		Message:   message,
		Timestamp: time.Now(),
	}, nil
}
