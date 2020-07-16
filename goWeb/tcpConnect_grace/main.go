package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
)

func main() {
	fmt.Printf("pid1：=%d", os.Getpid())
	//你会发现这个进程使用完全相同的参数os.Args启动了一个新进程。
	execSpec := &syscall.ProcAttr{
		Env:   os.Environ(),
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
	}

	fork, err := syscall.ForkExec(os.Args[0], os.Args, execSpec)
	syscall.Dup2(int(fork), int(os.Stderr.Fd()))

	fmt.Printf("pid2：=%d", os.Getpid())

	fmt.Println("Hello World!")
	//tcp监听服务
	service := ":8010"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		fmt.Printf("ResolveTCPAddr")
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Printf("ListenTCP")
	}

	for {
		conn, err := listener.Accept()
		defer conn.Close()
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				fmt.Println("Stop accepting connections")
				return
			}
		}
		conn.Write([]byte{'f', 'i', 'n', 'i', 's', 'h'})
	}
}
