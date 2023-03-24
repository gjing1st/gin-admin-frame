#!/bin/bash
#导入tar
cd
docker load < gaf-frontend.tar
docker load < gaf-backend.tar
#重启docker compose
cd /home/app/gaf
docker-compose  down
docker-compose up -d
echo "finish"
