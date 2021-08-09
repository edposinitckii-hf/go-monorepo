package pkg

import "context"

// SvcClient represents client for service-two Server
type SvcClient struct{}

func NewSvcClient() *SvcClient {
	return &SvcClient{}
}

// Do makes call to a service
func (c SvcClient) Do(ctx context.Context) error {
	return nil
}
