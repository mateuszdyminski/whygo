#!/usr/bin/env bash

docker build -t dummy .

docker run -p 9001:9001 -d --name dummy1 --rm dummy /app/dummy 9001
docker run -p 9002:9002 -d --name dummy2 --rm dummy /app/dummy 9002
docker run -p 9003:9003 -d --name dummy3 --rm dummy /app/dummy 9003
