FROM arm64v8/golang:1.19-alpine3.15
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
     apk add --no-cache git sqlite-libs sqlite-dev build-base && rm -rf /var/cache/apk/*
VOLUME /siyuanzhong/gobuild
WORKDIR /siyuanzhong/gobuild
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOPROXY=https://goproxy.cn,direct
ENTRYPOINT ["/bin/sh"]

