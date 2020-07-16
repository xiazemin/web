package main

import (
	"time"
)

var quit chan struct{} = make(chan struct{})

func f() {
	<-quit
}

func main() {
	for i := 0; i < 10000; i++ {
		go f()
	}

	go goppf() //启用跟踪查看
	for {
		time.Sleep(1 * time.Second)
	}
}
