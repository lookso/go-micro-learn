package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v3/logger"

	members "members/proto/members"
)

type Members struct{}

func (e *Members) Handle(ctx context.Context, msg *members.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *members.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
