FROM golang:1.9-alpine
MAINTAINER zterry <zterry@qq.com>

RUN apk add --no-cache git supervisor

## https://hub.docker.com/r/vimagick/polipo/~/dockerfile/
RUN set -xe \
    && apk add --no-cache build-base openssl \
    && wget https://github.com/jech/polipo/archive/master.zip -O polipo.zip \
    && unzip polipo.zip \
    && cd polipo-master \
    && make \
    && install polipo /usr/local/bin/ \
    && cd .. \
    && rm -rf polipo.zip polipo-master \
    && mkdir -p /usr/share/polipo/www /var/cache/polipo \
    && apk del build-base openssl

RUN go get github.com/shadowsocks/shadowsocks-go/cmd/shadowsocks-local

ADD polipo.conf /etc/polipo.conf
ADD supervisord.conf /etc/supervisor/conf.d/supervisord.conf

EXPOSE 1080
EXPOSE 8123

#CMD shadowsocks-local -c /etc/shadowsocks/config.json -m rc4-md5
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]
