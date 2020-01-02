#!/bin/bash

TAG=latest
PROJECT=hailongz/kk-user

rm -rf main
echo "[GO BUILD] [$PROJECT:$TAG] >>"
docker run --rm -v `pwd`:/main:rw -v $GOPATH:/go:rw hailongz/kk-gobuild:latest go build
echo "[OK]"
echo "[DOCKER BUILD] [$PROJECT:$TAG] >>"
docker build -t $PROJECT:$TAG .
echo "[OK]"
echo "[DOCKER PUSH] [$PROJECT:$TAG] >>"
docker push $PROJECT:$TAG
echo "[OK]"

