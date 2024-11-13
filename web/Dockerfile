# node:20
# 从cache镜像中中获取构建
FROM cows-yarn-cache:latest as build
# 声明工作目录
WORKDIR /gva_web/

# 拷贝整个web项目到当前工作目录
COPY . .

# ENV HTTP_PROXY=http://192.168.1.4:7890
# ENV HTTPS_PROXY=http://192.168.1.4:7890

RUN  yarn && yarn build

# RUN unset HTTP_PROXY && unset HTTPS_PROXY


# ===================================================== 以下为多阶段构建 ==========================================================

# 声明镜像来源为nginx:alpine, alpine 镜像小
FROM nginx:alpine

# 镜像编写者及邮箱
LABEL MAINTAINER="781408517@qq.com"

# 不设置运行阶段的代理环境变量
# 从.docker-compose/nginx/conf.d/目录拷贝my.conf到容器内的/etc/nginx/conf.d/my.conf
COPY .docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf

# 拷贝证书文件到容器内的指定位置
COPY certificate/ai-everywhere.cn_bundle.crt /etc/nginx/ssl/certificate.crt
COPY certificate/ai-everywhere.cn.key /etc/nginx/ssl/private.key

# 从第一阶段进行拷贝文件
COPY --from=0 /gva_web/dist /usr/share/nginx/html

# 查看/etc/nginx/nginx.conf文件
RUN cat /etc/nginx/nginx.conf

# 查看 /etc/nginx/conf.d/my.conf
RUN cat /etc/nginx/conf.d/my.conf

# 查看 文件是否拷贝成功
RUN ls -al /usr/share/nginx/html
