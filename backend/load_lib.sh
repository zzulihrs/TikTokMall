#!/bin/bash

# backend目录执行
ROOT_PATH=$PWD
APP_PATH="$ROOT_PATH/app"

# 微服务列表，每个微服务的目录名称
declare -a SERVICES=("cart" "checkout" "email" "frontend" "order" "payment" "product" "user")

# 启动所有微服务
for service in "${SERVICES[@]}"; do
    echo "app/$service"
    echo "loading libs "
    # 进入微服务目录
    pushd "$APP_PATH/$service" > /dev/null
    go mod tidy
    echo "libs loaded"
    echo "========================="
done
echo "common"
echo "loading libs "
# 进入common目录
pushd "$ROOT_PATH/common" > /dev/null
go mod tidy
echo "libs loaded"
echo "========================="

echo "rpc_gen"
echo "loading libs "
# 进入common目录
pushd "$ROOT_PATH/rpc_gen" > /dev/null
go mod tidy
echo "libs loaded"
echo "========================="