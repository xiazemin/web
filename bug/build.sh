#!/bin/bash
#curd=`pwd`
 #GOPATH=$GOPATH:/usr/local/go/:$curd
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bug main.go
chmod +x bug
scp bug xia@10.96.83.51:/home/aladdin/

#scp argus xia@10.179.87.83:/home/aladdin/
#nohup ./argus >>argus.log &

#
#
# nohup ./bug >bug.log &