FROM golang:1.22-rc-alpine

RUN apk update && apk add git

RUN mkdir -p /go/src/github.com/api

WORKDIR /go/src/github.com/api

ADD . /go/src/github.com/api
