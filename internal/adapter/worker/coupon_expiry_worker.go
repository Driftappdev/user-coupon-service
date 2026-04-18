package cron

import (
	"context"
	"log"
	"time"
)

type ExpiryService interface {
	Expire(ctx context.Context) error
}

type Worker struct {
	svc    ExpiryService
	ticker *time.Ticker
	stop   chan struct{}
}

func New(svc ExpiryService) *Worker {

	return &Worker{
		svc:    svc,
		ticker: time.NewTicker(10 * time.Minute),
		stop:   make(chan struct{}),
	}
}

func (w *Worker) Start() {

	go func() {

		for {

			select {

			case <-w.ticker.C:

				err := w.svc.Expire(context.Background())

				if err != nil {
					log.Println("[Worker] expire error:", err)
				}

			case <-w.stop:

				return
			}
		}
	}()
}

func (w *Worker) Stop() {

	w.ticker.Stop()

	close(w.stop)
}

