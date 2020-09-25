package main

import (
	"fmt"

	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
)

func main() {
	registry := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	webService, err := registry.GetService("micro-grpc-members-service")
	// 从etcd中获取服务信息
	//Address:10.60.103.159:8090 , Id:0af17e80-48f1-4977-978b-8e0cf0aed164 ,Metadata:map[]
	if err != nil {
		fmt.Printf("get service err %v", err)
		return
	}
	//随机方式，返回的 next 是一个没有参数的函数
	next := selector.Random(webService)
	//node 里面有 服务的 地址 ID 元数据
	node, err := next()
	fmt.Printf("Address:%v , Id:%v ,Metadata:%v\n", node.Address, node.Id, node.Metadata)
}
