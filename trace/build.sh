CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o trace main.go
scp trace xia@10.96.83.51:/home/src/github.com/xiazemin/trace