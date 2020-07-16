package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime/pprof"
	"time"
)

func hello(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "hello world")
}

func cpuProfile() {
	now := time.Now()
	f, err := os.Create(fmt.Sprintf("/Users/didi/goLang/src/pprof/net/tmp/cpu.profile.%4d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "profile cpu err:%v", err)
		return
	}
	time.Sleep(time.Second * 10)
	pprof.StartCPUProfile(f)
	time.Sleep(time.Second * 30)
	pprof.StopCPUProfile()
}

func memProfile() {
	now := time.Now()
	f, err := os.Create(fmt.Sprintf("/Users/didi/goLang/src/pprof/net/tmp/mem.profile.%4d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "profile mem err:%v", err)
		return
	}
	time.Sleep(time.Second * 10)
	pprof.WriteHeapProfile(f)
	time.Sleep(time.Second * 30)
	f.Close()
}

func main() {
	go cpuProfile()
	go memProfile()
	print("hello world\n")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8001", nil)
	// go func() {
	// 	log.Println(http.ListenAndServe("localhost:6060", nil))
	// }()
}
