FROM golang

RUN apt-get update
RUN apt-get install -y gcc libc-dev
RUN apt-get install -y make
RUN apt-get install -y nginx

ENV SERVER_PATH=/opt/server

ADD dist ${SERVER_PATH}

RUN ls ${SERVER_PATH}

RUN mv ${SERVER_PATH}/dist/html /usr/share/nginx/

WORKDIR ${SERVER_PATH}/dist

EXPOSE 80
EXPOSE 8080

ENTRYPOINT ./graduation