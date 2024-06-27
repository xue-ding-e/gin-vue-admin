# 从cache镜像中中获取构建
FROM cows-yarn-cache:latest as build
WORKDIR /gva_web/
# 拷贝整个web项目到当前工作目录
COPY . .

# PRODUCTION 根据实际生产更改代理
ENV HTTP_PROXY=http://192.168.1.148:7890
ENV HTTPS_PROXY=http://192.168.1.148:7890

RUN  yarn && yarn build

# ==================================================== 以下为多阶段构建 ==========================================================

# 声明镜像来源为alpine:latest
FROM nginx:alpine
#PRODUCTION 根据实际情况修改邮箱(前端)
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"
ENV HTTP_PROXY=http://192.168.1.148:7890
ENV HTTPS_PROXY=http://192.168.1.148:7890

COPY .docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf
#PRODUCTION 填写证书位置
# 拷贝证书文件到容器内的指定位置
#COPY certificate/ai-everywhere.cn_bundle.crt /etc/nginx/ssl/certificate.crt
#COPY certificate/ai-everywhere.cn.key /etc/nginx/ssl/private.key

# 从第一阶段进行拷贝文件
COPY --from=0 /gva_web/dist /usr/share/nginx/html
# 查看/etc/nginx/nginx.conf文件
RUN cat /etc/nginx/nginx.conf
# 查看 /etc/nginx/conf.d/my.conf
RUN cat /etc/nginx/conf.d/my.conf
# 查看 文件是否拷贝成功
RUN ls -al /usr/share/nginx/html
