#!/bin/bash
# 停止并清理容器
set -e

docker-compose down

echo "服务已停止并清理。"
