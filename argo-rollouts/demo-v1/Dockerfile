# 设置基础镜像
FROM golang:1.20.7-alpine3.17 as builder

# 设置工作目录
WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on

# cache modules

# 将代码复制到容器中
COPY . .

# 构建可执行文件
RUN CGO_ENABLED=0 go build -o main .


# 暴露端口
EXPOSE 8081

# 运行应用程序
CMD ["./main"]