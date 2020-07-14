package main

import "fmt"
import (
	"os"
	"path/filepath"
	"log"
	"os/user"
	"net/http/cgi"
	"net/http"
)

func myHandler(w http.ResponseWriter,r *http.Request){
	if u, err := user.Current(); err == nil {
		fmt.Println("用户ID: " + u.Uid)
		fmt.Println("主组ID: " + u.Gid)
		fmt.Println("用户名: " + u.Username)
		fmt.Println("主组名: " + u.Name)
		fmt.Println("家目录: " + u.HomeDir)
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	ip := os.Getenv("REMOTE_ADDR")
	fmt.Printf("Content-type: text/plain\n\n")
	fmt.Println(ip)
	w.Write([]byte("Content-type: text/plain\n\n"))
}
func main() {
	cgi.Serve(http.HandlerFunc(myHandler))

}
