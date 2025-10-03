# SmartCampus 前端开发计划 (Windows Fluent UI)

## 📋 项目概述

SmartCampus 前端采用 React + Fluent UI 技术栈，提供现代化的 Windows 风格用户体验，支持教师、学生和管理员三种角色的统一工作台界面。

## 🛠️ 前端技术栈

### 核心框架
```json
{
  "框架": "React 18 + TypeScript 5.0",
  "构建工具": "Vite 5.0",
  "包管理器": "PNPM",
  "UI组件库": "@fluentui/react-components (v9)",
  "状态管理": "Redux Toolkit + RTK Query",
  "路由": "React Router v6",
  "HTTP客户端": "Axios",
  "表单处理": "React Hook Form + Zod"
}
```

### Fluent UI v9 组件库
```typescript
// 主要组件包
import {
  Button,
  Input,
  Select,
  Table,
  Card,
  Nav,
  Layout,
  Avatar,
  Badge,
  MessageBar,
  Spinner
} from "@fluentui/react-components";

// 图标库
import {
  SearchRegular,
  PersonRegular,
  SettingsRegular,
  MailRegular
} from "@fluentui/react-icons";
```

## 🎨 设计系统

### Fluent Design 设计原则
- **Light** - 轻量、快速、高效
- **Depth** - 层次感和深度
- **Motion** - 流畅的动画效果
- **Material** - 真实的材质感
- **Scale** - 可扩展的设计系统

### 主题配置
```typescript
// src/theme/theme.ts
import { createLightTheme, createDarkTheme } from '@fluentui/react-components';

export const lightTheme = createLightTheme({
  // 品牌色 - 智慧校园蓝
  colorBrandBackground: '#0078D4',
  colorBrandBackgroundHover: '#106EBE',
  colorBrandBackgroundPressed: '#005A9E',
  
  // 中性色
  colorNeutralForeground1: '#242424',
  colorNeutralForeground2: '#424242',
  colorNeutralBackground1: '#FFFFFF',
  
  // 成功、警告、错误色
  colorStatusSuccess: '#107C10',
  colorStatusWarning: '#D83B01',
  colorStatusDanger: '#D13438'
});

export const darkTheme = createDarkTheme({
  colorBrandBackground: '#0088FF',
  colorBrandBackgroundHover: '#0078D4',
  // ... 其他颜色配置
});
```

### 设计 Token
```typescript
// src/theme/tokens.ts
export const designTokens = {
  // 间距系统 (基于4px)
  spacing: {
    xs: '4px',
    s: '8px',
    m: '16px',
    l: '24px',
    xl: '32px',
    xxl: '48px'
  },
  
  // 圆角
  borderRadius: {
    small: '4px',
    medium: '8px',
    large: '12px'
  },
  
  // 阴影
  shadow: {
    low: '0 1px 2px rgba(0,0,0,0.12)',
    medium: '0 2px 8px rgba(0,0,0,0.16)',
    high: '0 4px 16px rgba(0,0,0,0.20)'
  },
  
  // 动画
  animation: {
    fast: '150ms ease',
    normal: '200ms ease',
    slow: '300ms ease'
  }
};
```

## 🏗️ 项目架构

### 目录结构
```
src/
├── components/           # 可复用组件
│   ├── ui/              # 基础UI组件
│   │   ├── Button/
│   │   ├── Card/
│   │   └── Table/
│   ├── layout/          # 布局组件
│   │   ├── AppLayout/
│   │   ├── SideNav/
│   │   └── TopBar/
│   └── business/        # 业务组件
│       ├── UserCard/
│       ├── AssignmentCard/
│       └── GradeChart/
├── pages/               # 页面组件
│   ├── auth/            # 认证页面
│   ├── dashboard/       # 仪表盘
│   ├── classes/         # 班级管理
│   ├── assignments/     # 作业管理
│   ├── grades/          # 成绩管理
│   └── messages/        # 消息中心
├── hooks/               # 自定义Hooks
│   ├── useAuth.ts
│   ├── useWebSocket.ts
│   └── useFileUpload.ts
├── services/            # API服务
│   ├── api/
│   ├── websocket/
│   └── storage/
├── store/               # 状态管理
│   ├── slices/
│   └── api/
├── utils/               # 工具函数
│   ├── formatters.ts
│   ├── validators.ts
│   └── constants.ts
├── types/               # TypeScript类型定义
└── assets/              # 静态资源
```

### 组件架构设计
```typescript
// 应用入口结构
<FluentProvider theme={currentTheme}>
  <ReduxProvider store={store}>
    <Router>
      <AppLayout>
        <SideNav />
        <MainContent>
          <Routes>
            <Route path="/" element={<Dashboard />} />
            <Route path="/classes" element={<ClassManagement />} />
            <Route path="/assignments" element={<AssignmentCenter />} />
            {/* ... 其他路由 */}
          </Routes>
        </MainContent>
      </AppLayout>
    </Router>
  </ReduxProvider>
</FluentProvider>
```

## 📱 页面设计

### 通用布局组件
```typescript
// src/components/layout/AppLayout.tsx
export const AppLayout: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  return (
    <div className={styles.appLayout}>
      <TopBar />
      <div className={styles.mainContainer}>
        <SideNav />
        <main className={styles.mainContent}>
          {children}
        </main>
      </div>
      <NotificationContainer />
    </div>
  );
};
```

### 1. 登录页面
```typescript
// src/pages/auth/LoginPage.tsx
export const LoginPage: React.FC = () => {
  return (
    <div className={styles.loginContainer}>
      <Card className={styles.loginCard}>
        <div className={styles.logoSection}>
          <SchoolIcon fontSize={48} />
          <Title>SmartCampus</Title>
          <Text>智慧校园统一工作台</Text>
        </div>
        
        <LoginForm />
        
        <div className={styles.footer}>
          <Text size={200}>© 2024 SmartCampus. All rights reserved.</Text>
        </div>
      </Card>
    </div>
  );
};
```

### 2. 仪表盘页面 (角色差异化)
```typescript
// src/pages/dashboard/DashboardPage.tsx
export const DashboardPage: React.FC = () => {
  const { user } = useAuth();
  
  return (
    <div className={styles.dashboard}>
      <PageHeader title="仪表盘" description="欢迎回来" />
      
      <div className={styles.dashboardGrid}>
        {/* 通用组件 */}
        <QuickActions />
        <RecentActivities />
        <SystemNotifications />
        
        {/* 角色特定组件 */}
        {user.role === 'teacher' && <TeacherDashboard />}
        {user.role === 'student' && <StudentDashboard />}
        {user.role === 'admin' && <AdminDashboard />}
      </div>
    </div>
  );
};
```

### 3. 班级管理页面
```typescript
// src/pages/classes/ClassManagementPage.tsx
export const ClassManagementPage: React.FC = () => {
  return (
    <div className={styles.classManagement}>
      <PageHeader 
        title="班级管理" 
        description="管理班级和学生信息"
        actions={<CreateClassButton />}
      />
      
      <div className={styles.content}>
        <ClassList />
        <ClassDetails />
      </div>
    </div>
  );
};
```

## 🔄 状态管理设计

### Redux Store 结构
```typescript
// src/store/store.ts
export const store = configureStore({
  reducer: {
    auth: authSlice,
    user: userSlice,
    classes: classesSlice,
    assignments: assignmentsSlice,
    grades: gradesSlice,
    messages: messagesSlice,
    // RTK Query APIs
    api: apiReducer
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(apiMiddleware)
});

// API 服务定义
export const api = createApi({
  reducerPath: 'api',
  baseQuery: fetchBaseQuery({
    baseUrl: '/api/v1',
    prepareHeaders: (headers) => {
      const token = localStorage.getItem('token');
      if (token) {
        headers.set('authorization', `Bearer ${token}`);
      }
      return headers;
    }
  }),
  tagTypes: ['User', 'Class', 'Assignment', 'Grade', 'Message'],
  endpoints: (builder) => ({
    // 用户相关
    login: builder.mutation<LoginResponse, LoginRequest>({
      query: (credentials) => ({
        url: '/auth/login',
        method: 'POST',
        body: credentials
      })
    }),
    
    // 班级相关
    getClasses: builder.query<Class[], void>({
      query: () => '/classes',
      providesTags: ['Class']
    }),
    
    // 作业相关
    createAssignment: builder.mutation<Assignment, CreateAssignmentRequest>({
      query: (assignment) => ({
        url: '/assignments',
        method: 'POST',
        body: assignment
      }),
      invalidatesTags: ['Assignment']
    })
  })
});
```

## 🌐 API 集成服务

### HTTP 服务层
```typescript
// src/services/api/apiClient.ts
class ApiClient {
  private client: AxiosInstance;
  
  constructor() {
    this.client = axios.create({
      baseURL: import.meta.env.VITE_API_BASE_URL,
      timeout: 10000
    });
    
    this.setupInterceptors();
  }
  
  private setupInterceptors() {
    // 请求拦截器
    this.client.interceptors.request.use(
      (config) => {
        const token = localStorage.getItem('token');
        if (token) {
          config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
      },
      (error) => Promise.reject(error)
    );
    
    // 响应拦截器
    this.client.interceptors.response.use(
      (response) => response,
      (error) => {
        if (error.response?.status === 401) {
          this.handleUnauthorized();
        }
        return Promise.reject(error);
      }
    );
  }
  
  // API 方法
  async get<T>(url: string, params?: any): Promise<T> {
    const response = await this.client.get(url, { params });
    return response.data;
  }
  
  async post<T>(url: string, data?: any): Promise<T> {
    const response = await this.client.post(url, data);
    return response.data;
  }
  
  // 文件上传
  async uploadFile(url: string, file: File, onProgress?: (progress: number) => void) {
    const formData = new FormData();
    formData.append('file', file);
    
    const response = await this.client.post(url, formData, {
      headers: { 'Content-Type': 'multipart/form-data' },
      onUploadProgress: (progressEvent) => {
        if (onProgress && progressEvent.total) {
          const progress = (progressEvent.loaded / progressEvent.total) * 100;
          onProgress(Math.round(progress));
        }
      }
    });
    
    return response.data;
  }
}
```

### WebSocket 服务
```typescript
// src/services/websocket/websocketService.ts
class WebSocketService {
  private socket: WebSocket | null = null;
  private reconnectAttempts = 0;
  private maxReconnectAttempts = 5;
  
  connect() {
    const token = localStorage.getItem('token');
    this.socket = new WebSocket(`${import.meta.env.VITE_WS_BASE_URL}?token=${token}`);
    
    this.socket.onopen = () => {
      console.log('WebSocket connected');
      this.reconnectAttempts = 0;
    };
    
    this.socket.onmessage = (event) => {
      this.handleMessage(JSON.parse(event.data));
    };
    
    this.socket.onclose = () => {
      this.handleReconnect();
    };
  }
  
  private handleMessage(message: WebSocketMessage) {
    switch (message.type) {
      case 'NEW_MESSAGE':
        store.dispatch(addMessage(message.data));
        break;
      case 'ASSIGNMENT_GRADED':
        store.dispatch(updateAssignmentGrade(message.data));
        break;
      case 'SYSTEM_NOTIFICATION':
        store.dispatch(showNotification(message.data));
        break;
    }
  }
  
  sendMessage(type: string, data: any) {
    if (this.socket?.readyState === WebSocket.OPEN) {
      this.socket.send(JSON.stringify({ type, data }));
    }
  }
}
```

## 🎯 角色特定功能

### 教师工作台组件
```typescript
// src/components/business/TeacherDashboard.tsx
export const TeacherDashboard: React.FC = () => {
  const { data: classes } = useGetClassesQuery();
  const { data: pendingAssignments } = useGetPendingAssignmentsQuery();
  
  return (
    <>
      <Card className={styles.teacherCard}>
        <Title>教学概览</Title>
        <div className={styles.statsGrid}>
          <StatCard 
            title="负责班级" 
            value={classes?.length || 0} 
            icon={<ClassIcon />}
          />
          <StatCard 
            title="待批改作业" 
            value={pendingAssignments?.length || 0} 
            icon={<AssignmentIcon />}
          />
          <StatCard 
            title="未读消息" 
            value={12} 
            icon={<MailIcon />}
          />
        </div>
      </Card>
      
      <QuickGradePanel />
      <ClassPerformanceChart />
    </>
  );
};
```

### 学生工作台组件
```typescript
// src/components/business/StudentDashboard.tsx
export const StudentDashboard: React.FC = () => {
  const { data: upcomingAssignments } = useGetUpcomingAssignmentsQuery();
  const { data: recentGrades } = useGetRecentGradesQuery();
  
  return (
    <>
      <Card className={styles.studentCard}>
        <Title>学习概览</Title>
        <div className={styles.statsGrid}>
          <StatCard 
            title="待完成作业" 
            value={upcomingAssignments?.length || 0} 
            icon={<AssignmentIcon />}
          />
          <StatCard 
            title="平均成绩" 
            value="85.5" 
            icon={<GradeIcon />}
          />
          <StatCard 
            title="出勤率" 
            value="96%" 
            icon={<AttendanceIcon />}
          />
        </div>
      </Card>
      
      <AssignmentDeadlines />
      <GradeTrendChart />
    </>
  );
};
```

## 📦 构建和部署配置

### Vite 配置
```typescript
// vite.config.ts
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  server: {
    port: 3000,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true
      },
      '/ws': {
        target: 'http://localhost:8080',
        ws: true
      }
    }
  },
  build: {
    outDir: 'dist',
    sourcemap: true,
    rollupOptions: {
      output: {
        manualChunks: {
          vendor: ['react', 'react-dom'],
          fluentui: ['@fluentui/react-components'],
          utils: ['axios', 'date-fns', 'lodash']
        }
      }
    }
  },
  css: {
    modules: {
      localsConvention: 'camelCase'
    }
  }
});
```

### 环境配置
```typescript
// src/config/environment.ts
export const environment = {
  // 开发环境
  development: {
    apiBaseUrl: 'http://localhost:8080/api/v1',
    wsBaseUrl: 'ws://localhost:8080/ws',
    enableDebug: true
  },
  
  // 生产环境
  production: {
    apiBaseUrl: '/api/v1',
    wsBaseUrl: `ws://${window.location.host}/ws`,
    enableDebug: false
  }
};

export const getEnvironment = () => {
  return import.meta.env.MODE === 'production' 
    ? environment.production 
    : environment.development;
};
```

## 🚀 实施计划

### 第一阶段：项目搭建 (第1-2周)
```bash
# 1. 创建项目
pnpm create vite smartcampus-frontend --template react-ts
cd smartcampus-frontend

# 2. 安装依赖
pnpm add @fluentui/react-components @fluentui/react-icons
pnpm add @reduxjs/toolkit react-redux
pnpm add react-router-dom axios
pnpm add react-hook-form @hookform/resolvers zod

# 3. 配置代码规范
pnpm add -D eslint @typescript-eslint/eslint-plugin prettier
```

### 第二阶段：基础架构 (第3-4周)
- [ ] 配置 Fluent UI 主题系统
- [ ] 实现应用布局组件
- [ ] 设置路由和导航
- [ ] 配置状态管理 (Redux Toolkit)
- [ ] 实现认证流程

### 第三阶段：核心页面开发 (第5-8周)
- [ ] 仪表盘页面 (各角色)
- [ ] 班级管理页面
- [ ] 作业管理页面
- [ ] 成绩管理页面
- [ ] 消息中心页面

### 第四阶段：高级功能 (第9-10周)
- [ ] 实时通信 (WebSocket)
- [ ] 文件上传和管理
- [ ] 数据可视化图表
- [ ] 离线支持
- [ ] 响应式设计优化

### 第五阶段：测试和优化 (第11-12周)
- [ ] 单元测试和集成测试
- [ ] 性能优化
- [ ] 无障碍访问 (a11y)
- [ ] 打包和部署优化

## 🎨 Fluent UI 最佳实践

### 组件使用规范
```typescript
// 正确的 Fluent UI 组件使用
export const UserProfile: React.FC = () => {
  return (
    <Card>
      <CardHeader
        image={<Avatar name="张三" />}
        header={<Text weight="semibold">张三</Text>}
        description={<Text>三年级二班</Text>}
      />
      <CardFooter>
        <Button appearance="primary" icon={<EditRegular />}>
          编辑资料
        </Button>
      </CardFooter>
    </Card>
  );
};
```

### 自定义样式
```typescript
// 使用 makeStyles 进行样式封装
import { makeStyles } from '@fluentui/react-components';

const useStyles = makeStyles({
  dashboardGrid: {
    display: 'grid',
    gridTemplateColumns: 'repeat(auto-fit, minmax(300px, 1fr))',
    gap: '16px',
    padding: '16px'
  },
  
  statCard: {
    transition: 'transform 0.2s ease',
    ':hover': {
      transform: 'translateY(-2px)'
    }
  }
});
```

这个前端开发计划提供了完整的 Fluent UI 集成方案，确保应用具有现代化的 Windows 风格用户体验，同时保持高性能和良好的可维护性。