GOARCH=wasm GOOS=js go build -o test.wasm main.go
 cp /usr/local/go/misc/wasm/* .