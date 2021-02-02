from golang:1.14-alpine as builder
# 签名
MAINTAINER xingzheng "3024979463@qq.com"
# RUN rm -rf /etc/apt/sources.list
# ADD ./sources.list /etc/apt/
# install xz
# RUN apt-key adv --recv-keys --keyserver keyserver.ubuntu.com 3B4FE6ACC0B21F32 && apt-get update && apt-get install -y xz-utils 
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
# && rm -rf /var/lib/apt/lists/*
# install UPX
# ADD ./upx /bin
# RUN chmod a+x /bin/upx

WORKDIR $GOPATH/src
ADD . $GOPATH/src

RUN mkdir -p $GOPATH/src/golang.org/x

RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go mod tidy && go build -o main ./main.go && pwd


# use a minimal alpine image
FROM alpine:latest
# add ca-certificates in case you need them
RUN apk add --no-cache ca-certificates
# set working directory
WORKDIR /root
# copy the binary from builder
COPY --from=builder /go/src/main .
COPY --from=builder /go/src/config/application.yml ./config/
COPY ./make-dir.sh /root
RUN chmod a+x /root/make-dir.sh 

EXPOSE 5556

CMD sh /root/make-dir.sh && ls /home && ./main
