package main

import (
	"errors"
	"log"
	"os"
	"runtime"
	"syscall"
)

func main() {

}
func daemon(chdir bool, closeStd bool) (err error) {
	var ret, ret2 uintptr
	var er syscall.Errno
	darwin := runtime.GOOS == "darwin"
	// already a daemon
	if syscall.Getppid() == 1 {
		return
	}
	// fork off the parent process
	ret, ret2, er = syscall.RawSyscall(syscall.SYS_FORK, 0, 0, 0)
	if er != 0 {
		err = errors.New("fork fail")
		return
	}
	// failure
	if ret2 < 0 {
		err = errors.New("fork fail")
		os.Exit(-1)
	}
	// handle exception for darwin
	if darwin && ret2 == 1 {
		ret = 0
	}
	// if we got a good PID, then we call exit the parent process.
	if ret > 0 {
		os.Exit(0)
	}
	/* Change the file mode mask */
	_ = syscall.Umask(0)
	// create a new SID for the child process
	s_ret, s_errno := syscall.Setsid()
	if s_errno != nil {
		log.Printf("Error: syscall.Setsid errno: %d", s_errno)
	}
	if s_ret < 0 {
		err = errors.New("setsid fail")
		return
	}
	if chdir {
		os.Chdir("/")
	}
	if closeStd {
		f, e := os.OpenFile("/dev/null", os.O_RDWR, 0)
		if e == nil {
			fd := f.Fd()
			syscall.Dup2(int(fd), int(os.Stdin.Fd()))
			syscall.Dup2(int(fd), int(os.Stdout.Fd()))
			syscall.Dup2(int(fd), int(os.Stderr.Fd()))
		}
	}
	return
}
