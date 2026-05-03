package natsadmin

import (
	"context"
	"encoding/json"

	adminport "dift_user_insentive/user-coupon-service/internal/interface/service/admin"
	adminhandler "dift_user_insentive/user-coupon-service/internal/service_logic/handler/admin"
)

type CommandConsumer struct {
	handler *adminhandler.CommandHandler
}

func NewCommandConsumer(handler *adminhandler.CommandHandler) *CommandConsumer {
	return &CommandConsumer{handler: handler}
}

func (c *CommandConsumer) Handle(ctx context.Context, _ string, payload []byte) error {
	var cmd adminport.Command
	if err := json.Unmarshal(payload, &cmd); err != nil {
		return err
	}
	return c.handler.Execute(ctx, cmd)
}

