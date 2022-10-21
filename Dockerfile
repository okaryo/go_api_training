FROM golang:1.19-alpine

RUN apk update && apk add --no-cache alpine-sdk build-base git

COPY ./ /go/src/app

WORKDIR /go/src/app
