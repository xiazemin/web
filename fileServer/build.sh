#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o fileServer  main.go
scp fileServer xiazemin@10.96.83.51:/home/aladdin
scp ~/Downloads/conf.tar  xiazemin@10.96.83.51:/home/aladdin/download
