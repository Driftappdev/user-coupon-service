package admincontrol

import "context"

type Command struct {
	Action         string         `json:"action"`
	Payload        map[string]any `json:"payload"`
	IdempotencyKey string         `json:"idempotency_key"`
}

type Service struct{}

func NewService() *Service { return &Service{} }
func (s *Service) Execute(_ context.Context, _ Command) error { return nil }

