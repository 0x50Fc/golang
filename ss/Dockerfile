FROM hailongz/kk-golang:latest

COPY ./main /bin/kk-app

RUN chmod +x /bin/kk-app

COPY ./app.json /app.josn

RUN mkdir /webapp/

EXPOSE 80

ENV KK_CONFIG_FILE /app.json

ENV KK_NODE node-0

ENV KK_VERSION 1.0

WORKDIR /webapp/

CMD kk-app
