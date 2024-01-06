FROM golang:1.21

RUN apk update && apk add git

RUN mkdir -p /go/src/github.com/app-plate

WORKDIR /go/src/github.com/app-plate

ADD . /go/src/github.com/app-plate
