#!/bin/bash

set -v

docker build . -t hailongz/kk-golang:latest
docker push hailongz/kk-golang:latest
