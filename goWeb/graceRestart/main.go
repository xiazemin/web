// main
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"syscall"
	"time"
)

func main() {

	var listener1 *net.TCPListener

	if os.Getenv("_GRACEFUL_RESTART") == "true" {
		file := os.NewFile(3, "/tmp/sock-go-graceful-restart")

		listener1, err := net.FileListener(file)
		if err != nil {
			// handle
		}
		var ok bool
		listener1, ok = listener1.(*net.TCPListener)
		if !ok {
			// handle
		}
	} else {
		//listener1, err = new()ListenerWithPort(12345)
	}

	listenerFile1, err := listener1.File()

	if err != nil {
		log.Fatalln("Fail to get socket file descriptor:", err)
	}

	listenerFd := listenerFile1.Fd() // Set a flag for the new process start process

	os.Setenv("_GRACEFUL_RESTART", "true")
	execSpec1 := &syscall.ProcAttr{
		Env:   os.Environ(),
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd(), listenerFd},
	} // Fork exec the new version of your server
	fork1, err := syscall.ForkExec(os.Args[0], os.Args, execSpec1)

	syscall.Dup2(int(fork1), int(os.Stderr.Fd()))
	//由于标准库里提供了sync.WaitGroup结构体
	wg := sync.WaitGroup{}
	for {
		conn1, err := listener1.Accept()
		if err != nil {
			log.Fatal("connect failed ")

		}
		wg.Add(1)
		go func() {
			handle(conn1)
			wg.Done()
		}()
	}

	timeout := time.NewTimer(time.Minute)
	wait := make(chan struct{})
	go func() {
		wg.Wait()
		wait <- struct{}{}
	}()
	select {
	case <-timeout.C:
		//return nil
		//wait //, timeout.Error
	case <-wait:
		//return nil
	}

}

func Log(v ...interface{}) {
	fmt.Println(v...)
}

func handle(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		Log(conn.RemoteAddr().String(), "receive data length:", n)
		Log(conn.RemoteAddr().String(), "receive data:", buffer[:n])
		Log(conn.RemoteAddr().String(), "receive data string:", string(buffer[:n]))
	}
}
