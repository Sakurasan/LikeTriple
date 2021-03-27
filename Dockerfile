FROM  golang:1.15.9-alpine as builder
LABEL maintainer="C菌"
WORKDIR  /go/src
COPY  .  .
RUN GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o  LikeTriple


FROM python:3-alpine
# RUN apk add -U tzdata
# RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
 && apk --no-cache add ffmpeg \
 && pip3 install --no-cache-dir you-get \
 && mkdir -p /download && mkdir -p /app

WORKDIR /app
COPY  --from=builder /go/src/LikeTriple  .

CMD  ["./LikeTriple"]
