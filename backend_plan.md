# SmartCampus 后端开发计划 (Go + PostgreSQL)

## 📋 项目概述

SmartCampus 是一个基于 Go 语言和 PostgreSQL 的智慧校园系统后端，采用微服务架构，为教师、学生和管理员提供完整的教育管理解决方案。

## 🛠️ 技术栈选择

### 核心框架
```go
// Web 框架
- Gin (高性能 HTTP 框架)
- Gorm (ORM 框架)
- Viper (配置管理)
- Zap (结构化日志)

// 数据库
- PostgreSQL 15 (主数据库)
- Redis 7.0 (缓存和会话)

// 认证授权
- JWT-Go (令牌认证)
- Bcrypt (密码加密)
- Casbin (权限管理)

// 实时通信
- Gorilla WebSocket
- Redis Pub/Sub

// 消息队列
- Asynq (分布式任务队列)

// 文件处理
- 本地文件存储 + 云存储适配器
```

### 选择理由
- **Gin**: 高性能、低内存占用，适合高并发场景
- **PostgreSQL**: 强大的关系型数据库，支持复杂查询和事务
- **Redis**: 高速缓存，提升系统性能
- **Gorm**: 简化数据库操作，提高开发效率

## 🗄️ 数据库设计

### 核心表结构

#### 用户体系表
```sql
-- 用户表
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'teacher', 'student')),
    full_name VARCHAR(100) NOT NULL,
    avatar_url VARCHAR(255),
    phone VARCHAR(20),
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'inactive', 'suspended')),
    last_login_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 用户档案表
CREATE TABLE user_profiles (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    student_id VARCHAR(50) UNIQUE, -- 学生学号
    teacher_id VARCHAR(50) UNIQUE, -- 教师工号
    department VARCHAR(100),
    grade VARCHAR(50),
    birth_date DATE,
    address TEXT,
    emergency_contact JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### 学校组织表
```sql
-- 学校表
CREATE TABLE schools (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(200) NOT NULL,
    code VARCHAR(50) UNIQUE NOT NULL,
    address TEXT,
    phone VARCHAR(20),
    email VARCHAR(100),
    principal VARCHAR(100),
    settings JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 班级表
CREATE TABLE classes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    school_id UUID NOT NULL REFERENCES schools(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    grade VARCHAR(50) NOT NULL,
    classroom VARCHAR(50),
    capacity INTEGER DEFAULT 40,
    head_teacher_id UUID REFERENCES users(id),
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 班级学生关联表
CREATE TABLE class_students (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    class_id UUID NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
    student_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    seat_number INTEGER,
    enrollment_date DATE DEFAULT CURRENT_DATE,
    status VARCHAR(20) DEFAULT 'active',
    UNIQUE(class_id, student_id)
);
```

#### 课程教学表
```sql
-- 课程表
CREATE TABLE courses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(200) NOT NULL,
    code VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    credit INTEGER DEFAULT 1,
    teacher_id UUID NOT NULL REFERENCES users(id),
    class_id UUID NOT NULL REFERENCES classes(id) ON DELETE CASCADE,
    schedule JSONB NOT NULL, -- {weekday: 1, start_time: "08:00", end_time: "09:40"}
    room VARCHAR(50),
    status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 作业表
CREATE TABLE assignments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    assignment_type VARCHAR(50) DEFAULT 'homework',
    max_score DECIMAL(5,2) DEFAULT 100.00,
    due_date TIMESTAMP NOT NULL,
    attachments JSONB DEFAULT '[]',
    status VARCHAR(20) DEFAULT 'published',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 作业提交表
CREATE TABLE assignment_submissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    assignment_id UUID NOT NULL REFERENCES assignments(id) ON DELETE CASCADE,
    student_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content TEXT,
    attachments JSONB DEFAULT '[]',
    submitted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(20) DEFAULT 'submitted',
    UNIQUE(assignment_id, student_id)
);

-- 成绩表
CREATE TABLE grades (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    submission_id UUID NOT NULL REFERENCES assignment_submissions(id) ON DELETE CASCADE,
    score DECIMAL(5,2),
    feedback TEXT,
    graded_by UUID REFERENCES users(id),
    graded_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### 实时通信表
```sql
-- 消息表
CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sender_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    receiver_type VARCHAR(20) NOT NULL CHECK (receiver_type IN ('user', 'class', 'broadcast')),
    receiver_id UUID, -- 用户ID或班级ID
    content TEXT NOT NULL,
    message_type VARCHAR(20) DEFAULT 'text',
    attachments JSONB DEFAULT '[]',
    status VARCHAR(20) DEFAULT 'sent',
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    read_at TIMESTAMP
);

-- 通知表
CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(200) NOT NULL,
    content TEXT NOT NULL,
    notification_type VARCHAR(50) NOT NULL,
    target_audience VARCHAR(20) CHECK (target_audience IN ('all', 'teachers', 'students', 'class')),
    target_id UUID, -- 特定班级ID
    priority VARCHAR(20) DEFAULT 'normal',
    expires_at TIMESTAMP,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 🔐 权限系统设计

### 角色定义
```go
type Role string

const (
    RoleAdmin   Role = "admin"
    RoleTeacher Role = "teacher" 
    RoleStudent Role = "student"
)

// 权限定义
type Permission string

const (
    // 系统管理权限
    PermissionSystemManage Permission = "system:manage"
    PermissionUserManage   Permission = "user:manage"
    PermissionSchoolManage Permission = "school:manage"
    
    // 教学管理权限
    PermissionClassManage  Permission = "class:manage"
    PermissionCourseManage Permission = "course:manage"
    
    // 作业相关权限
    PermissionAssignmentCreate Permission = "assignment:create"
    PermissionAssignmentGrade  Permission = "assignment:grade"
    PermissionAssignmentSubmit Permission = "assignment:submit"
    
    // 成绩权限
    PermissionGradeView   Permission = "grade:view"
    PermissionGradeManage Permission = "grade:manage"
    
    // 消息权限
    PermissionMessageSend     Permission = "message:send"
    PermissionMessageBroadcast Permission = "message:broadcast"
)
```

### 角色权限矩阵
| 权限 | 管理员 | 教师 | 学生 |
|------|--------|------|------|
| system:manage | ✅ | ❌ | ❌ |
| user:manage | ✅ | ❌ | ❌ |
| school:manage | ✅ | ❌ | ❌ |
| class:manage | ✅ | ✅(自己班级) | ❌ |
| course:manage | ✅ | ✅(自己课程) | ❌ |
| assignment:create | ✅ | ✅ | ❌ |
| assignment:grade | ✅ | ✅(自己课程) | ❌ |
| assignment:submit | ✅ | ❌ | ✅ |
| grade:view | ✅ | ✅(自己课程) | ✅(自己成绩) |
| grade:manage | ✅ | ✅(自己课程) | ❌ |
| message:send | ✅ | ✅ | ✅ |
| message:broadcast | ✅ | ✅(自己班级) | ❌ |

### Casbin 权限配置
```ini
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)
```

## 🏗️ 系统架构

### 微服务划分
```
SmartCampus Backend
├── API Gateway
│   ├── 路由分发
│   ├── 认证中间件
│   ├── 限流控制
│   └── 日志记录
├── 用户服务 (User Service)
│   ├── 用户管理
│   ├── 认证授权
│   └── 权限验证
├── 教学服务 (Teaching Service)
│   ├── 班级管理
│   ├── 课程管理
│   └── 学生管理
├── 作业服务 (Assignment Service)
│   ├── 作业发布
│   ├── 作业提交
│   └── 作业批改
├── 成绩服务 (Grade Service)
│   ├── 成绩录入
│   ├── 成绩统计
│   └── 成绩分析
├── 消息服务 (Message Service)
│   ├── 实时聊天
│   ├── 通知推送
│   └── 广播消息
└── 文件服务 (File Service)
    ├── 文件上传
    ├── 文件存储
    └── 文件分发
```

## 📅 实施步骤

### 第一阶段：项目基础搭建 (第1-2周)

#### 1.1 项目初始化
```bash
# 创建项目结构
mkdir -p smartcampus/{cmd,internal,api,pkg,config,migrations,scripts}
cd smartcampus

# 初始化 Go 模块
go mod init github.com/your-org/smartcampus

# 安装核心依赖
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/spf13/viper
go get -u go.uber.org/zap
```

#### 1.2 配置管理系统
```go
// config/config.go
package config

type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    Redis    RedisConfig    `mapstructure:"redis"`
    JWT      JWTConfig      `mapstructure:"jwt"`
}

type ServerConfig struct {
    Port         string `mapstructure:"port"`
    Mode         string `mapstructure:"mode"`
    ReadTimeout  int    `mapstructure:"read_timeout"`
    WriteTimeout int    `mapstructure:"write_timeout"`
}

type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    User     string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    DBName   string `mapstructure:"dbname"`
    SSLMode  string `mapstructure:"sslmode"`
}

type RedisConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    Password string `mapstructure:"password"`
    DB       int    `mapstructure:"db"`
}

type JWTConfig struct {
    Secret string `mapstructure:"secret"`
    Expire int    `mapstructure:"expire"`
}
```

#### 1.3 数据库迁移
```bash
# 创建迁移脚本
mkdir -p migrations/001_initial

# 使用 golang-migrate
migrate create -ext sql -dir migrations/001_initial -seq create_initial_tables
```

### 第二阶段：核心功能开发 (第3-8周)

#### 2.1 用户认证系统 (第3周)
```go
// internal/auth/jwt.go
package auth

type JWTManager struct {
    secretKey     string
    tokenDuration time.Duration
}

func (manager *JWTManager) Generate(user *models.User) (string, error) {
    // JWT 令牌生成
}

func (manager *JWTManager) Verify(tokenString string) (*Claims, error) {
    // JWT 令牌验证
}

// internal/auth/middleware.go
func AuthMiddleware(jwtManager *auth.JWTManager) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 认证中间件
    }
}
```

#### 2.2 用户管理服务 (第4周)
```go
// internal/services/user_service.go
type UserService struct {
    db        *gorm.DB
    validator *validator.Validate
}

func (s *UserService) CreateUser(req *CreateUserRequest) (*UserResponse, error) {
    // 创建用户
}

func (s *UserService) GetUserByID(userID string) (*UserResponse, error) {
    // 获取用户信息
}

func (s *UserService) UpdateUser(userID string, req *UpdateUserRequest) error {
    // 更新用户信息
}
```

#### 2.3 教学管理服务 (第5-6周)
```go
// internal/services/teaching_service.go
type TeachingService struct {
    db *gorm.DB
}

func (s *TeachingService) CreateClass(req *CreateClassRequest) (*ClassResponse, error) {
    // 创建班级
}

func (s *TeachingService) AddStudentToClass(classID, studentID string) error {
    // 添加学生到班级
}

func (s *TeachingService) GetClassStudents(classID string) ([]*UserResponse, error) {
    // 获取班级学生列表
}
```

#### 2.4 作业服务 (第7周)
```go
// internal/services/assignment_service.go
type AssignmentService struct {
    db      *gorm.DB
    storage file.Storage
}

func (s *AssignmentService) CreateAssignment(req *CreateAssignmentRequest) (*AssignmentResponse, error) {
    // 创建作业
}

func (s *AssignmentService) SubmitAssignment(req *SubmitAssignmentRequest) error {
    // 提交作业
}

func (s *AssignmentService) GradeAssignment(req *GradeAssignmentRequest) error {
    // 批改作业
}
```

#### 2.5 实时消息服务 (第8周)
```go
// internal/services/message_service.go
type MessageService struct {
    db        *gorm.DB
    redis     *redis.Client
    hub       *ws.Hub
}

func (s *MessageService) SendMessage(req *SendMessageRequest) error {
    // 发送消息
}

func (s *MessageService) BroadcastToClass(classID string, message *Message) error {
    // 班级广播
}

// internal/ws/hub.go
type Hub struct {
    clients    map[*Client]bool
    broadcast  chan []byte
    register   chan *Client
    unregister chan *Client
}
```

### 第三阶段：API 接口开发 (第9-12周)

#### 3.1 RESTful API 设计
```go
// api/routes/routes.go
func SetupRouter(router *gin.Engine, deps *Dependencies) {
    // 公开路由
    public := router.Group("/api/v1")
    {
        public.POST("/auth/login", deps.AuthHandler.Login)
        public.POST("/auth/refresh", deps.AuthHandler.RefreshToken)
    }

    // 需要认证的路由
    auth := router.Group("/api/v1")
    auth.Use(middleware.AuthMiddleware(deps.JWTManager))
    {
        // 用户管理
        auth.GET("/users/profile", deps.UserHandler.GetProfile)
        auth.PUT("/users/profile", deps.UserHandler.UpdateProfile)
        
        // 班级管理
        auth.GET("/classes", deps.ClassHandler.ListClasses)
        auth.POST("/classes", deps.ClassHandler.CreateClass)
        auth.GET("/classes/:id/students", deps.ClassHandler.GetClassStudents)
        
        // 作业管理
        auth.GET("/assignments", deps.AssignmentHandler.ListAssignments)
        auth.POST("/assignments", deps.AssignmentHandler.CreateAssignment)
        auth.POST("/assignments/:id/submit", deps.AssignmentHandler.SubmitAssignment)
        
        // 消息系统
        auth.GET("/messages", deps.MessageHandler.ListMessages)
        auth.POST("/messages", deps.MessageHandler.SendMessage)
        auth.GET("/ws", deps.WSHandler.HandleWebSocket)
    }

    // 管理员路由
    admin := auth.Group("")
    admin.Use(middleware.RoleMiddleware("admin"))
    {
        admin.GET("/admin/users", deps.UserHandler.ListUsers)
        admin.POST("/admin/users", deps.UserHandler.CreateUser)
        admin.PUT("/admin/users/:id", deps.UserHandler.UpdateUser)
    }
}
```

#### 3.2 WebSocket 实时通信
```go
// internal/ws/client.go
type Client struct {
    hub     *Hub
    conn    *websocket.Conn
    send    chan []byte
    userID  string
    role    string
}

func (c *Client) readPump() {
    // 读取消息
}

func (c *Client) writePump() {
    // 发送消息
}
```

### 第四阶段：测试和优化 (第13-14周)

#### 4.1 单元测试
```go
// internal/services/user_service_test.go
func TestUserService_CreateUser(t *testing.T) {
    // 初始化测试数据库
    db := setupTestDB()
    defer cleanupTestDB(db)
    
    service := NewUserService(db)
    
    req := &CreateUserRequest{
        Username: "testuser",
        Email:    "test@example.com",
        Password: "password123",
        Role:     "student",
        FullName: "Test User",
    }
    
    user, err := service.CreateUser(req)
    assert.NoError(t, err)
    assert.Equal(t, "testuser", user.Username)
    assert.Equal(t, "student", user.Role)
}
```

#### 4.2 性能优化
```go
// 数据库查询优化
func (s *UserService) GetUserWithProfile(userID string) (*UserResponse, error) {
    var user User
    err := s.db.
        Preload("Profile").
        Preload("Classes").
        Where("id = ?", userID).
        First(&user).Error
        
    return user.ToResponse(), err
}

// Redis 缓存
func (s *UserService) GetUserByID(userID string) (*UserResponse, error) {
    cacheKey := fmt.Sprintf("user:%s", userID)
    
    // 尝试从缓存获取
    cached, err := s.redis.Get(cacheKey).Result()
    if err == nil {
        var user UserResponse
        json.Unmarshal([]byte(cached), &user)
        return &user, nil
    }
    
    // 从数据库获取并缓存
    user, err := s.getUserFromDB(userID)
    if err != nil {
        return nil, err
    }
    
    userJSON, _ := json.Marshal(user)
    s.redis.Set(cacheKey, userJSON, time.Hour)
    
    return user, nil
}
```

### 第五阶段：部署和监控 (第15-16周)

#### 5.1 Docker 部署
```dockerfile
# Dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/config ./config

EXPOSE 8080
CMD ["./main"]
```

#### 5.2 健康检查
```go
// api/handlers/health.go
func (h *HealthHandler) Check(c *gin.Context) {
    // 数据库健康检查
    db, _ := h.db.DB()
    if err := db.Ping(); err != nil {
        c.JSON(503, gin.H{"status": "error", "message": "Database unavailable"})
        return
    }
    
    // Redis 健康检查
    if _, err := h.redis.Ping().Result(); err != nil {
        c.JSON(503, gin.H{"status": "error", "message": "Redis unavailable"})
        return
    }
    
    c.JSON(200, gin.H{
        "status":    "ok",
        "timestamp": time.Now().Format(time.RFC3339),
        "version":   "1.0.0",
    })
}
```

## 📊 项目里程碑

| 阶段 | 时间 | 交付物 |
|------|------|--------|
| 项目搭建 | 第1-2周 | 项目结构、配置管理、数据库迁移 |
| 核心功能 | 第3-8周 | 用户系统、教学管理、作业系统、消息系统 |
| API 开发 | 第9-12周 | RESTful API、WebSocket、文档 |
| 测试优化 | 第13-14周 | 单元测试、性能优化、安全审计 |
| 部署上线 | 第15-16周 | Docker 部署、监控系统、文档 |

## 🔧 开发规范

### 代码规范
- 使用 `gofmt` 和 `goimports` 格式化代码
- 遵循 Go 语言官方代码规范
- 编写完整的单元测试，覆盖率 >80%
- 使用 `golint` 和 `staticcheck` 进行代码检查

### API 设计规范
- RESTful 风格 API 设计
- 统一的响应格式
- 适当的 HTTP 状态码
- 完整的 API 文档

### 数据库规范
- 使用迁移脚本管理数据库变更
- 为常用查询字段创建索引
- 避免 N+1 查询问题
- 定期进行数据库性能优化

这个开发计划提供了完整的后端开发路线图，从技术栈选择到具体的实施步骤，确保项目能够按时高质量交付。