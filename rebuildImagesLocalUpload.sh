#!/bin/bash
git pull
docker build  -t gva-yarn-cache -f ./web/Dockerfile-cache ./web
docker-compose -f deploy/docker-compose/docker-compose.yaml build
# 本地目录和远程服务器信息
#LOCAL_DIR="/home/qing/programs/myOthers/test/gva_test/livestock-automation-hub_-backend-and-admin"
REMOTE_HOST="服务器ip或域名"
REMOTE_USER="服务器登录帐号"
PASSWORD="test@abcCBA"
REMOTE_DIR="目标文件夹"
TAR_FILE="本地打包为什么文件名字(建议以.tar结尾)"
COMPOSE_FILE="docker-compose.yaml" #运行时的docker-compose设置不是打包时候的
NGINX_FILE="web/.docker-compose/nginx/conf.d/my.conf" #注意这些地方一定不要有空格哪怕等号左右都不要有
CONFIG_FILE="server/config.docker.yaml"  #后端配置文件
# 为没有标签的镜像重新打标签 如果docker-compose里没有做那么需要打上标签
# docker tag 2bb37f74d436 gva-server:latest
# docker tag 5e542ea21b0c gva-web:latest

# 提取所有镜像名并保存为 tar 文件
#docker save -o "$TAR_FILE" mysql:8.0.21 redis:6.0.6 gva-server:latest gva-web:latest
#不提取所有只提取更新的
docker save -o "$TAR_FILE" gva-server:latest gva-web:latest
# 创建远程目录
sshpass -p "$PASSWORD" ssh -o StrictHostKeyChecking=no "$REMOTE_USER@$REMOTE_HOST" "mkdir -p $REMOTE_DIR"

# 将项目文件、tar 文件和 docker-compose 文件传输到云服务器
# 直接使用images不用上传本地项目文件
#sshpass -p "$PASSWORD" scp -o StrictHostKeyChecking=no -r "$LOCAL_DIR" "$REMOTE_USER@$REMOTE_HOST:$REMOTE_DIR"
#sshpass -p "$PASSWORD" scp -o StrictHostKeyChecking=no "$TAR_FILE" "$COMPOSE_FILE" "$NGINX_FILE" "$REMOTE_USER@$REMOTE_HOST:$REMOTE_DIR"
#如果有后端服务器的config.docker.yaml需要上传
sshpass -p "$PASSWORD" scp -o StrictHostKeyChecking=no "$TAR_FILE" "$COMPOSE_FILE" "$NGINX_FILE" "$CONFIG_FILE" "$REMOTE_USER@$REMOTE_HOST:$REMOTE_DIR"

# 登录到云服务器，加载镜像并启动服务
sshpass -p "$PASSWORD" ssh -o StrictHostKeyChecking=no "$REMOTE_USER@$REMOTE_HOST" << EOF
cd "$REMOTE_DIR"
docker load -i "$TAR_FILE"
docker-compose down
docker-compose up -d
EOF

# 删除本地的 tar 文件
# rm "$TAR_FILE"
