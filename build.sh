#!/bin/bash

CURRENT_COMMIT=$(git rev-parse --short HEAD)
CURENT_DATE=$(date '+%Y%m%d%H%M%S')
BUILD_NAME=atmm$CURENT_DATE-$CURRENT_COMMIT
rm ./bin/* 2> /dev/null
env GOOS=windows go build -race -o ./bin/$BUILD_NAME-windows.exe
env GOOS=linux go build -o ./bin/$BUILD_NAME-linux