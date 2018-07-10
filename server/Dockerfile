FROM golang:rc-alpine
MAINTAINER Ankur Gupta

RUN apk update && apk upgrade
RUN apk add --no-cache git mercurial

RUN go get github.com/julienschmidt/httprouter

WORKDIR /app

ENTRYPOINT go build -o bin/server src/server.go && bin/server
