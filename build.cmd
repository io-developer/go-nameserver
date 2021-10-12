@echo off

set GOBIN=%~dp0bin\
set CGO_ENABLED=0

go build -o %~dp0bin\go-nameserver.exe -tags netgo -a
::go install
