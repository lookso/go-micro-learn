package main

import (
	micro "github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"go-micro-learn/products/handler"

	products "go-micro-learn/products/proto/products"
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

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
