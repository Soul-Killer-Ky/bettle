FROM golang:1.18 AS builder

ARG APP

COPY . /src
WORKDIR /src/app/$APP

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

ARG APP

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/app/$APP/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

ENV entry=$APP

#CMD ["./server", "-conf", "/data/conf"]
CMD ./$entry -conf /data/conf
