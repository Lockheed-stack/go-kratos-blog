FROM golang:1.21 AS builder

COPY . /src
WORKDIR /src
ENV CGO_ENABLED=0
RUN GOPROXY=https://goproxy.cn make build

# FROM debian:stable-slim

# RUN apt-get update && apt-get install -y --no-install-recommends \
# 		ca-certificates  \
#         netbase \
#         && rm -rf /var/lib/apt/lists/ \
#         && apt-get autoremove -y && apt-get autoclean -y
FROM alpine
COPY --from=builder /src/bin /app

# set time zone
RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]
