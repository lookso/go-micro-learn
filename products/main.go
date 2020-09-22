package main

import (
	log "github.com/micro/go-micro/v3/logger"
	"github.com/micro/go-micro"
	"products/handler"
	"products/subscriber"

	products "products/proto/products"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.products"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	products.RegisterProductsHandler(service.Server(), new(handler.Products))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.products", service.Server(), new(subscriber.Products))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
