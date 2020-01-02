
FROM golang:1.13.5-alpine

RUN echo "Asia/shanghai" >> /etc/timezone

COPY localtime /etc/localtime

RUN apk add gcc musl-dev git && \
    rm -rf /var/cache/apk/* 

RUN go env

RUN go version

WORKDIR /main


