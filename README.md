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


protoc --proto_path=. --micro_out=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto:. --go_out=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto:. proto/posts/posts.proto

### micro cmd
micro server &
micro new members
micro run members
micro status

micro env set local
micro env

列出已经注册的服务
peanut-b360hd3 :: ~ » micro list services
go.micro.service.members
go.micro.service.orders
go.micro.service.products
