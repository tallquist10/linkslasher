#!/bin/bash
PORT="${2:-3000}"
SQL_MODE="${3:-SQLITE}"
docker build -t "linkslasher:${1}" --build-arg GO_VERSION="${1}" .
if [ "$(docker ps -a | grep linkslasher:${1})" ]; then 
    docker stop "linkslasher${1}"
    docker rm "linkslasher${1}"
fi
docker run -d -e SQL_MODE=$SQL_MODE -p $PORT:8080 --name "linkslasher${1}" "linkslasher:${1}"