#!/bin/bash  
docker rm -f base
docker build --tag base .
docker run --network=kong-net -d -p 8079:8079 --name base base
docker system prune -f
docker ps
exit
