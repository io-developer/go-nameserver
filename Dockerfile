FROM debian:stable-slim

RUN apt-get update \
    && apt-get install -y --no-install-recommends net-tools iputils-ping \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

ADD bin/go-nameserver /go-nameserver

EXPOSE 53

ENTRYPOINT ["/go-nameserver"]