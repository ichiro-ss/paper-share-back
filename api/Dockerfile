FROM golang:1.21

RUN apk update && apk add git

RUN mkdir -p /go/src/github.com/api

WORKDIR /go/src/github.com/api

ADD . /go/src/github.com/api
