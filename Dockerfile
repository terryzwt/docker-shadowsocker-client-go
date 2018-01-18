FROM golang:1.9-alpine
MAINTAINER zterry <zterry@qq.com>

RUN apk add --no-cache git polipo supervisor
RUN go get github.com/shadowsocks/shadowsocks-go/cmd/shadowsocks-local

ADD polipo.conf /etc/polipo.conf
ADD supervisord.conf /etc/supervisor/conf.d/supervisord.conf

EXPOSE 1080
EXPOSE 8123

#CMD shadowsocks-local -c /etc/shadowsocks/config.json -m rc4-md5
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]
