package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"go-micro-learn/orders/handler"
	"go-micro-learn/orders/proto/orders"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.orders"),
		micro.Version("latest"),
		micro.Registry(etcdReg),
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
