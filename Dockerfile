FROM golang:1.14 as builder

WORKDIR /tmp/proj/
ENV GOBIN="$(pwd)/bin"
ENV CGO_ENABLED=0

COPY . ./

RUN go build -o ./bin/go-nameserver -tags netgo -a
#RUN go install

FROM debian:stable-slim as runtime

RUN apt-get update \
    && apt-get install -y --no-install-recommends net-tools iputils-ping \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /tmp/proj/bin/go-nameserver /go-nameserver

EXPOSE 53/udp
EXPOSE 53/tcp

ENTRYPOINT ["/go-nameserver"]