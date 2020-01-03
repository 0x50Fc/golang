FROM hailongz/kk-golang:latest

COPY ./main /bin/kk-email

RUN chmod +x /bin/kk-email

COPY ./app.json /app.josn

EXPOSE 80

ENV KK_CONFIG_FILE /app.json

WORKDIR /

CMD kk-email $KK_CONFIG_FILE
