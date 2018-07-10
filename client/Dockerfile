FROM golang:rc-alpine
MAINTAINER Ankur Gupta

RUN apk update && apk upgrade

WORKDIR /app

ENTRYPOINT go build -o bin/client src/client.go && bin/client
