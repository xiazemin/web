package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	fmt.Printf("pid1：=%d\n", os.Getpid())
	parent := syscall.Getppid()
	//你会发现这个进程使用完全相同的参数os.Args启动了一个新进程。
	execSpec := &syscall.ProcAttr{
		Env:   os.Environ(),
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
	}

	fork, err := syscall.ForkExec(os.Args[0], os.Args, execSpec)
	if err != nil {
		fmt.Printf("error")
	}

	log.Printf("main: Killing parent pid: %v\n", parent)
	syscall.Kill(parent, syscall.SIGTERM)

	syscall.Dup2(int(fork), int(os.Stderr.Fd()))

	fmt.Printf("pid2：=%d\n", os.Getpid())
	syscall.Kill(syscall.Getppid(), syscall.SIGTERM)
}
