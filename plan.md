## 2. Workbench 统一Web应用设计方案

### 📋 应用概述

Workbench 是一个统一的智慧校园Web应用，根据用户角色（教师/学生）动态展示相应功能界面，提供完整的教学管理和学习体验。

### 🎨 统一设计语言

#### 视觉设计系统
```
色彩体系：
- 主色调：智慧蓝 (#1E40AF)
- 辅助色：活力橙 (#EA580C) - 教师强调色
- 辅助色：成长绿 (#16A34A) - 学生强调色
- 中性色：灰阶 palette (#F8FAFC → #0F172A)

字体系统：
- 英文：Inter, system-ui
- 中文：PingFang SC, HarmonyOS Sans SC
- 代码：JetBrains Mono

设计Token：
- 间距：4px基数 (4,8,12,16,20,24,32,40,48,64,80,96)
- 圆角：小(4px)、中(8px)、大(12px)
- 阴影：3级阴影系统
```

#### 角色差异化设计
```css
/* 教师主题 */
.teacher-theme {
  --primary: #1E40AF;
  --primary-foreground: #EFF6FF;
  --accent: #EA580C;
}

/* 学生主题 */  
.student-theme {
  --primary: #16A34A;
  --primary-foreground: #F0FDF4;
  --accent: #1E40AF;
}
```

### 🛠️ 技术栈设计

#### 前端技术栈
```typescript
// 核心框架
- React 18 + TypeScript 5.0
- Vite 5.0 (构建工具)
- PNPM (包管理器)

// UI框架
- Ant Design 5.x + 自定义主题
- Styled Components (CSS-in-JS)
- Framer Motion (动画)

// 状态管理
- Redux Toolkit + RTK Query
- React Hook Form (表单)
- React Router v6 (路由)

// 功能库
- Socket.IO Client 4.7 (实时通信)
- React PDF (文档预览)
- CodeMirror 6 (代码编辑器)
- Chart.js (数据可视化)

// 工具库
- Day.js (日期处理)
- Axios (HTTP客户端)
- Zod (数据验证)
```

#### 后端技术栈
```typescript
// 运行时
- Node.js 18+ LTS
- TypeScript 5.0

// 框架
- NestJS 10.0 (企业级框架)
- Express (底层HTTP)

// 数据库
- PostgreSQL 15 (主数据库)
- Redis 7.0 (缓存/会话)

// ORM与工具
- Prisma 5.0 (ORM)
- Class Validator (数据验证)
- JWT (认证)
- Bcrypt (加密)

// 文件处理
- Multer (文件上传)
- Sharp (图片处理)
```

### 🏗️ 应用架构

#### 前端架构
```
Workbench Web App
├── 核心层
│   ├── 认证模块 (RBAC权限)
│   ├── 路由守卫 (角色路由)
│   └── 主题管理 (动态换肤)
├── 通用组件层
│   ├── 布局组件
│   ├── 业务组件
│   └── UI组件
├── 功能模块层
│   ├── 仪表盘 (角色专属)
│   ├── 课程管理
│   ├── 作业系统
│   ├── 成绩管理
│   ├── 实时通信
│   └── 个人中心
└── 服务层
    ├── API服务
    ├── WebSocket服务
    └── 工具函数
```

#### 后端微服务架构
```
Workbench Backend
├── API网关
│   ├── 请求路由
│   ├── 身份验证
│   ├── 速率限制
│   └── 日志记录
├── 核心服务
│   ├── 用户服务 (认证、权限、资料)
│   ├── 课程服务 (课程、班级、课表)
│   ├── 作业服务 (作业、提交、批改)
│   ├── 成绩服务 (成绩、统计、分析)
│   ├── 消息服务 (聊天、通知、广播)
│   └── 文件服务 (上传、存储、管理)
└── 支撑服务
    ├── 数据库 (PostgreSQL)
    ├── 缓存 (Redis)
    ├── 消息队列 (Bull Queue)
    └── 实时通信 (Socket.IO)
```

### 🔧 核心功能模块

#### 通用功能模块
```
通用功能
├── 用户认证
│   ├── 登录/注册
│   ├── 权限验证
│   └── 会话管理
├── 个人中心
│   ├── 个人信息
│   ├── 消息通知
│   └── 偏好设置
└── 系统功能
    ├── 文件管理
    ├── 搜索功能
    └── 帮助中心
```

#### 教师专属功能
```
教师工作台
├── 教学仪表盘
│   ├── 课程概览
│   ├── 待办事项
│   └── 快速操作
├── 班级管理
│   ├── 学生管理
│   ├── 座位安排
│   └── 考勤统计
├── 教学工具
│   ├── 作业发布
│   ├── 在线批改
│   ├── 成绩录入
│   └── 课堂活动
└── 数据分析
    ├── 学情分析
    ├── 成绩统计
    └── 教学报告
```

#### 学生专属功能
```
学生学习台
├── 学习空间
│   ├── 我的课程
│   ├── 学习进度
│   └── 课程表
├── 作业中心
│   ├── 作业列表
│   ├── 作业提交
│   └── 成绩查询
├── 协作学习
│   ├── 班级讨论
│   ├── 学习小组
│   └── 资源分享
└── 个人成长
    ├── 学习档案
    ├── 成就系统
    └── 学习统计
```

### 🐳 Docker部署配置

#### 1. 前端Dockerfile
```dockerfile
# 前端 Dockerfile
FROM node:18-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制包管理文件
COPY package.json pnpm-lock.yaml* ./

# 安装 pnpm
RUN npm install -g pnpm

# 安装依赖
RUN pnpm install --frozen-lockfile

# 复制源代码
COPY . .

# 构建应用
RUN pnpm build

# 生产阶段
FROM nginx:alpine

# 复制构建产物
COPY --from=builder /app/dist /usr/share/nginx/html

# 复制 nginx 配置
COPY nginx.conf /etc/nginx/nginx.conf

# 复制启动脚本
COPY docker-entrypoint.sh /
RUN chmod +x /docker-entrypoint.sh

# 暴露端口
EXPOSE 80

# 启动 nginx
ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]
```

#### 2. 前端Nginx配置
```nginx
# nginx.conf
events {
    worker_connections 1024;
}

http {
    include /etc/nginx/mime.types;
    default_type application/octet-stream;

    # 日志格式
    log_format main '$remote_addr - $remote_user [$time_local] "$request" '
                    '$status $body_bytes_sent "$http_referer" '
                    '"$http_user_agent" "$http_x_forwarded_for"';

    access_log /var/log/nginx/access.log main;
    error_log /var/log/nginx/error.log warn;

    # Gzip压缩
    gzip on;
    gzip_vary on;
    gzip_min_length 1024;
    gzip_types
        text/plain
        text/css
        text/xml
        text/javascript
        application/javascript
        application/xml+rss
        application/json;

    server {
        listen 80;
        server_name localhost;
        root /usr/share/nginx/html;
        index index.html;

        # 静态资源缓存
        location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg|woff|woff2)$ {
            expires 1y;
            add_header Cache-Control "public, immutable";
        }

        # SPA路由支持
        location / {
            try_files $uri $uri/ /index.html;
        }

        # API代理
        location /api/ {
            proxy_pass http://backend:3001;
            proxy_http_version 1.1;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection 'upgrade';
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
            proxy_cache_bypass $http_upgrade;
        }

        # WebSocket代理
        location /socket.io/ {
            proxy_pass http://backend:3001;
            proxy_http_version 1.1;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection "upgrade";
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }
    }
}
```

#### 3. 前端启动脚本
```bash
#!/bin/bash
# docker-entrypoint.sh

# 替换环境变量
envsubst < /usr/share/nginx/html/env.template.js > /usr/share/nginx/html/env.js

# 启动 nginx
exec "$@"
```

#### 4. 后端Dockerfile
```dockerfile
# 后端 Dockerfile
FROM node:18-alpine AS builder

# 安装构建依赖
RUN apk add --no-cache \
    python3 \
    make \
    g++ \
    libc6-compat

WORKDIR /app

# 复制包文件
COPY package.json pnpm-lock.yaml* ./
RUN npm install -g pnpm && pnpm install --frozen-lockfile

# 复制源代码
COPY . .

# 构建应用
RUN pnpm build

# 生产阶段
FROM node:18-alpine AS production

# 安装运行时依赖
RUN apk add --no-cache \
    dumb-init \
    libc6-compat

WORKDIR /app

# 创建非root用户
RUN addgroup -g 1001 -S nodejs && \
    adduser -S nextjs -u 1001

# 复制构建产物和依赖
COPY --from=builder --chown=nextjs:nodejs /app/dist ./dist
COPY --from=builder --chown=nextjs:nodejs /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./

# 复制 Prisma 相关文件
COPY --from=builder /app/prisma ./prisma

# 生成 Prisma 客户端
RUN npx prisma generate

# 切换用户
USER nextjs

# 暴露端口
EXPOSE 3001

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD node dist/health-check.js

# 启动应用
CMD ["dumb-init", "node", "dist/main.js"]
```

#### 5. Docker Compose配置
```yaml
# docker-compose.yml
version: '3.8'

services:
  # 前端服务
  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    ports:
      - "80:80"
    environment:
      - API_BASE_URL=http://localhost:3001
    depends_on:
      - backend
    networks:
      - campus-network

  # 后端服务
  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    ports:
      - "3001:3001"
    environment:
      - NODE_ENV=production
      - DATABASE_URL=postgresql://postgres:password@postgres:5432/smartcampus
      - REDIS_URL=redis://redis:6379
      - JWT_SECRET=your-jwt-secret
      - UPLOAD_PATH=/app/uploads
    volumes:
      - uploads:/app/uploads
    depends_on:
      - postgres
      - redis
    networks:
      - campus-network

  # 数据库
  postgres:
    image: postgres:15-alpine
    environment:
      - POSTGRES_DB=smartcampus
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=password
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./init-db:/docker-entrypoint-initdb.d
    networks:
      - campus-network

  # Redis缓存
  redis:
    image: redis:7-alpine
    command: redis-server --appendonly yes
    volumes:
      - redis_data:/data
    networks:
      - campus-network

  # 反向代理 (可选)
  nginx-proxy:
    image: nginx:alpine
    ports:
      - "443:443"
      - "80:80"
    volumes:
      - ./nginx/conf.d:/etc/nginx/conf.d
      - ./ssl:/etc/nginx/ssl
    depends_on:
      - frontend
    networks:
      - campus-network

volumes:
  postgres_data:
  redis_data:
  uploads:

networks:
  campus-network:
    driver: bridge
```

#### 6. 环境配置模板
```javascript
// frontend/public/env.template.js
window.env = {
  API_BASE_URL: '${API_BASE_URL:-http://localhost:3001}',
  WS_BASE_URL: '${WS_BASE_URL:-ws://localhost:3001}',
  NODE_ENV: '${NODE_ENV:-production}',
  VERSION: '${VERSION:-1.0.0}'
};
```

### 🚀 部署说明

#### 构建和运行
```bash
# 构建所有服务
docker-compose build

# 启动服务
docker-compose up -d

# 查看日志
docker-compose logs -f

# 停止服务
docker-compose down
```

#### 环境变量配置
创建 `.env` 文件：
```env
# 数据库配置
DATABASE_URL=postgresql://postgres:password@postgres:5432/smartcampus
REDIS_URL=redis://redis:6379

# JWT配置
JWT_SECRET=your-super-secret-jwt-key
JWT_EXPIRES_IN=7d

# 文件上传
MAX_FILE_SIZE=50MB
UPLOAD_PATH=/app/uploads

# 应用配置
API_BASE_URL=http://localhost:3001
NODE_ENV=production
```