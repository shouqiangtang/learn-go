FROM alpine:latest

RUN apk update && apk add bash

ADD ./localip-linux-amd64 /app

RUN chmod 777 /app

ENTRYPOINT [ "/app" ]