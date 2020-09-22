package main

import (
	micro "github.com/micro/go-micro"
	"log"
	"members/handler"
	members "members/proto/members"
	"members/subscriber"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.members"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	members.RegisterMembersHandler(service.Server(), new(handler.Members))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.members", service.Server(), new(subscriber.Members))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
