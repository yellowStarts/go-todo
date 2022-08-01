####### base
FROM golang:1.17-alpine AS base
WORKDIR /app

ENV GO111MODULE=on
ENV GOOS="linux"
ENV CGO_ENABLED=0
# 设置时区
ENV TZ=Asia/Shanghai
# 设置 Go 代理
ENV GOPROXY=https://goproxy.cn,direct

# RUN apk update \
#     && apk add --no-cache \
#     ca-certificates \
#     curl \
#     tzdata \
#     git \
#     && update-ca-certificates

RUN apk --no-cache add ca-certificates

####### dev
FROM base AS dev
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest && go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 5000
EXPOSE 2345

ENTRYPOINT ["air", "-c", ".air.toml"]

####### builder
FROM base AS builder
WORKDIR /app

COPY . .

RUN go mod download \
    && go mod verify

RUN go build -o ./todo .

####### prod
FROM alpine:latest AS prod
WORKDIR /usr/local/bin/

COPY --from=builder /app/todo .
EXPOSE 5000

CMD ["./todo"]