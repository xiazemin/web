#把main.go build成WebAssembly(简写为wasm)二进制文件
GOOS=js GOARCH=wasm go build -o lib.wasm helloworld.go
#把JavaScript依赖拷贝到当前路径

 cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .


#$ ls  /usr/local/go/misc/wasm/
#go_js_wasm_exec wasm_exec.html  wasm_exec.js

#创建一个index.html文件，并引入wasm_exec.js文件，调用刚才build的lib.wasm
