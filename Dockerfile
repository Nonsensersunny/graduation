FROM golang

RUN apt-get update
RUN apt-get install -y gcc libc-dev
RUN apt-get install -y make
RUN apt-get install -y nginx

ENV SERVER_PATH=/opt/server

ADD dist ${SERVER_PATH}

RUN pwd

RUN ls ${SERVER_PATH}

WORKDIR ${SERVER_PATH}

EXPOSE 80
EXPOSE 8080

CMD ["nginx", "start"]
ENTRYPOINT ./dist/graduation