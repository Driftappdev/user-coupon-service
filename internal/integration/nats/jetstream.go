package nats

import (
	"context"

	"github.com/nats-io/nats.go"
)

type HandlerFunc func(ctx context.Context, subject string, payload []byte) error

type JetStreamConsumer struct {
	js nats.JetStreamContext
}

func NewJetStreamConsumer(js nats.JetStreamContext) *JetStreamConsumer {
	return &JetStreamConsumer{js: js}
}

func (c *JetStreamConsumer) Subscribe(
	ctx context.Context,
	subject string,
	durable string,
	handler HandlerFunc,
) error {

	_, err := c.js.Subscribe(
		subject,
		func(msg *nats.Msg) {

			err := handler(ctx, msg.Subject, msg.Data)

			if err != nil {
				msg.Nak()
				return
			}

			msg.Ack()
		},
		nats.Durable(durable),
		nats.ManualAck(),
	)

	return err
}

