# micro-server v2版本
微服务架构实践

### 目录结构,每个目录代表一个独立的系统服务

```
├── members  会员服务
│   ├── handler
│   ├── proto
│   │   └── members
│   └── subscriber
├── orders  订单系统
│   ├── handler
│   ├── proto
│   │   └── orders
│   └── subscriber
└── products  商品系统
    ├── handler
    ├── proto
    │   └── products
    └── subscriber

```

### 接入计划
- [ ] 引入ZipKin 分布式链路追踪
- [ ] GRpc

```
生成pb文件和micro文件
make proto
or
protoc --proto_path=. --micro_out=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto:. --go_out=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto:. proto/members/members.proto

### micro cmd
micro server &
生成模板文件
micro new members
subscriber是用来订阅服务的,目前用不到可以删除,然后可以删除go.mod之类的文件,自己go mod init去生成
micro run members
micro status

micro env set local
» micro env
* local      none
  server     127.0.0.1:8081
  platform   proxy.micro.mu

列出已经注册的服务
» micro list services
go.micro.service.members
go.micro.service.orders
go.micro.service.products
```

参考文档:
https://xueyuanjun.com/post/20965#toc-2
https://github.com/micro-in-cn/tutorials/tree/master/microservice-in-micro