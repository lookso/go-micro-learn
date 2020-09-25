package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"go-micro-learn/members/proto/members"
)

var (
	membersSrvClient members.MembersService
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	GRpcService := micro.NewService(
		micro.Registry(etcdReg),
	)
	GRpcService.Init()

	membersSrvClient = members.NewMembersService("micro-grpc-members-service", client.DefaultClient)
	memReq := &members.Request{
		Name: "peanut",
	}
	resp, err := membersSrvClient.Call(context.TODO(), memReq)
	if err != nil || resp == nil {
		fmt.Println("GRpc members resp err", err)
		return
	}
	fmt.Println(resp.Msg)
}
