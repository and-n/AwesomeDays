#!/usr/bin/env bash
echo "Installation..."
git pull

cd ./src/

go build server.go

pkill -f "server"

nohup ./server & >> ../server.log

echo "completed"
