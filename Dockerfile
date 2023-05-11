FROM arm64v8/golang:1.19-alpine3.15 AS build-env
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
     apk add --no-cache git sqlite-libs sqlite-dev build-base && rm -rf /var/cache/apk/*
ADD . /siyuanzhong/gobuild
WORKDIR /siyuanzhong/gobuild
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOPROXY=https://goproxy.cn,direct
# 澶嶅埗椤圭洰涓殑 go.mod 鍜?go.sum鏂囦欢骞朵笅杞戒緷璧栦俊鎭?
RUN go build -o engineserver ./server/main.go
FROM alpine
WORKDIR /siyuanzhong
COPY --from=build-env /siyuanzhong/gobuild/engineserver /siyuanzhong/
COPY --from=build-env /siyuanzhong/gobuild/database /siyuanzhong/
ENV SERVER_PORT 10080
ENV MQTT_BROKER tcp://192.168.8.251:1883
ENV DATABASE_DIR /siyuanzhong/
EXPOSE 10080
CMD ["/siyuanzhong/engineserver"]

