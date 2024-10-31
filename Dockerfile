# 使用官方 Go 镜像
FROM golang:1.20-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建微服务
RUN go build -o lucianaUserServer .

# 使用更小的镜像
FROM alpine:latest
WORKDIR /root/

# 复制二进制文件
COPY --from=builder /app/lucianaUserServer .

# 暴露端口
EXPOSE 50051  # 根据你的 gRPC 服务端口设置

# 运行微服务
CMD ["./lucianaUserServer"]
