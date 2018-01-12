#!/bin/bash
mkdir bin
cp VHostTemplate.conf bin/VHostTemplate.conf
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/apacheVHost main.go
