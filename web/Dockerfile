FROM node:20.11.0

WORKDIR /gva_web/
# 拷贝整个web项目到当前工作目录
COPY . .

# TODO PRODUCTION 根据实际生产更改代理
# --unsafe-perm 允许以root权限运行
#RUN npm config set registry https://192.168.1.148:7890 \
#    && npm install --unsafe-perm \
#    && RUN npm run build

#这个用不了 出错问淼子
RUN npm config set registry https://192.168.1.148:7890 && yarn && yarn build

FROM nginx:alpine
#TODO PRODUCTION 根据实际情况修改邮箱(前端)
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

COPY .docker-compose/nginx/conf.d/my.conf /etc/nginx/conf.d/my.conf
#TODO PRODUCTION 填写证书位置
# 拷贝证书文件到容器内的指定位置
#COPY certificate/ai-everywhere.cn_bundle.crt /etc/nginx/ssl/certificate.crt
#COPY certificate/ai-everywhere.cn.key /etc/nginx/ssl/private.key

# 从第一阶段进行拷贝文件
COPY --from=0 /gva_web/dist /usr/share/nginx/html
RUN cat /etc/nginx/nginx.conf
RUN cat /etc/nginx/conf.d/my.conf
RUN ls -al /usr/share/nginx/html
