#!/bin/bash
# 部署脚本：构建并启动容器
set -e

# 构建镜像
DOCKER_BUILDKIT=1 docker build -t smartcampus-backend .

# 启动容器（如有 docker-compose.yml 可用）
docker-compose up -d

echo "部署完成，服务已启动。"
