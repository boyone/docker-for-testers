FROM golang:1.12 as builder
WORKDIR /module

FROM alpine:3
RUN apk add tzdata && \
    cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && \
    echo "Asia/Bangkok" >  /etc/timezone && \
    apk del tzdata
WORKDIR /root/
ENV GIN_MODE release
