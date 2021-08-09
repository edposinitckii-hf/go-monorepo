package log

import "go.uber.org/zap"

// NewLogger creates new logger
func NewLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}
