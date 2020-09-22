package main

import (
	log "github.com/micro/go-micro/v3/logger"
	"github.com/micro/go-micro"
	"orders/handler"
	"orders/subscriber"

	orders "orders/proto/orders"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.orders"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	orders.RegisterOrdersHandler(service.Server(), new(handler.Orders))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.orders", service.Server(), new(subscriber.Orders))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
