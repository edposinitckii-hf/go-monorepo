package internal

import (
	"context"

	"go.uber.org/zap"
)

type SvcTwoClient interface {
	Do(ctx context.Context) error
}

type service struct {
	client SvcTwoClient
	logger *zap.Logger
}

func NewService(client SvcTwoClient, logger *zap.Logger) *service {
	return &service{client: client, logger: logger}
}

func (s service) Execute(ctx context.Context) error {
	s.logger.Info("calling to a service-two via client")
	return s.client.Do(ctx)
}
