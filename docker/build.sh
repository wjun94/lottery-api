#!bin/bash

echo "打包"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go