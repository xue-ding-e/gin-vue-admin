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

#PRODUCTION 设置代理环境变量
#ENV HTTP_PROXY=http://192.168.1.148:7890
#ENV HTTPS_PROXY=http://192.168.1.148:7890
#PRODUCTION 根据自己的项目修改实际的代理或者不需要   设置为env -w GOPROXY=https://goproxy.cn,direct  上面是默认
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

#RUN go env -w GO111MODULE=on \
#    && go env -w CGO_ENABLED=0 \
#    && go env \
#    && go mod tidy \
#    && go build -o server .

#PRODUCTION 取消代理环境变量
#RUN unset HTTP_PROXY HTTPS_PROXY
# ==================================================== 以下为多阶段构建 ==========================================================

# 声明镜像来源为alpine:latest

FROM alpine:latest
#PRODUCTION 根据生产环境实际配置邮箱
# 镜像编写者及邮箱
LABEL MAINTAINER="781408517@qq.com"
#ENV HTTP_PROXY=http://192.168.1.148:7890
#ENV HTTPS_PROXY=http://192.168.1.148:7890
# 设置时区
ENV TZ=Asia/Shanghai
#PRODUCTION 临时设置代理环境变量来下载包
RUN export HTTP_PROXY=http://192.168.1.4:7890 \
    && export HTTPS_PROXY=http://192.168.1.4:7890 \
    && apk update \
    && apk add --no-cache tzdata openntpd \
    && unset HTTP_PROXY HTTPS_PROXY \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
#PRODUCTION 取消代理环境变量
RUN unset HTTP_PROXY HTTPS_PROXY

# 安装 tzdata 和 openntpd
#RUN apk update && apk add --no-cache tzdata openntpd \
#    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone \

# 声明工作目录
WORKDIR /go/src/gin-vue-admin

# 把/go/src/gin-vue-admin中的可执行文件以及配置文件(resource模板文件)添加进入docker中
COPY --from=0 /go/src/gin-vue-admin/server ./
COPY --from=0 /go/src/gin-vue-admin/config.docker.yaml ./
COPY --from=0 /go/src/gin-vue-admin/resource ./

# 挂载目录：如果使用了sqlite数据库，容器命令示例：docker run -d -v /宿主机路径/gva.db:/go/src/github.com/flipped-aurora/gin-vue-admin/server/gva.db -p 8888:8888 --name gva-server-v1 gva-server:1.0
# VOLUME ["/go/src/github.com/flipped-aurora/gin-vue-admin/server"]

EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml
