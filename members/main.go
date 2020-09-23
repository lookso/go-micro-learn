package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"go-micro-learn/members/handler"
	"go-micro-learn/members/proto/members"
	"go-micro-learn/members/router"
)

func main() {
	// 服务发现
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	gr := router.InitRouters()
	webService := web.NewService(
		web.Name("micro-members-web-api"),
		web.Address(":8090"),
		web.Handler(gr),
		web.Registry(etcdReg),
	)
	webService.Run()

	// grpc服务
	//New Service
	service := micro.NewService(
		micro.Name("go.micro.service.members"),
		micro.Version("latest"),
	)
	// Initialise service
	service.Init()
	//Register Handler
	members.RegisterMembersHandler(service.Server(), new(handler.Members))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
