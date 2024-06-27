# 声明镜像来源为golang:alpine
FROM golang:alpine as builder
# 声明工作目录
WORKDIR /go/src/gin-vue-admin
# 拷贝整个server项目到工作目录
COPY . .

#官方默认
#RUN go env -w GO111MODULE=on \
#    && go go env -w GOPROXY=https://goproxy.cn,direct \
#    && go env -w CGO_ENABLED=0 \
#    && go env \
#    && go mod tidy \
#    && go build -o server .

#PRODUCTION
ENV HTTP_PROXY=http://192.168.1.148:7890
ENV HTTPS_PROXY=http://192.168.1.148:7890
#PRODUCTION 根据自己的项目修改实际的代理或者不需要   设置为env -w GOPROXY=https://goproxy.cn,direct  上面是默认
RUN go env -w GO111MODULE=on \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

# ==================================================== 以下为多阶段构建 ==========================================================

# 声明镜像来源为alpine:latest

FROM alpine:latest
#PRODUCTION 根据生产环境实际配置邮箱
# 镜像编写者及邮箱
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"
ENV HTTP_PROXY=http://192.168.1.148:7890
ENV HTTPS_PROXY=http://192.168.1.148:7890
# 设置时区
ENV TZ=Asia/Shanghai
# 安装 tzdata 和 openntpd
RUN apk update && apk add --no-cache tzdata openntpd \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone \
# 声明工作目录
WORKDIR /go/src/gin-vue-admin

COPY --from=0 /go/src/gin-vue-admin ./

EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml

