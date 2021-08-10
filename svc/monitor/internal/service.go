package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/edposinitckii-hf/monorepo/svc/echo/pkg/http"
	"go.uber.org/zap"
)

type MessageEcho struct {
	Message   string
	Timestamp time.Time
}

type EchoClient interface {
	Do(context.Context, string) (*http.MessageEcho, error)
}

type monitor struct {
	client EchoClient
	logger *zap.Logger
}

func NewService(client EchoClient, logger *zap.Logger) *monitor {
	return &monitor{client: client, logger: logger}
}

func (s monitor) Execute(ctx context.Context, message string) (*MessageEcho, error) {
	s.logger.Info("echo via client")
	echo, err := s.client.Do(ctx, message)
	if err != nil {
		return nil, err
	}

	if echo.Error != "" {
		return nil, fmt.Errorf("failed to echo message: %s", echo.Error)
	}

	echoMessage := MessageEcho{
		Message:   echo.Message,
		Timestamp: echo.Timestamp,
	}

	return &echoMessage, nil
}
