FROM  golang:alpine as builder
LABEL maintainer="C菌"
WORKDIR  /go/src
COPY  .  .
ENV GO111MODULE=off
RUN go build  -o  LikeTriple


FROM python:3-alpine
# RUN apk add -U tzdata
# RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
RUN  apk --no-cache add ffmpeg \
 && pip3 install --no-cache-dir you-get \
 && mkdir -p /download && mkdir -p /app

WORKDIR /app
COPY  --from=builder /go/src/LikeTriple  .

CMD  ["./LikeTriple"]
