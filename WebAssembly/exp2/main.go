package main

import (
	"strconv"
	//"github.com/gopherjs/gopherjs/js"
	"syscall/js"
)
// 传入value1, value2, result三个元素的id，将value1+value2结果赋给result元素
func add(args js.Value,ids []js.Value) interface{}{

	// 根据id获取输入值
	value1 := js.Global().Get("document").Call("getElementById", ids[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", ids[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)
	// 将相加结果set给result元素
	js.Global().Get("document").Call("getElementById", ids[2].String()).Set("value", int1+int2)
	return nil
}

// 添加监听事件
func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
}

func main() {
	c := make(chan struct{}, 0)
	println("Go WebAssembly Initialized!")
	registerCallbacks()

	<-c
}