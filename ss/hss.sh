#!/bin/bash

TAG=1.0
PROJECT=registry.cn-shenzhen.aliyuncs.com/hss/hss-kk-app

echo "[GO BUILD] [$PROJECT:$TAG] >>"
docker run --rm -v `pwd`:/main:rw -v $GOPATH:/go:rw registry.dpool.sina.com.cn/kk/kk-gobuild:latest go build
echo "[OK]"
echo "[DOCKER BUILD] [$PROJECT:$TAG] >>"
docker build -t $PROJECT:$TAG .
echo "[OK]"
echo "[DOCKER PUSH] [$PROJECT:$TAG] >>"
docker push $PROJECT:$TAG
echo "[OK]"
rm -rf main

