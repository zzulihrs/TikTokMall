#!/bin/bash

# 检查 log 文件夹是否存在
if [ ! -d "log" ]; then
    # 如果不存在，则创建 log 文件夹
    mkdir log
fi
make run-server svc=frontend 2>&1 > log/frontend.log &
make run-server svc=cart 2>&1 > log/cart.log &
make run-server svc=checkout 2>&1 > log/checkout.log &
make run-server svc=email 2>&1 > log/email.log &
make run-server svc=order 2>&1 > log/order.log &
make run-server svc=payment 2>&1 > log/payment.log &
make run-server svc=product 2>&1 > log/product.log &
make run-server svc=user 2>&1 > log/user.log &