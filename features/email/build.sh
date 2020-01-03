#!/bin/bash

TAG=latest
PROJECT=hailongz/kk-email

rm -rf main
echo "[GO BUILD] [$PROJECT:$TAG] >>"
docker run --rm -v `pwd`:/main:rw -v $GOPATH:/go:rw hailongz/kk-gobuild:latest go build
echo "[OK]"
