package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v3/logger"

	orders "orders/proto/orders"
)

type Orders struct{}

func (e *Orders) Handle(ctx context.Context, msg *orders.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *orders.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
