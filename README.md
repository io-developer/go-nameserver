# go-nameserver
Golang DNS nameserver supporting hostname health check and foreign DNS upstream

[![Docker Automated build](https://img.shields.io/docker/automated/iodeveloper/go-nameserver.svg)](https://hub.docker.com/repository/docker/iodeveloper/go-nameserver/)

## Docker
[![Docker Hub repository](http://dockeri.co/image/iodeveloper/go-nameserver)](https://registry.hub.docker.com/r/iodeveloper/go-nameserver)

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
In server logs:
![image](https://user-images.githubusercontent.com/6779324/82758535-cb31bb00-9def-11ea-989d-d721636fbd63.png)

DIG response:
![image](https://user-images.githubusercontent.com/6779324/82758550-da186d80-9def-11ea-8022-24471327abc5.png)
