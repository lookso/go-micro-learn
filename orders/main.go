package main

import (
	log "github.com/micro/go-micro/v2/logger"
	micro "github.com/micro/go-micro/v2"
	"go-micro-learn/orders/handler"

	orders "go-micro-learn/orders/proto/orders"
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

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
