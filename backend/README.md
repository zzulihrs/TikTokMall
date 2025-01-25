# TikTokMall

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
go work sync
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

Kubernetes