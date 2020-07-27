# syntax=docker/dockerfile:experimental
FROM reg.haochang.tv/base/golang:1.14.4 as builder

ARG BOT_ACCESS_TOKEN

ENV GOPROXY=https://goproxy.haochang.tv,https://goproxy.cn,direct \
    GOPRIVATE=gitlab.haochang.tv \
    CGO_ENABLED=0 \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /src
COPY . .
RUN git config --global url."https://bot:${BOT_ACCESS_TOKEN}@gitlab.haochang.tv".insteadOf "https://gitlab.haochang.tv"
# 使用时修改go build后面参数
RUN --mount=type=cache,target=/go go build -o bin/app cmd/server/main.go \
  && go build -o bin/app2 cmd/server/main.go



FROM reg.haochang.tv/base/alpine:3.10

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
 && apk update \
 && apk add --no-cache \
 && apk add ca-certificates openssl openssl-dev
COPY --from=builder /src/bin /usr/local/bin/
COPY ./configs/server /usr/local/etc/app

# 修改具体命令和参数
CMD ["app", "-config=/usr/local/etc/app/dev/config.toml"]
