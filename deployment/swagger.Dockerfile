FROM golang:1.12-alpine AS build-env
ENV GO111MODULE=on

WORKDIR /go/src/swagger
COPY swagger/ /go/src/swagger
RUN apk update \
&& apk add --no-cache git \
&& rm -rf /var/cache/apk/* \
&& rm -rf /tmp/*
RUN go build -o main
ENTRYPOINT ./main