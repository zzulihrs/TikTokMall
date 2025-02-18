# TikTokMall

# 开发环境配置
* go: 1.23.1
* protoc: 3.6.1
```bash
# !!! 自行设置 GOPATH
rm $GOPATH/bin/air
rm $GOPATH/bin/cwgo
# air
go install github.com/air-verse/air@v1.61.7
# cwgo
go install github.com/cloudwego/cwgo@v0.1.2
```
Docker Images
```bash
docker pull consul:1.15.4
docker pull mysql:8
docker pull redis:7.2.4-alpine
docker pull nats:2.10.24-alpine
docker pull prom/prometheus:v3.0.0
docker pull grafana/grafana:11.4.0
docker pull jaegertracing/all-in-one:1.64.0
docker pull docker.io/bitnami/etcd:3.5
docker pull grafana/loki:2.9.2
docker pull grafana/promtail:2.9.2
```
SVC
``` bash
# backend 目录下
# 下载依赖
bash load_lib.sh
# 启动 mysql、redis 等中间件
docker compose up -d
# 启动所有服务, 输出在 log/${svc}.log 中
bash start_service.sh
# 关闭所有服务
bash kill_service.sh
```

Docker Start
```bash
# backend 目录下
# build images
make build-all v=v1.0.0
# 启动 mysql、redis 等中间件
docker compose up -d
# start service
docker compose -f deploy/docker-compose-svc.yaml up -d
# clean service
docker compose -f deploy/docker-compose-svc.yaml down
```

# 目录结构

```bash
.
├── app             # 主要业务逻辑
│  ├── cart         # 购物车服务
│  ├── checkout     # 结算服务
│  ├── email        # 邮件服务
│  ├── frontend     # 前端的http服务
│  ├── order        # 订单服务
│  ├── payment      # 支付服务
│  ├── product      # 商品服务
│  └── user         # 用户服务
├── common          
│  ├── clientsuite  # 客户端配置
│  ├── mtl          # 指标
│  └── serversuite  # 服务端配置
├── rpc_gen         # 自动生成的 rpc client 代码，在 backend 目录下，通过 make gen-client svc=user 可单独生成
│  ├── kitex_gen
│  └── rpc
├── db              
│  └── sql/ini/databases.sql    # mysql 的初始化文件，需将该文件映射到 /docker-entrypoint-initdb.d/ 目录下
├── deploy                      # 部署相关文件
│  ├── config                   # 配置文件
│  ├── docker-compose-svc.yaml  # 整个服务的 docker-compose 文件（不包括 mysql 等中间件）
│  ├── Dockerfile.frontend      # frontend 的 Dockerfile
│  ├── Dockerfile.svc           # svc（除去 frontend） 的 Dockerfile
│  ├── gomall-dev-app.yaml      # k8s 的 svc 配置
│  ├── gomall-dev-base.yaml     # k8s 的 mysql 等中间件的配置
│  ├── gomall-dev-cluster.yaml  # kind 的 cluster 配置
│  └── grafana.json             # grafana 的 dashboard 配置文件
├── idl                 # 接口定义文件
│  ├── api.proto        # 通用 api 定义
│  ├── cart.proto       # 购物车服务的 api 定义
│  ├── checkout.proto   # 结算服务的 api 定义
│  ├── echo.proto       # 测试 proto
│  ├── echo.thrift      # 测试 thrift
│  ├── email.proto      # 邮件服务的 api 定义
│  ├── frontend         # 前端 http 服务的 api 定义目录
│  ├── order.proto      # 订单服务的 api 定义
│  ├── payment.proto    # 支付服务的 api 定义
│  ├── product.proto    # 商品服务的 api 定义
│  └── user.proto       # 用户服务的 api 定义
├── docker-compose.yaml # mysql、redis 等中间件的 docker-compose 文件
├── Dockerfile          # 制造 docker 镜像的 Dockerfile 模板
├── Makefile
├── kill_service.sh     # 关闭所有 go 服务的脚本
├── start_service.sh    # 启动所有服务的脚本
└── README.md
```

业务代码目录结构，以 user 服务为例：
```bash
.
├── biz                 # 业务逻辑，主要编写的部分
│  ├── dal              # 数据访问层
│  ├── model            # 数据模型
│  └── service          # 服务层
├── build.sh
├── conf                # 配置文件
│  ├── conf.go
│  ├── dev          
│  ├── online
│  └── test
├── docker-compose.yaml # 依赖的中间件, 不用
├── handler.go          # rpc 服务的 handler
├── main.go             # 入口文件
└── .env                # godotenv 加载的环境变量配置
```