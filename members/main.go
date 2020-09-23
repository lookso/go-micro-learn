package main

import (
	micro "github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"go-micro-learn/members/handler"
	members "go-micro-learn/members/proto/members"
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

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
