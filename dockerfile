FROM golang:alpine AS builder_test

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /app
COPY ./ffmpeg_demo.go ./ffmpeg_demo.go
COPY ./go.mod go.sum ./
RUN go mod download
RUN go build -ldflags="-s -w" -o /app/output


FROM golang:alpine AS object

RUN apk update && apk add --no-cache ffmpeg

WORKDIR /app
COPY ./gura.mp4 ./app/gura.mp4
COPY --from=builder_test /app/output /app/output

