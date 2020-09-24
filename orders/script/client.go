package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	"go-micro-learn/orders/proto/orders"
)

var (
	ordersSrvClient orders.OrdersService
)

func main() {
	ordersSrvClient = orders.NewOrdersService("go.micro.service.orders",client.DefaultClient)
	memReq := &orders.Request{
		Name: "peanut",
	}
	resp, err := ordersSrvClient.Call(context.TODO(), memReq)
	if err != nil || resp == nil {
		fmt.Println("GRpc members resp err", err)
		return
	}
	fmt.Println(resp.Msg)
}
