package nats

import (
	"context"
	"fmt"
	"hash/fnv"
	"time"

	"dift_user_insentive/user-coupon-service/internal/servicecore/queue"

	"github.com/nats-io/nats.go"
)

type HandlerFunc func(ctx context.Context, subject string, payload []byte) error

type JetStreamConsumer struct {
	js    nats.JetStreamContext
	inbox *queue.Inbox
	dlq   *queue.DLQ
}

func NewJetStreamConsumer(js nats.JetStreamContext) *JetStreamConsumer {
	inbox := queue.NewInbox(queue.InboxConfig{
		DedupeEnabled:      false,
		DefaultMaxAttempts: 3,
	})
	dlq, _ := queue.NewDLQ(queue.DLQConfig{
		Storage:         queue.NewInMemoryDLQStorage(),
		DefaultTTL:      24 * time.Hour,
		PoisonThreshold: 3,
		ReplayInbox:     inbox,
	})
	return &JetStreamConsumer{js: js, inbox: inbox, dlq: dlq}
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
			err := c.enqueueAndHandle(ctx, msg, handler)
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

func (c *JetStreamConsumer) enqueueAndHandle(
	ctx context.Context,
	msg *nats.Msg,
	handler HandlerFunc,
) error {
	attempts := 1
	if meta, err := msg.Metadata(); err == nil && meta != nil && meta.NumDelivered > 0 {
		attempts = int(meta.NumDelivered)
	}
	item := &queue.Message{
		ID:          buildMessageID(msg.Subject, msg.Data),
		SenderID:    "jetstream",
		Topic:       msg.Subject,
		Priority:    queue.PriorityHigh,
		Body:        append([]byte(nil), msg.Data...),
		MaxAttempts: 3,
		Attempts:    attempts,
		Timestamp:   time.Now().UTC(),
	}
	if err := c.inbox.Submit(item); err != nil {
		return err
	}
	queued, err := c.inbox.Consume(ctx)
	if err != nil {
		return err
	}
	if err := handler(ctx, queued.Topic, queued.Body); err != nil {
		if queued.Attempts >= queued.MaxAttempts {
			_ = c.dlq.Send(ctx, queued, err)
			return nil
		}
		return err
	}
	return nil
}

func buildMessageID(subject string, payload []byte) string {
	h := fnv.New64a()
	_, _ = h.Write([]byte(subject))
	_, _ = h.Write(payload)
	return fmt.Sprintf("%x", h.Sum64())
}
