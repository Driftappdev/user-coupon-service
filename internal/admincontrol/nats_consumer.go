package admincontrol

import (
	"context"
	"encoding/json"
)

type NATSConsumer struct{ svc *Service }

func NewNATSConsumer(svc *Service) *NATSConsumer { return &NATSConsumer{svc: svc} }

func (c *NATSConsumer) Handle(ctx context.Context, _ string, payload []byte) error {
	var cmd Command
	if err := json.Unmarshal(payload, &cmd); err != nil {
		return err
	}
	return c.svc.Execute(ctx, cmd)
}

