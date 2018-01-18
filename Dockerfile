FROM golang

MAINTAINER zterry <zterry@qq.com>

RUN go get github.com/shadowsocks/shadowsocks-go/cmd/shadowsocks-local

EXPOSE 1080

CMD shadowsocks-local -c /etc/shadowsocks/config.json -m rc4-md5
