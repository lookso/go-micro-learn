package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	products "products/proto/products"
)

type Products struct{}

func (e *Products) Handle(ctx context.Context, msg *products.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *products.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
