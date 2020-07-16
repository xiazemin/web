// main
package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"syscall"
)

func main() {
	var netListener *net.TCPListener

	service := ":9090"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	netListener, err = net.ListenTCP("tcp", tcpAddr)

	//file0 := os.NewFile(3, "/tmp/sock-go-graceful-restart")
	//netListener1, _ := net.FileListener(file0)
	//netListener = netListener1.(*net.TCPListener)

	var file *os.File
	file, _ = netListener.File() // this returns a Dup()
	path := "./graceSimple"
	args := []string{
		"-graceful"}

	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{file}

	err = cmd.Start()
	if err != nil {
		log.Fatalf("gracefulRestart: Failed to launch, error: %v", err)
	}

	server := &http.Server{Addr: "0.0.0.0:8888"}

	var gracefulChild bool
	var l net.Listener
	//	var err error

	flag.BoolVar(&gracefulChild, "-graceful", false, "listen on fd open 3 (internal use only)")

	if gracefulChild {
		log.Print("main: Listening to existing file descriptor 3.")
		f := os.NewFile(3, "")
		l, err = net.FileListener(f)
	} else {
		log.Print("main: Listening on a new file descriptor.")
		l, err = net.Listen("tcp", server.Addr)
	}

	if gracefulChild {
		parent := syscall.Getppid()
		log.Printf("main: Killing parent pid: %v", parent)
		syscall.Kill(parent, syscall.SIGTERM)
	}

	server.Serve(l)

}

var httpWg sync.WaitGroup

type gracefulListener struct {
	net.Listener
	stop    chan error
	stopped bool
}

func (gl *gracefulListener) Accept() (c net.Conn, err error) {
	c, err = gl.Listener.Accept()
	if err != nil {
		return
	}

	c = gracefulConn{Conn: c}

	httpWg.Add(1)
	return
}

func newGracefulListener(l net.Listener) (gl *gracefulListener) {
	gl = &gracefulListener{Listener: l, stop: make(chan error)}
	go func() {
		_ = <-gl.stop
		gl.stopped = true
		gl.stop <- gl.Listener.Close()
	}()
	return
}

func (gl *gracefulListener) Close() error {
	if gl.stopped {
		return syscall.EINVAL
	}
	gl.stop <- nil
	return <-gl.stop
}

func (gl *gracefulListener) File() *os.File {
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
