#!/bin/bash

echo start building

env GOOS=linux GOARCH=amd64 go build -mod=vendor -o bot_msg_console main.go
scp bot_msg_console root@39.96.21.121:/home/works/zhouyi/bot_msg_console

rm -f bot_msg_console 
echo build finished