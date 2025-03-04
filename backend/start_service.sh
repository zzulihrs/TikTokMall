#!/bin/bash
# backend目录执行
ROOT_PATH=$PWD/app
# 微服务列表，每个微服务的目录名称
declare -a SERVICES=("cart" "checkout" "email" "frontend" "order" "payment" "product" "user" "merchant")
sleep 10 # 等docker启动完成
# 创建数据文件夹（如果不存在）
mkdir -p ./data/mysql/ ./data/redis/

# 检查并修改 ./data/mysql/ 目录的权限
mysql_permissions=$(stat -c "%a" ./data/mysql/)
if [ "$mysql_permissions" != "777" ]; then
    sudo chmod -R 777 ./data/mysql/
fi

# 检查并修改 ./data/redis/ 目录的权限
redis_permissions=$(stat -c "%a" ./data/redis/)
if [ "$redis_permissions" != "777" ]; then
    sudo chmod -R 777 ./data/redis/
fi
# 启动所有微服务
for service in "${SERVICES[@]}"; do
    echo "Starting $service..."
    # 进入微服务目录
    pushd "$ROOT_PATH/$service" > /dev/null

    # 检查是否成功进入目录
    if [ $? -ne 0 ]; then
        echo "Failed to enter directory for $service. Skipping..."
        continue
    fi

    # 创建日志文件夹（如果不存在）
    LOG_DIR="$ROOT_PATH/$service/log"
    if [ ! -d "$LOG_DIR" ]; then
        mkdir -p "$LOG_DIR"
    fi

    # 日志文件路径
    LOG_FILE="$LOG_DIR/$service.log"

    go mod tidy
    # 使用 nohup 启动微服务，并将输出重定向到日志文件
    nohup air > "$LOG_FILE" 2>&1 &
    pid=$!

    # 输出对应的进程ID，可以选择kill某个微服务
    echo "Started $service with PID $pid"
    # 返回到脚本开始时的目录
    popd > /dev/null
    echo "$service started."

    # 等待 5 秒, 同时启动多个服务太占内存
    sleep 5
    echo "sleep 5s."
done

echo "All services have been started."