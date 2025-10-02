# Workbench 智慧校园系统 - Go + PostgreSQL + Redis 架构设计

## 🚀 技术栈架构

### 后端技术栈
```yaml
编程语言: Go 1.21+
Web框架: 
  - Gin (HTTP API)
  - GORM (ORM)
  - go-redis (Redis客户端)
数据库:
  - PostgreSQL 15+ (主数据存储)
  - Redis 7.0+ (缓存/会话/消息队列)
消息队列: 
  - Redis Streams (轻量级消息队列)
  - 或 NSQ (可选，用于高吞吐场景)
实时通信:
  - WebSocket (gorilla/websocket)
  - Server-Sent Events (SSE)
工具库:
  - validator (数据验证)
  - jwt-go (认证)
  - bcrypt (加密)
  - zap (日志)
  - viper (配置管理)
```

### 前端技术栈 (保持不变，与plan.md一致)
```typescript
React 18 + TypeScript + Vite + Ant Design
```

## 🏗️ 系统架构设计

### 后端微服务架构
```
Workbench Backend (Go)
├── API网关层
│   ├── 请求路由 & 负载均衡
│   ├── JWT认证中间件
│   ├── 速率限制 (Redis)
│   ├── 请求日志 & 审计
│   └── CORS处理
├── 业务服务层
│   ├── 用户服务 (auth.go, user.go)
│   ├── 课程服务 (course.go, class.go)
│   ├── 作业服务 (assignment.go, submission.go)
│   ├── 成绩服务 (grade.go, analysis.go)
│   ├── 消息服务 (message.go, notification.go)
│   └── 文件服务 (upload.go, file.go)
├── 数据访问层
│   ├── PostgreSQL (GORM)
│   ├── Redis (缓存 & 会话)
│   └── 数据库连接池
└── 支撑服务层
    ├── 定时任务 (cron)
    ├── 实时通信 (WebSocket)
    ├── 文件存储 (本地/MinIO)
    └── 监控告警 (Prometheus)
```

## 🗄️ 数据库设计

### PostgreSQL 核心表结构

```sql
-- 用户表
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('teacher', 'student', 'admin')),
    full_name VARCHAR(100) NOT NULL,
    avatar_url VARCHAR(255),
    is_active BOOLEAN DEFAULT true,
    last_login_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- 课程表
CREATE TABLE courses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(200) NOT NULL,
    description TEXT,
    code VARCHAR(50) UNIQUE NOT NULL,
    teacher_id UUID NOT NULL REFERENCES users(id),
    semester VARCHAR(50),
    credits INTEGER DEFAULT 0,
    is_published BOOLEAN DEFAULT false,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- 课程-学生关联表
CREATE TABLE course_students (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL REFERENCES courses(id),
    student_id UUID NOT NULL REFERENCES users(id),
    enrolled_at TIMESTAMPTZ DEFAULT NOW(),
    status VARCHAR(20) DEFAULT 'active',
    UNIQUE(course_id, student_id)
);

-- 作业表
CREATE TABLE assignments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL REFERENCES courses(id),
    title VARCHAR(200) NOT NULL,
    description TEXT,
    due_date TIMESTAMPTZ NOT NULL,
    max_score INTEGER DEFAULT 100,
    assignment_type VARCHAR(50) DEFAULT 'homework',
    attachments JSONB,
    created_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- 作业提交表
CREATE TABLE submissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    assignment_id UUID NOT NULL REFERENCES assignments(id),
    student_id UUID NOT NULL REFERENCES users(id),
    content TEXT,
    attachments JSONB,
    submitted_at TIMESTAMPTZ DEFAULT NOW(),
    status VARCHAR(20) DEFAULT 'submitted',
    score INTEGER,
    feedback TEXT,
    graded_by UUID REFERENCES users(id),
    graded_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE(assignment_id, student_id)
);

-- 消息表
CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sender_id UUID NOT NULL REFERENCES users(id),
    receiver_id UUID REFERENCES users(id),
    course_id UUID REFERENCES courses(id),
    title VARCHAR(200),
    content TEXT NOT NULL,
    message_type VARCHAR(50) DEFAULT 'notification',
    is_read BOOLEAN DEFAULT false,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- 创建索引优化查询性能
CREATE INDEX idx_courses_teacher_id ON courses(teacher_id);
CREATE INDEX idx_course_students_course_id ON course_students(course_id);
CREATE INDEX idx_course_students_student_id ON course_students(student_id);
CREATE INDEX idx_assignments_course_id ON assignments(course_id);
CREATE INDEX idx_submissions_assignment_id ON submissions(assignment_id);
CREATE INDEX idx_submissions_student_id ON submissions(student_id);
CREATE INDEX idx_messages_receiver_id ON messages(receiver_id);
CREATE INDEX idx_messages_course_id ON messages(course_id);
CREATE INDEX idx_messages_created_at ON messages(created_at DESC);
```

### Redis 数据结构设计

```go
// Redis Key 设计模式
type RedisKeys struct {
    // 用户会话
    UserSession(userID string) string { return fmt.Sprintf("session:%s", userID) }
    
    // 课程缓存
    CourseCache(courseID string) string { return fmt.Sprintf("course:%s", courseID) }
    
    // 作业缓存
    AssignmentCache(assignmentID string) string { return fmt.Sprintf("assignment:%s", assignmentID) }
    
    // 限流器
    RateLimit(key string) string { return fmt.Sprintf("ratelimit:%s", key) }
    
    // 在线用户
    OnlineUsers() string { return "online:users" }
    
    // 消息队列
    MessageQueue() string { return "queue:messages" }
}
```

## 🔧 Go 项目结构

```
workbench-backend/
├── cmd/
│   └── api/
│       └── main.go                 # 应用入口
├── internal/
│   ├── config/
│   │   └── config.go              # 配置管理
│   ├── database/
│   │   ├── postgres.go            # PostgreSQL连接
│   │   └── redis.go               # Redis连接
│   ├── models/
│   │   ├── user.go
│   │   ├── course.go
│   │   ├── assignment.go
│   │   └── ...                    # 其他数据模型
│   ├── handlers/
│   │   ├── auth.go
│   │   ├── user.go
│   │   ├── course.go
│   │   └── ...                    # HTTP处理器
│   ├── services/
│   │   ├── auth_service.go
│   │   ├── course_service.go
│   │   ├── assignment_service.go
│   │   └── ...                    # 业务逻辑层
│   ├── repositories/
│   │   ├── user_repo.go
│   │   ├── course_repo.go
│   │   ├── assignment_repo.go
│   │   └── ...                    # 数据访问层
│   ├── middleware/
│   │   ├── auth.go
│   │   ├── cors.go
│   │   ├── logger.go
│   │   └── ratelimit.go           # 中间件
│   ├── utils/
│   │   ├── jwt.go
│   │   ├── password.go
│   │   ├── validator.go
│   │   └── ...                    # 工具函数
│   └── websocket/
│       ├── hub.go
│       ├── client.go
│       └── handler.go             # WebSocket服务
├── pkg/
│   ├── response/
│   │   └── response.go            # 统一响应格式
│   └── cache/
│       └── cache.go               # 缓存工具
├── deployments/
│   ├── Dockerfile
│   ├── docker-compose.yml
│   └── nginx.conf
├── scripts/
│   ├── migrate.go                 # 数据库迁移
│   └── seed.go                    # 数据填充
├── go.mod
└── go.sum
```

## 🎯 核心API设计

### 认证模块
```go
// auth.go
type AuthHandler struct {
    userService services.UserService
    jwtUtil     utils.JWTUtil
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.BadRequest(c, "无效的请求参数")
        return
    }
    
    user, err := h.userService.Authenticate(req.Username, req.Password)
    if err != nil {
        response.Unauthorized(c, "用户名或密码错误")
        return
    }
    
    token, err := h.jwtUtil.GenerateToken(user.ID, user.Role)
    if err != nil {
        response.ServerError(c, "生成token失败")
        return
    }
    
    response.Success(c, LoginResponse{
        Token: token,
        User:  user.ToDTO(),
    })
}
```

### 课程服务
```go
// course_service.go
type CourseService struct {
    courseRepo repositories.CourseRepository
    cache      cache.Cache
}

func (s *CourseService) GetTeacherCourses(teacherID string, page, pageSize int) (*PaginationResponse, error) {
    cacheKey := fmt.Sprintf("teacher_courses:%s:%d:%d", teacherID, page, pageSize)
    
    // 尝试从缓存获取
    var result PaginationResponse
    if err := s.cache.Get(cacheKey, &result); err == nil {
        return &result, nil
    }
    
    // 缓存未命中，查询数据库
    courses, total, err := s.courseRepo.FindByTeacherID(teacherID, page, pageSize)
    if err != nil {
        return nil, err
    }
    
    result = PaginationResponse{
        Data:  courses,
        Total: total,
        Page:  page,
    }
    
    // 写入缓存，过期时间5分钟
    s.cache.Set(cacheKey, result, 5*time.Minute)
    
    return &result, nil
}
```

## 🐳 容器化部署

### 后端 Dockerfile
```dockerfile
# 多阶段构建
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 安装依赖
RUN apk add --no-cache git ca-certificates tzdata

# 复制go模块文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# 生产镜像
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/

# 复制二进制文件
COPY --from=builder /app/main .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# 创建非root用户
RUN addgroup -g 1001 -S app && \
    adduser -u 1001 -S app -G app

# 切换用户
USER app

# 暴露端口
EXPOSE 8080

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

CMD ["./main"]
```

### 更新后的 docker-compose.yml
```yaml
version: '3.8'

services:
  # Go后端服务
  backend:
    build:
      context: ./workbench-backend
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    environment:
      - APP_ENV=production
      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_NAME=smartcampus
      - DB_USER=postgres
      - DB_PASSWORD=password
      - REDIS_URL=redis:6379
      - JWT_SECRET=your-super-secret-jwt-key
    depends_on:
      - postgres
      - redis
    networks:
      - campus-network
    healthcheck:
      test: ["CMD", "wget", "--no-verbose", "--tries=1", "--spider", "http://localhost:8080/health"]
      interval: 30s
      timeout: 10s
      retries: 3

  # 前端服务 (保持不变)
  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    ports:
      - "80:80"
    environment:
      - API_BASE_URL=http://backend:8080
    depends_on:
      - backend
    networks:
      - campus-network

  # PostgreSQL数据库
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
    command: >
      postgres 
      -c shared_preload_libraries=pg_stat_statements 
      -c pg_stat_statements.track=all 
      -c max_connections=200 
      -c shared_buffers=256MB 
      -c effective_cache_size=1GB

  # Redis缓存
  redis:
    image: redis:7-alpine
    command: redis-server --appendonly yes --maxmemory 512mb --maxmemory-policy allkeys-lru
    volumes:
      - redis_data:/data
    networks:
      - campus-network

volumes:
  postgres_data:
  redis_data:

networks:
  campus-network:
    driver: bridge
```

## ⚡ 性能优化配置

### PostgreSQL 配置优化
```sql
-- 高并发配置
ALTER SYSTEM SET max_connections = 200;
ALTER SYSTEM SET shared_buffers = '256MB';
ALTER SYSTEM SET effective_cache_size = '1GB';
ALTER SYSTEM SET work_mem = '4MB';
ALTER SYSTEM SET maintenance_work_mem = '64MB';
ALTER SYSTEM SET random_page_cost = 1.1;
ALTER SYSTEM SET effective_io_concurrency = 200;
```

### Go 服务配置
```go
// config/config.go
type Config struct {
    Server struct {
        Port         string        `mapstructure:"port"`
        ReadTimeout  time.Duration `mapstructure:"read_timeout"`
        WriteTimeout time.Duration `mapstructure:"write_timeout"`
        IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
    } `mapstructure:"server"`
    
    Database struct {
        MaxOpenConns int `mapstructure:"max_open_conns"`
        MaxIdleConns int `mapstructure:"max_idle_conns"`
        ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
    } `mapstructure:"database"`
    
    Redis struct {
        PoolSize int `mapstructure:"pool_size"`
        MinIdleConns int `mapstructure:"min_idle_conns"`
    } `mapstructure:"redis"`
}

// 推荐的配置值
func DefaultConfig() *Config {
    return &Config{
        Server: struct{
            Port string
            ReadTimeout time.Duration
            WriteTimeout time.Duration  
            IdleTimeout time.Duration
        }{
            Port: "8080",
            ReadTimeout: 15 * time.Second,
            WriteTimeout: 15 * time.Second,
            IdleTimeout: 60 * time.Second,
        },
        Database: struct{
            MaxOpenConns int
            MaxIdleConns int
            ConnMaxLifetime time.Duration
        }{
            MaxOpenConns: 25,
            MaxIdleConns: 25,
            ConnMaxLifetime: 5 * time.Minute,
        },
        Redis: struct{
            PoolSize int
            MinIdleConns int
        }{
            PoolSize: 10,
            MinIdleConns: 5,
        },
    }
}
```

## 🔒 安全设计

### JWT 认证流程
```go
// middleware/auth.go
func AuthMiddleware(jwtUtil utils.JWTUtil) gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            response.Unauthorized(c, "缺少认证token")
            c.Abort()
            return
        }
        
        claims, err := jwtUtil.ValidateToken(tokenString)
        if err != nil {
            response.Unauthorized(c, "无效的token")
            c.Abort()
            return
        }
        
        // 将用户信息存入上下文
        c.Set("userID", claims.UserID)
        c.Set("userRole", claims.Role)
        c.Next()
    }
}

// 基于角色的访问控制
func RBAC(allowedRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole, exists := c.Get("userRole")
        if !exists {
            response.Unauthorized(c, "未认证用户")
            c.Abort()
            return
        }
        
        for _, role := range allowedRoles {
            if userRole == role {
                c.Next()
                return
            }
        }
        
        response.Forbidden(c, "权限不足")
        c.Abort()
    }
}
```

这个新的架构设计充分利用了 Go 语言的高并发特性和 PostgreSQL + Redis 的高性能组合，能够很好地支撑智慧校园系统的高并发访问和复杂的业务逻辑需求。