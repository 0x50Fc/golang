FROM alpine:3.11

RUN echo "Asia/shanghai" >> /etc/timezone

COPY localtime /etc/localtime

RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/*

RUN update-ca-certificates

