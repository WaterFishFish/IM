#!/bin/bash

reso_addr='registry.cn-hangzhou.aliyuncs.com/lazy-chat/social-rpc-dev'
tag='latest'

pod_idb="192.168.88.128"

container_name="lazy-chat-social-rpc-test"

docker stop ${container_name}

docker rm ${container_name}

docker rmi ${reso_addr}:${tag}

docker pull ${reso_addr}:${tag}


# 如果需要指定配置文件的
# docker run -p 10001:8080 --network imooc_easy-im -v /easy-im/config/user-rpc:/user/conf/ --name=${container_name} -d ${reso_addr}:${tag}
docker run -p 10001:10001 -e pod_IP=${pod_idb} --name=${container_name} -d ${reso_addr}:${tag}