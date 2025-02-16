#!/bin/bash

# backend目录执行

ROOT_PATH=$PWD
declare -a MODULES=("common" "rpc_gen")
# 启动所有微服务
for module in "${MODULES[@]}"; do
    echo "Tidying $module"
    # 进入微服务目录
    pushd "$ROOT_PATH/$module" > /dev/null
    # 检查是否成功进入目录
    if [ $? -ne 0 ]; then
        echo "Failed to enter directory for $module. Skipping..."
        continue
    fi
    go mod tidy
    echo "$module tidied."
    echo "==================="
done

APP_PATH=$ROOT_PATH/app
# 声明一个数组，存储所有微服务的目录名称
declare -a SERVICES=("cart" "checkout" "email" "frontend" "order" "payment" "product" "user")
for service in "${SERVICES[@]}"; do
    echo "Tidying $service"
    # 进入微服务目录
    pushd "$APP_PATH/$service" > /dev/null
    # 检查是否成功进入目录
    if [ $? -ne 0 ]; then
        echo "Failed to enter directory for $service. Skipping..."
        continue
    fi
    go mod tidy
    echo "$service tidied."
    echo "==================="
done