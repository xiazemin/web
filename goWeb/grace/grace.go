package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"time"
)

//var netListener *net.TCPListener
var netListener GracefulListener
var ptrOffset = make([]uintptr, 0, 128)

//var ptr = []uint{0}

func ForkProcess() {
	//ptrOffset = make(uint)

	//	service := ":8888"
	//	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	//	netListener, err = net.ListenTCP("tcp", tcpAddr)

	file := netListener.File() // this returns a Dup()
	//path := "./grace"
	path := os.Args[0]
	args := []string{
		"-graceful"}

	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{file}

	env := append(
		os.Environ(),
	)
	cmd.Env = env

	err = cmd.Start()
	if err != nil {
		log.Fatalf("gracefulRestart: Failed to launch, error: %v", err)
	}

}

var (
	gracefulChild bool
	l             net.Listener
	err           error
)

//var server = &http.Server{Addr: "127.0.0.1:8888"}

var server = &http.Server{
	Addr:           "127.0.0.1:8888",
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 16,
}

func ChildInit() {

	flag.BoolVar(&gracefulChild, "graceful", false, "listen on fd open 3 (internal use only)")
	fmt.Printf("before new child process%d", os.Getpid())
	gracefulChild = true
	if gracefulChild {
		log.Print("main: Listening to existing file descriptor 3.")
		f := os.NewFile(3, "")
		l, err = net.FileListener(f)
	} else {
		log.Print("main: Listening on a new file descriptor.")
		l, err = net.Listen("tcp", server.Addr)
	}
	fmt.Printf("after new child process %d", os.Getpid())
}

func KillParent() {
	if gracefulChild {
		parent := os.Getppid() //syscall.Getppid()
		fmt.Printf("main: Killing parent pid: %v", parent)
		syscall.Kill(parent, syscall.SIGTERM)
	}
	fmt.Printf("after kill parent%d", os.Getpid())
	//server.Serve(l)
	netListener1 := newGracefulListener(l)
	server.Serve(netListener1)
}

type GracefulListener struct {
	net.Listener
	stop    chan error
	stopped bool
}

var httpWg sync.WaitGroup

func (gl *GracefulListener) Accept() (c net.Conn, err error) {
	c, err = gl.Listener.Accept()
	if err != nil {
		return
	}

	c = gracefulConn{Conn: c}

	httpWg.Add(1)
	return
}
func newGracefulListener(l net.Listener) (gl *GracefulListener) {
	gl = &GracefulListener{Listener: l, stop: make(chan error)}
	//	go func() {
	//		_ = <-gl.stop
	//		gl.stopped = true
	//		gl.stop <- gl.Listener.Close()
	//	}()
	return
}

func (gl *GracefulListener) Close() error {
	if gl.stopped {
		return syscall.EINVAL
	}
	gl.stop <- nil
	return <-gl.stop
}

func (gl *GracefulListener) File() *os.File {
	tl := gl.Listener.(*net.TCPListener)
	fl, _ := tl.File()
	return fl
}

type gracefulConn struct {
	net.Conn
}

func (w gracefulConn) Close() error {
	httpWg.Done()
	return w.Conn.Close()
}

//usage
//netListener = newGracefulListener(l)
//server.Serve(netListener)

//server := &http.Server{
//        Addr:           "0.0.0.0:8888",
//        ReadTimeout:    10 * time.Second,
//        WriteTimeout:   10 * time.Second,
//        MaxHeaderBytes: 1 << 16}
//}
