FROM golang:1.19.1-alpine

RUN apk update && apk add git

RUN mkdir -p /go/src/github.com/app-plate

WORKDIR /go/src/github.com/app-plate

ADD . /go/src/github.com/app-plate
