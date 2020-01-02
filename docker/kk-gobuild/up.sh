#!/bin/bash

set -v

docker build . -t hailongz/kk-gobuild:latest
docker push hailongz/kk-gobuild:latest
