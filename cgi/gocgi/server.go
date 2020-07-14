package main

import (
	"log"
	"net/http"
	"net/http/cgi"
	"path/filepath"
	"os"
	//"os/exec"
	//"fmt"
	"fmt"
	"os/user"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {


		if u, err := user.Current(); err == nil {
			fmt.Println("用户ID: " + u.Uid)
			fmt.Println("主组ID: " + u.Gid)
			fmt.Println("用户名: " + u.Username)
			fmt.Println("主组名: " + u.Name)
			fmt.Println("家目录: " + u.HomeDir)
		}

		handler := new(cgi.Handler)

		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}

		//cmd := exec.Command("pwd")
		//fmt.Println(cmd)
		//dir:=cmd.Dir
		handler.Path = "/usr/local/go/bin/go"//dir+"/cgi"
		script := dir+"/cgi" + r.URL.Path
		log.Println(handler.Path)
		handler.Dir =  dir+"/cgi"

		//f, err := os.Create(script)
		//defer f.Close()
		//if err = f.Chmod(0777); err !=nil{
		  //fmt.Println(err,script)
		//}
		//args := []string{"sudo go run -x ", script}
		//args := []string{"php ", script}
		args := []string{"run", script}
		handler.Args = append(handler.Args, args...)
		handler.Env = append(handler.Env, "GOPATH=/Users/didi/goLang")
		handler.Env = append(handler.Env, "GOROOT=/usr/local/go")
		handler.Env = append(handler.Env, "PATH="+dir+"/cgi")
		//export TMPDIR=/run/shm
		handler.Env = append(handler.Env,"TMPDIR="+dir+"/cgi")

		log.Println(handler.Args)

		handler.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":8086", nil))
}