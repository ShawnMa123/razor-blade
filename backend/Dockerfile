# 使用官方Go镜像作为构建环境
FROM golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装必要的系统包
RUN apk add --no-cache git gcc musl-dev sqlite-dev

# 复制go.mod和go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用 - 使用兼容的 SQLite 构建标签
RUN CGO_ENABLED=1 GOOS=linux \
    CGO_CFLAGS="-D_LARGEFILE64_SOURCE" \
    go build -tags "sqlite_omit_load_extension" \
    -ldflags="-s -w" \
    -o main ./cmd/server

# 使用轻量级的alpine镜像作为运行环境
FROM alpine:latest

# 安装必要的运行时包
RUN apk --no-cache add ca-certificates sqlite

# 设置工作目录
WORKDIR /app

# 从构建镜像复制编译好的二进制文件
COPY --from=builder /app/main .
COPY --from=builder /app/config.yaml .

# 创建数据目录
RUN mkdir -p /app/data

# 设置权限
RUN chmod +x ./main

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./main"]