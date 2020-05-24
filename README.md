# go-nameserver
Golang DNS nameserver supporting hostname health check and foreign DNS upstream

[![Docker Automated build](https://img.shields.io/docker/automated/iodeveloper/go-nameserver.svg)](https://hub.docker.com/repository/docker/iodeveloper/go-nameserver/)

## Docker
[![Docker Hub repository](http://dockeri.co/image/iodeveloper/go-nameserver)](https://registry.hub.docker.com/u/iodeveloper/go-nameserver/)

`iodeveloper/go-nameserver:latest`

## Example docker-compose.yml
```yml
version: '3.4'

services:
  local:
    image: iodeveloper/go-nameserver:latest
    restart: always
    command: ["--upstream", "tun:53", "--verbose"]
    ports:
       - '1053:53/udp'
    volumes:
       - ./records-local.json:/records.json

  tun:
    image: iodeveloper/go-nameserver:latest
    restart: always
    command: ["--upstream", "1.1.1.1:53", "--verbose"]
    ports:
       - '1054:53/udp'
    volumes:
       - ./records-tun.json:/records.json
```


## Example local run

Run server:
```bash
go run main.go --listen '0.0.0.0:1053' --records './records.json' --upstream '8.8.8.8:53' --verbose
```

And get dns in other cli:
```bash
dig node2.local @127.0.0.1 -p 1053
```

Server logs:


DIG response:
