FROM registry.dpool.sina.com.cn/kk/kk-golang:latest

COPY ./main /bin/kk-notice

RUN chmod +x /bin/kk-notice

COPY ./app.json /app.josn

EXPOSE 80

ENV KK_CONFIG_FILE /app.json

ENV KK_NODE node-0

ENV KK_VERSION 1.0

WORKDIR /

CMD kk-notice
