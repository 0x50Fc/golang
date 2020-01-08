FROM mysql:5.6

RUN echo "Asia/shanghai" >> /etc/timezone

COPY localtime /etc/localtime

COPY ./mysql.cnf /etc/mysql/conf.d/mysql.cnf

