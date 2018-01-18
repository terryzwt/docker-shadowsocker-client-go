FROM golang:1.9-alpine
MAINTAINER zterry <zterry@qq.com>

RUN apk add --no-cache git
RUN go get github.com/shadowsocks/shadowsocks-go/cmd/shadowsocks-local

EXPOSE 1080

CMD shadowsocks-local -c /etc/shadowsocks/config.json -m rc4-md5
