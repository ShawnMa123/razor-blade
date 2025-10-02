# Razor Blade 本地运行和调试指南

本文档详细说明如何在本地环境中运行和调试 Razor Blade 项目。

## 前置要求

### 必需软件

1. **Go 1.21+**
   - 下载地址: https://golang.org/dl/
   - 验证安装: `go version`

2. **Node.js 18+**
   - 下载地址: https://nodejs.org/
   - 验证安装: `node --version` 和 `npm --version`

3. **Git**
   - 下载地址: https://git-scm.com/
   - 验证安装: `git --version`

### 可选软件

1. **Docker & Docker Compose** (用于容器化运行)
   - Docker Desktop: https://www.docker.com/products/docker-desktop
   - 验证安装: `docker --version` 和 `docker-compose --version`

## 运行方式

### 方式一：本地直接运行 (推荐开发使用)

这种方式适合日常开发，支持热重载，便于调试。

#### 步骤 1: 准备项目

```bash
# 克隆项目 (如果还没有)
git clone <your-repository-url>
cd razor-blade

# 查看项目结构
ls -la
```

#### 步骤 2: 启动后端服务

```bash
# 进入后端目录
cd backend

# 下载Go依赖
go mod tidy

# 创建数据存储目录
mkdir -p data

# 检查配置文件
cat config.yaml

# 启动后端服务
go run cmd/server/main.go
```

**预期输出:**
```
INFO[2024-01-15T10:30:00Z] Starting Razor-Blade server...
INFO[2024-01-15T10:30:00Z] Database connected successfully
INFO[2024-01-15T10:30:00Z] Database migration completed
INFO[2024-01-15T10:30:00Z] Server starting on port 8080
```

**验证后端运行:**
```bash
# 健康检查
curl http://localhost:8080/health

# 预期响应
{"success":true,"data":{"status":"healthy"},"message":"服务正常运行"}
```

#### 步骤 3: 启动前端服务

打开新的终端窗口：

```bash
# 进入前端目录
cd frontend

# 安装npm依赖
npm install

# 启动开发服务器
npm run dev
```

**预期输出:**
```
  VITE v4.5.0  ready in 500 ms

  ➜  Local:   http://localhost:3000/
  ➜  Network: use --host to expose
  ➜  press h to show help
```

**访问应用:**
- 前端界面: http://localhost:3000
- 后端API: http://localhost:8080

#### 步骤 4: 验证功能

1. 打开浏览器访问 http://localhost:3000
2. 应该看到 Razor Blade 的主界面
3. 尝试添加一个剃须刀
4. 尝试添加一个刀片
5. 尝试创建一条使用记录

### 方式二：Docker Compose 运行

这种方式接近生产环境，适合测试完整部署流程。

#### 开发环境运行

```bash
# 在项目根目录下
cd razor-blade

# 使用开发配置启动 (后端端口会暴露)
docker-compose -f docker-compose.dev.yml up --build

# 后台运行
docker-compose -f docker-compose.dev.yml up --build -d
```

#### 生产环境运行

```bash
# 使用生产配置启动
docker-compose up --build

# 后台运行
docker-compose up --build -d
```

#### Docker 管理命令

```bash
# 查看运行状态
docker-compose ps

# 查看日志
docker-compose logs

# 实时跟踪日志
docker-compose logs -f

# 停止服务
docker-compose down

# 停止服务并清理数据
docker-compose down -v

# 重建服务
docker-compose build --no-cache
```

## 开发调试

### 后端调试

#### 使用 VS Code 调试

创建 `.vscode/launch.json`:

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Razor Blade Backend",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/backend/cmd/server/main.go",
            "cwd": "${workspaceFolder}/backend",
            "env": {
                "RAZOR_BLADE_SERVER_MODE": "debug",
                "RAZOR_BLADE_LOG_LEVEL": "debug"
            }
        }
    ]
}
```

#### 使用 Delve 调试器

```bash
# 安装 delve
go install github.com/go-delve/delve/cmd/dlv@latest

# 启动调试服务器
cd backend
dlv debug cmd/server/main.go

# 在调试器中设置断点并运行
(dlv) break main.main
(dlv) continue
```

#### 日志调试

修改 `backend/config.yaml` 启用详细日志:

```yaml
log:
  level: "debug"
  format: "text"
```

### 前端调试

#### 使用浏览器开发者工具

1. 打开 Chrome/Firefox 开发者工具 (F12)
2. 在 Network 标签查看 API 请求
3. 在 Console 标签查看错误信息
4. 使用 Vue DevTools 扩展调试组件状态

#### 使用 VS Code 调试

安装推荐的VS Code扩展:
- Vue Language Features (Volar)
- TypeScript Vue Plugin (Volar)
- Vue VSCode Snippets

## 环境变量配置

### 后端环境变量

```bash
# 服务器配置
export RAZOR_BLADE_SERVER_PORT=8080
export RAZOR_BLADE_SERVER_MODE=debug  # debug, release, test

# 数据库配置
export RAZOR_BLADE_DATABASE_PATH=./data/razor-blade.db

# 日志配置
export RAZOR_BLADE_LOG_LEVEL=debug    # debug, info, warn, error
export RAZOR_BLADE_LOG_FORMAT=text    # text, json
```

### 前端环境变量

创建 `frontend/.env.local`:

```env
# API基础URL (开发环境通常不需要设置，会使用vite.config.ts中的代理)
VITE_API_BASE_URL=http://localhost:8080/api/v1

# 开发模式
VITE_MODE=development
```

## 常见问题解决

### 1. 后端问题

**问题: 端口 8080 被占用**
```bash
# 查找占用进程
netstat -tulpn | grep :8080
# 或
lsof -i :8080

# 杀死进程 (谨慎操作)
kill -9 <PID>

# 或使用其他端口
export RAZOR_BLADE_SERVER_PORT=8081
```

**问题: Go 模块下载失败**
```bash
# 设置Go代理 (中国用户)
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOSUMDB=sum.golang.google.cn

# 重新下载依赖
go mod tidy
```

**问题: 数据库文件权限错误**
```bash
# 检查数据目录
ls -la backend/data/

# 修复权限
chmod 755 backend/data/
chmod 644 backend/data/razor-blade.db  # 如果文件已存在
```

### 2. 前端问题

**问题: npm 安装依赖失败**
```bash
# 清理缓存
npm cache clean --force

# 删除 node_modules 重新安装
rm -rf node_modules package-lock.json
npm install

# 或使用 yarn
yarn install
```

**问题: 前端无法连接后端**
- 检查后端服务是否正常运行 (`curl http://localhost:8080/health`)
- 检查 `frontend/vite.config.ts` 中的代理配置
- 查看浏览器开发者工具的网络请求

**问题: TypeScript 编译错误**
```bash
# 检查 TypeScript 配置
npx tsc --noEmit

# 安装类型定义
npm install --save-dev @types/node
```

### 3. Docker 问题

**问题: Docker 构建失败**
```bash
# 清理 Docker 缓存
docker system prune -a

# 查看构建日志
docker-compose build --no-cache --progress=plain

# 单独构建服务
docker-compose build backend
docker-compose build frontend
```

**问题: 容器无法启动**
```bash
# 查看容器日志
docker-compose logs backend
docker-compose logs frontend

# 进入容器调试
docker-compose exec backend sh
docker-compose exec frontend sh
```

**问题: 数据不持久化**
- 检查 `docker-compose.yml` 中的 volumes 配置
- 确保数据卷映射正确

### 4. 网络问题

**问题: API 请求跨域错误**
- 检查后端 CORS 配置 (`backend/internal/middleware/middleware.go`)
- 确保允许前端域名访问

**问题: 前端路由 404**
- 检查 Vue Router 配置
- 确保 Nginx 配置支持 SPA 路由

## 性能优化建议

### 开发环境优化

1. **启用Go模块代理**
   ```bash
   go env -w GOPROXY=https://goproxy.cn,direct
   ```

2. **使用Node.js包管理器缓存**
   ```bash
   npm config set registry https://registry.npmmirror.com
   ```

3. **IDE配置优化**
   - 安装相关语言插件
   - 配置代码格式化
   - 启用实时错误检查

### 运行时优化

1. **后端优化**
   - 生产环境使用 `release` 模式
   - 适当调整日志级别
   - 监控数据库性能

2. **前端优化**
   - 生产构建启用压缩
   - 使用 CDN 加速静态资源
   - 启用浏览器缓存

## 安全注意事项

1. **不要提交敏感信息**
   - `.env` 文件包含敏感配置
   - 数据库文件包含用户数据

2. **生产环境配置**
   - 修改默认端口
   - 启用 HTTPS
   - 设置防火墙规则

3. **定期备份数据**
   ```bash
   # 备份数据库
   cp backend/data/razor-blade.db backup/razor-blade-$(date +%Y%m%d).db
   ```

## 获取帮助

如果遇到本文档未覆盖的问题:

1. 检查项目的 GitHub Issues
2. 查看相关技术文档:
   - [Go 官方文档](https://golang.org/doc/)
   - [Vue.js 官方文档](https://vuejs.org/)
   - [Docker 官方文档](https://docs.docker.com/)
3. 提交新的 Issue 描述问题