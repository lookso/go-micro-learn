# micro-server v2版本
微服务架构实践

### 目录结构,每个目录代表一个独立的系统服务

```
├── members  会员服务
│   ├── handler
│   ├── proto
│   │   └── members
├── orders  订单系统
│   ├── handler
│   ├── proto
│   │   └── orders
└── products  商品系统
    ├── handler
    ├── proto
    │   └── products

系统调用关系
members  grpc提供 user/info 调用 orders/list 查看订单列表
products 调用 user/info grpc提供 product/info
orders   调用 user/info和product/info grpc提供 order/list
```

### 接入计划
- [x] Gin框架
- [ ] 引入ZipKin 分布式链路追踪
- [ ] GRpc,Gateway
- [ ] Etcd
- [ ] ApiGateway

```
生成pb文件和micro文件
cd members && make proto
or
protoc --proto_path=. --micro_out=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto:. --go_out=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto:. proto/members/members.proto

### micro cmd

micro server &
micro env set local
» micro env
* local      none
  server     127.0.0.1:8081
  platform   proxy.micro.mu

生成模板文件
micro new members
subscriber是用来订阅服务的,目前用不到可以删除,然后可以删除go.mod之类的文件,自己go mod init去生成,v3版本默认已经去掉这个目录
micro run members
micro status

列出已经注册的服务 或者 micro web 通过浏览器界面查看
» micro list services
go.micro.service.members
go.micro.service.orders
go.micro.service.products
```

micro参考文档:

- https://xueyuanjun.com/post/20965#toc-2
- https://github.com/micro-in-cn/tutorials/tree/master/microservice-in-micro