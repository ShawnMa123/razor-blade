# Razor Blade Usage Tracker / 剃须刀使用记录器

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Go Version](https://img.shields.io/badge/Go-1.23+-brightgreen)](https://golang.org/)
[![Vue Version](https://img.shields.io/badge/Vue-3.3+-4FC08D)](https://vuejs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.2+-3178C6)](https://www.typescriptlang.org/)

## English | [中文](#中文文档)

A comprehensive web application for tracking razor and blade usage records, managing inventory, and analyzing usage patterns.

### 🌟 Features

- **📊 Usage Tracking**: Record detailed shaving sessions with ratings and experiences
- **📦 Inventory Management**: Automatic blade count reduction when changed
- **📅 Calendar View**: Visual calendar displaying usage history with ratings
- **📈 Statistics Dashboard**: Comprehensive analytics and usage insights
- **🔄 Quick Recording**: Fast and intuitive usage record creation
- **💾 Data Storage**: Support for both SQLite database and in-memory storage
- **🎨 Modern UI**: Clean, responsive interface built with Element Plus
- **📱 Mobile Friendly**: Optimized for desktop and mobile devices

### 🛠️ Technology Stack

#### Backend
- **Go 1.23+** - High-performance backend server
- **Gin** - HTTP web framework
- **GORM** - ORM for database operations
- **SQLite** - Lightweight database (with in-memory fallback)
- **Viper** - Configuration management
- **Logrus** - Structured logging

#### Frontend
- **Vue 3** - Progressive JavaScript framework
- **TypeScript** - Type-safe development
- **Element Plus** - Vue 3 UI component library
- **Pinia** - State management
- **Vue Router** - Client-side routing
- **Axios** - HTTP client
- **ECharts** - Data visualization
- **Day.js** - Date manipulation

### 🚀 Quick Start

#### Prerequisites
- **Go 1.23+**
- **Node.js 18+**
- **npm** or **yarn**

#### Installation

1. **Clone the repository**
```bash
git clone https://github.com/ShawnMa123/razor-blade.git
cd razor-blade
```

2. **Start the backend server**
```bash
cd backend
go mod tidy
go run cmd/server/main.go
```

3. **Start the frontend development server**
```bash
cd frontend
npm install
npm run dev
```

4. **Access the application**
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080

#### Docker Deployment

```bash
# Build and run with Docker Compose
docker-compose up --build -d

# Access the application
# Frontend: http://localhost:3000
# Backend API: http://localhost:8080
```

### 📖 Usage Guide

#### Adding Razors and Blades
1. Navigate to the management section
2. Add your razor information (brand, model, etc.)
3. Add blade types with inventory quantities

#### Recording Usage
1. Use the quick record form on the dashboard
2. Select your razor and blade combination
3. Add rating and experience notes
4. Mark if you changed the blade (auto-reduces inventory)

#### Viewing Statistics
- Check the dashboard for usage overview
- View the calendar for historical records
- Analyze patterns and preferences

### 🏗️ Architecture

```
razor-blade/
├── backend/                 # Go backend
│   ├── cmd/server/         # Application entry point
│   ├── internal/           # Internal packages
│   │   ├── handler/        # HTTP handlers
│   │   ├── service/        # Business logic
│   │   ├── repository/     # Data access layer
│   │   ├── model/          # Data models
│   │   └── middleware/     # HTTP middleware
│   ├── config.yaml         # Configuration file
│   └── go.mod              # Go dependencies
├── frontend/               # Vue.js frontend
│   ├── src/
│   │   ├── components/     # Vue components
│   │   ├── views/          # Page views
│   │   ├── stores/         # Pinia stores
│   │   ├── types/          # TypeScript types
│   │   └── main.ts         # Application entry
│   ├── package.json        # Dependencies
│   └── vite.config.ts      # Build configuration
├── docker-compose.yml      # Docker composition
└── README.md              # This file
```

### 🔧 Configuration

Backend configuration is managed through `backend/config.yaml`:

```yaml
server:
  port: 8080
  mode: debug

database:
  driver: sqlite
  dsn: "./data.db"

cors:
  allowed_origins:
    - "http://localhost:3000"
    - "http://localhost:3001"
    - "http://localhost:3002"
```

### 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### 📄 License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

### 🐛 Issues & Support

If you encounter any issues or have questions, please:
1. Check existing [issues](https://github.com/ShawnMa123/razor-blade/issues)
2. Create a new issue with detailed information
3. Include steps to reproduce the problem

---

## 中文文档

一个全面的剃须刀和刀片使用记录管理系统，支持库存管理和使用模式分析。

### 🌟 功能特性

- **📊 使用追踪**: 记录详细的剃须过程，包含评分和体验
- **📦 库存管理**: 更换刀片时自动减少库存数量
- **📅 日历视图**: 可视化日历显示使用历史和评分
- **📈 统计仪表板**: 全面的分析和使用洞察
- **🔄 快速记录**: 快速直观的使用记录创建
- **💾 数据存储**: 支持 SQLite 数据库和内存存储
- **🎨 现代界面**: 使用 Element Plus 构建的清洁响应式界面
- **📱 移动友好**: 针对桌面和移动设备优化

### 🛠️ 技术栈

#### 后端
- **Go 1.23+** - 高性能后端服务器
- **Gin** - HTTP Web 框架
- **GORM** - 数据库 ORM
- **SQLite** - 轻量级数据库（支持内存模式）
- **Viper** - 配置管理
- **Logrus** - 结构化日志

#### 前端
- **Vue 3** - 渐进式 JavaScript 框架
- **TypeScript** - 类型安全开发
- **Element Plus** - Vue 3 UI 组件库
- **Pinia** - 状态管理
- **Vue Router** - 客户端路由
- **Axios** - HTTP 客户端
- **ECharts** - 数据可视化
- **Day.js** - 日期处理

### 🚀 快速开始

#### 环境要求
- **Go 1.23+**
- **Node.js 18+**
- **npm** 或 **yarn**

#### 安装步骤

1. **克隆仓库**
```bash
git clone https://github.com/ShawnMa123/razor-blade.git
cd razor-blade
```

2. **启动后端服务器**
```bash
cd backend
go mod tidy
go run cmd/server/main.go
```

3. **启动前端开发服务器**
```bash
cd frontend
npm install
npm run dev
```

4. **访问应用**
- 前端: http://localhost:3000
- 后端 API: http://localhost:8080

#### Docker 部署

```bash
# 使用 Docker Compose 构建和运行
docker-compose up --build -d

# 访问应用
# 前端: http://localhost:3000
# 后端 API: http://localhost:8080
```

### 📖 使用指南

#### 添加剃须刀和刀片
1. 导航到管理部分
2. 添加剃须刀信息（品牌、型号等）
3. 添加刀片类型和库存数量

#### 记录使用
1. 使用仪表板上的快速记录表单
2. 选择剃须刀和刀片组合
3. 添加评分和体验备注
4. 标记是否更换刀片（自动减少库存）

#### 查看统计
- 查看仪表板了解使用概览
- 查看日历了解历史记录
- 分析使用模式和偏好

### 🏗️ 架构设计

```
razor-blade/
├── backend/                 # Go 后端
│   ├── cmd/server/         # 应用入口
│   ├── internal/           # 内部包
│   │   ├── handler/        # HTTP 处理器
│   │   ├── service/        # 业务逻辑
│   │   ├── repository/     # 数据访问层
│   │   ├── model/          # 数据模型
│   │   └── middleware/     # HTTP 中间件
│   ├── config.yaml         # 配置文件
│   └── go.mod              # Go 依赖
├── frontend/               # Vue.js 前端
│   ├── src/
│   │   ├── components/     # Vue 组件
│   │   ├── views/          # 页面视图
│   │   ├── stores/         # Pinia 状态管理
│   │   ├── types/          # TypeScript 类型
│   │   └── main.ts         # 应用入口
│   ├── package.json        # 依赖配置
│   └── vite.config.ts      # 构建配置
├── docker-compose.yml      # Docker 编排
└── README.md              # 本文件
```

### 🔧 配置说明

后端配置通过 `backend/config.yaml` 管理：

```yaml
server:
  port: 8080
  mode: debug

database:
  driver: sqlite
  dsn: "./data.db"

cors:
  allowed_origins:
    - "http://localhost:3000"
    - "http://localhost:3001"
    - "http://localhost:3002"
```

### 🤝 贡献指南

1. Fork 仓库
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

### 📄 许可证

本项目使用 GNU General Public License v3.0 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

### 🐛 问题与支持

如果遇到问题或有疑问，请：
1. 查看现有 [issues](https://github.com/ShawnMa123/razor-blade/issues)
2. 创建新 issue 并提供详细信息
3. 包含重现问题的步骤

---

**Made with ❤️ by ShawnMa123**