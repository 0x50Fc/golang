#!/bin/bash

set -v

docker build . -t registry.cn-beijing.aliyuncs.com/zkr/kk-phpmyadmin:latest
docker push registry.cn-beijing.aliyuncs.com/zkr/kk-phpmyadmin:latest
