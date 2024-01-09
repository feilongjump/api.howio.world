FROM golang:alpine as builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 将代码复制到容器中
COPY . /app

# 移动到工作目录：/app
WORKDIR /app

# 将我们的代码编译成二进制可执行文件 howio
RUN go build -o howio .


########## 创建一个小镜像 ########
FROM alpine:latest

ARG TZ=UTC
ENV TZ ${TZ}

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 配置文件
COPY env.local.toml /env.local.toml

# 从 builder 镜像中把 /app/howio 拷贝到当前目录
COPY --from=builder /app/howio /howio

# 声明服务端口
EXPOSE 8088

# 启动容器时运行的命令
CMD ["/howio"]
