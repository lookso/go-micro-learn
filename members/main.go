package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"go-micro-learn/members/handler/rpcs"
	"go-micro-learn/members/model"
	"go-micro-learn/members/proto/members"
	"go-micro-learn/members/router"
	"sync"
)

func main() {

	// 启动etcd集群
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	gr := router.InitRouters()
	webService := web.NewService(
		web.Name("micro-members-web-api"),
		web.Address(":8090"),
		web.Handler(gr),
		web.Registry(etcdReg), // 服务注册
	)
	// 查看etcd 注册信息
	//# etcdctl get /micro prefix
	// micro/registry/micro-members-web-api/0af17e80-48f1-4977-978b-8e0cf0aed164
	//{"name":"micro-members-web-api","version":"latest","metadata":null,"endpoints":null, "nodes":[{"id":"0af17e80-48f1-4977-978b-8e0cf0aed164","address":"10.60.103.159:8090","metadata":null}]}

	// web服务初始化
	webService.Init(
		web.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
		}),
	)

	// grpc服务
	//New Service
	grpcService := micro.NewService(
		micro.Name("go.micro.service.members"),
		micro.Version("latest"),
		micro.Registry(etcdReg), // 服务注册
	)
	// Initialise service
	grpcService.Init()

	//Register Handler
	members.RegisterMembersHandler(grpcService.Server(), new(rpcs.Members))

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		// Run service
		if err := grpcService.Run(); err != nil {
			log.Fatal(err)
		}
	}()
	go func() {
		defer wg.Done()
		if err := webService.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
}
